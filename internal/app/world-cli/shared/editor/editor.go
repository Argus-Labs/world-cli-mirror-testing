package editor

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/goccy/go-json"

	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/app/world-cli/shared/config"
	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/pkg/logger"
	"github.com/google/uuid"
	"github.com/rotisserie/eris"
	"golang.org/x/mod/modfile"
)

const (
	Dir = ".editor"

	downloadURLPattern = "https://github.com/Argus-Labs/cardinal-editor/releases/download/%s/cardinal-editor-%s.zip"

	httpTimeout = 2 * time.Second

	cardinalProjectIDPlaceholder = "__CARDINAL_PROJECT_ID__"

	cardinalPkgPath = "pkg.world.dev/world-engine/cardinal"

	versionMapURL = "https://raw.githubusercontent.com/Argus-Labs/cardinal-editor/main/version_map.json"
)

type Asset struct {
	BrowserDownloadURL string `json:"browser_download_url"`
}

type Release struct {
	Name   string  `json:"name"`
	Assets []Asset `json:"assets"`
}

// getDefaultCardinalVersionMap returns the default version map for Cardinal Editor.
// This is used when the version map cannot be fetched from the repository.
func getDefaultCardinalVersionMap() map[string]string {
	return map[string]string{
		"v1.2.2-beta": "v0.1.0",
		"v1.2.3-beta": "v0.1.0",
		"v1.2.4-beta": "v0.3.1",
		"v1.2.5-beta": "v0.3.1",
	}
}

func SetupCardinalEditor(rootDir string, gameDir string) error {
	// Get the version map
	cardinalVersionMap, err := getVersionMap(versionMapURL)
	if err != nil {
		logger.Warn("Failed to get version map, using default version map")
		cardinalVersionMap = getDefaultCardinalVersionMap()
	}

	// Check version
	cardinalVersion, err := getModuleVersion(filepath.Join(rootDir, gameDir, "go.mod"), cardinalPkgPath)
	if err != nil {
		return eris.Wrap(err, "failed to get cardinal version")
	}

	downloadVersion, versionExists := cardinalVersionMap[cardinalVersion]
	if !versionExists {
		// Get the latest release version
		latestReleaseVersion, err := getLatestReleaseVersion()
		if err != nil {
			return eris.Wrap(err, "failed to get latest release version")
		}
		downloadVersion = latestReleaseVersion
	}

	downloadURL := fmt.Sprintf(downloadURLPattern, downloadVersion, downloadVersion)

	// Check if the Cardinal Editor directory exists
	targetEditorDir := filepath.Join(rootDir, Dir)
	if _, err := os.Stat(targetEditorDir); !os.IsNotExist(err) {
		// Check the version of cardinal editor is appropriate
		if fileExists(filepath.Join(targetEditorDir, downloadVersion)) {
			// do nothing if the version is already downloaded
			return nil
		}

		// Remove the existing Cardinal Editor directory
		if err := os.RemoveAll(targetEditorDir); err != nil {
			logger.Error("Failed to remove existing Cardinal Editor directory", "error", err)
		}
	}

	configDir, err := config.GetCLIConfigDir()
	if err != nil {
		return err
	}

	editorDir, err := downloadReleaseIfNotCached(downloadVersion, downloadURL, configDir)
	if err != nil {
		return err
	}

	// rename version tag dir to .editor
	err = copyDir(editorDir, targetEditorDir)
	if err != nil {
		return err
	}

	// rename project id
	// "ce" prefix is added because guids can start with numbers, which is not allowed in js
	projectID := "ce" + strippedGUID()
	err = replaceProjectIDs(targetEditorDir, projectID)
	if err != nil {
		return err
	}

	// this file is used to check if the version is already downloaded
	err = addFileVersion(filepath.Join(targetEditorDir, downloadVersion))
	if err != nil {
		return err
	}

	return nil
}

// returns editor directory path, and error.
func downloadReleaseIfNotCached(version, downloadURL, configDir string) (string, error) {
	editorDir := filepath.Join(configDir, "editor")
	targetDir := filepath.Join(editorDir, version)

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		return targetDir, downloadAndUnzip(downloadURL, targetDir)
	}

	return targetDir, nil
}

func downloadAndUnzip(url string, targetDir string) error {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{
		Timeout: httpTimeout + 10*time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return eris.New("Failed to download Cardinal Editor from " + url)
	}
	defer closeAndSetError(resp.Body, &err)

	tmpZipFileName := "tmp.zip"
	file, err := os.Create(tmpZipFileName)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		closeAndLogError(file)
		return err
	}
	closeAndLogError(file)

	if err = unzipFile(tmpZipFileName, targetDir); err != nil {
		return err
	}

	return os.Remove(tmpZipFileName)
}

//nolint:gocognit // Makes sense to consolidate everything in one function
func unzipFile(filename string, targetDir string) error {
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer closeAndSetError(reader, &err)

	// save original folder name
	var originalDir string
	for i, file := range reader.File {
		if i == 0 {
			originalDir = file.Name
		}

		src, err := file.Open()
		if err != nil {
			return err
		}

		filePath, err := sanitizeExtractPath(filepath.Dir(targetDir), file.Name)
		if err != nil {
			closeAndLogError(src)
			return err
		}
		if file.FileInfo().IsDir() {
			err = os.MkdirAll(filePath, 0755)
			if err != nil {
				closeAndLogError(src)
				return err
			}
			closeAndLogError(src)
			continue
		}

		dst, err := os.Create(filePath)
		if err != nil {
			closeAndLogError(src)
			return err
		}

		_, err = io.Copy(dst, src) //nolint:gosec // zip file is from us
		if err != nil {
			closeAndLogError(src)
			closeAndLogError(dst)
			return err
		}
		closeAndLogError(src)
		closeAndLogError(dst)
	}

	if err = os.Rename(filepath.Join(filepath.Dir(targetDir), originalDir), targetDir); err != nil {
		return err
	}

	return nil
}

func sanitizeExtractPath(dst string, filePath string) (string, error) {
	dstPath := filepath.Join(dst, filePath)
	if strings.HasPrefix(dstPath, filepath.Clean(dst)) {
		return dstPath, nil
	}
	return "", eris.Errorf("%s: illegal file path", filePath)
}

func copyDir(src string, dst string) error {
	srcDir, err := os.ReadDir(src)
	if err != nil {
		return eris.New("Failed to read directory " + src)
	}

	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}

	for _, entry := range srcDir {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			// Recursively copy dirs
			if err := copyDir(srcPath, destPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, destPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer closeAndSetError(srcFile, &err)

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer closeAndSetError(destFile, &err)

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func replaceProjectIDs(editorDir string, newID string) error {
	assetsDir := filepath.Join(editorDir, "assets")
	files, err := os.ReadDir(assetsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".js") {
			content, err := os.ReadFile(filepath.Join(assetsDir, file.Name()))
			if err != nil {
				return err
			}

			newContent := strings.ReplaceAll(string(content), cardinalProjectIDPlaceholder, newID)

			err = os.WriteFile(filepath.Join(assetsDir, file.Name()), []byte(newContent), 0600)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func strippedGUID() string {
	u := uuid.New()
	return strings.ReplaceAll(u.String(), "-", "")
}

// addFileVersion add file with name version of cardinal editor.
func addFileVersion(version string) error {
	// Create the file
	file, err := os.Create(version)
	if err != nil {
		return err
	}
	defer closeAndSetError(file, &err)

	return nil
}

func getModuleVersion(gomodPath, modulePath string) (string, error) {
	// Read the go.mod file
	data, err := os.ReadFile(gomodPath)
	if err != nil {
		return "", err
	}

	// Parse the go.mod file
	modFile, err := modfile.Parse(gomodPath, data, nil)
	if err != nil {
		return "", err
	}

	// Iterate through the require statements to find the desired module
	for _, require := range modFile.Require {
		if require.Mod.Path == modulePath {
			return require.Mod.Version, nil
		}
	}

	// Return an error if the module is not found
	return "", eris.Errorf("module %s not found", modulePath)
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// getVersionMap fetches the JSON data from a URL and unmarshals it into a map[string]string.
func getVersionMap(url string) (map[string]string, error) {
	// Make an HTTP GET request
	client := &http.Client{
		Timeout: httpTimeout,
	}

	// Create a new request using http
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Send the request via a client
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer closeAndSetError(resp.Body, &err)

	// Check for HTTP error
	if resp.StatusCode != http.StatusOK {
		return nil, eris.Errorf("HTTP error: %d - %s", resp.StatusCode, resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a map
	var result map[string]string
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getLatestReleaseVersion() (string, error) {
	latestReleaseURL := "https://github.com/Argus-Labs/cardinal-editor/releases/latest"

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, latestReleaseURL, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{
		Timeout: httpTimeout,
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			// Return an error to prevent following redirects
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer closeAndSetError(resp.Body, &err)

	// Check if the status code is 302
	// GitHub responds with a 302 redirect to the actual latest release URL, which contains the version number
	if resp.StatusCode != http.StatusFound {
		return "", eris.New("Failed to fetch the latest release of Cardinal Editor")
	}

	redirectURL := resp.Header.Get("Location")
	latestReleaseVersion := strings.TrimPrefix(
		redirectURL,
		"https://github.com/Argus-Labs/cardinal-editor/releases/tag/",
	)

	return latestReleaseVersion, nil
}

// closeAndLogError closes the given closer and logs any error that occurs.
// This is used for non-defer close operations where we just want to log the error.
func closeAndLogError(closer io.Closer) {
	if err := closer.Close(); err != nil {
		logger.Error("Failed to close resource", "error", err)
	}
}

// closeAndSetError closes the given closer and sets the error if one occurs.
// This is used in defer statements where we want to preserve the original error.
func closeAndSetError(closer io.Closer, err *error) {
	if closeErr := closer.Close(); closeErr != nil {
		if *err == nil {
			*err = closeErr
		}
	}
}
