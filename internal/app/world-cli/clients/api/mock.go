package api

import (
	"context"
	"net/http"

	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/app/world-cli/models"
	"github.com/stretchr/testify/mock"
)

// Ensure MockClient implements the interface.
var _ ClientInterface = (*MockClient)(nil)

// MockClient is a mock implementation of ClientInterface.
type MockClient struct {
	mock.Mock
}

// API-specific methods

// GetUser mocks getting user information.
func (m *MockClient) GetUser(ctx context.Context) (models.User, error) {
	args := m.Called(ctx)
	user, ok := args.Get(0).(models.User)
	if !ok {
		return models.User{}, args.Error(1)
	}

	return user, args.Error(1)
}

// UpdateUser mocks updating user information.
func (m *MockClient) UpdateUser(ctx context.Context, name, email string) error {
	args := m.Called(ctx, name, email)
	return args.Error(0)
}

// UpdateUserRoleInOrganization mocks updating user role in organization.
func (m *MockClient) UpdateUserRoleInOrganization(ctx context.Context, orgID, userEmail, role string) error {
	args := m.Called(ctx, orgID, userEmail, role)
	return args.Error(0)
}

// InviteUserToOrganization mocks inviting a user to an organization.
func (m *MockClient) InviteUserToOrganization(ctx context.Context, orgID, userEmail, role string) error {
	args := m.Called(ctx, orgID, userEmail, role)
	return args.Error(0)
}

// GetOrganizations mocks getting organizations.
func (m *MockClient) GetOrganizations(ctx context.Context) ([]models.Organization, error) {
	args := m.Called(ctx)
	organizations, ok := args.Get(0).([]models.Organization)
	if !ok {
		return nil, args.Error(1)
	}

	return organizations, args.Error(1)
}

// GetOrganizationsInvitedTo mocks getting organization invitations.
func (m *MockClient) GetOrganizationsInvitedTo(ctx context.Context) ([]models.Organization, error) {
	args := m.Called(ctx)
	organizations, ok := args.Get(0).([]models.Organization)
	if !ok {
		return nil, args.Error(1)
	}

	return organizations, args.Error(1)
}

// AcceptOrganizationInvitation mocks accepting an organization invitation.
func (m *MockClient) AcceptOrganizationInvitation(ctx context.Context, orgID string) error {
	args := m.Called(ctx, orgID)
	return args.Error(0)
}

// GetProjects mocks getting projects for an organization.
func (m *MockClient) GetProjects(ctx context.Context, orgID string) ([]models.Project, error) {
	args := m.Called(ctx, orgID)
	projects, ok := args.Get(0).([]models.Project)
	if !ok {
		return nil, args.Error(1)
	}

	return projects, args.Error(1)
}

// LookupProjectFromRepo mocks looking up a project from repository information.
func (m *MockClient) LookupProjectFromRepo(ctx context.Context, repoURL, repoPath string) (models.Project, error) {
	args := m.Called(ctx, repoURL, repoPath)
	project, ok := args.Get(0).(models.Project)
	if !ok {
		return models.Project{}, args.Error(1)
	}

	return project, args.Error(1)
}

// GetOrganizationByID mocks getting an organization by ID.
func (m *MockClient) GetOrganizationByID(ctx context.Context, id string) (models.Organization, error) {
	args := m.Called(ctx, id)
	org, ok := args.Get(0).(models.Organization)
	if !ok {
		return models.Organization{}, args.Error(1)
	}

	return org, args.Error(1)
}

// GetProjectByID mocks getting a project by ID.
func (m *MockClient) GetProjectByID(ctx context.Context, orgID, projID string) (models.Project, error) {
	args := m.Called(ctx, orgID, projID)
	project, ok := args.Get(0).(models.Project)
	if !ok {
		return models.Project{}, args.Error(1)
	}

	return project, args.Error(1)
}

// CreateOrganization mocks creating an organization.
func (m *MockClient) CreateOrganization(
	ctx context.Context,
	name, slug string,
) (models.Organization, error) {
	args := m.Called(ctx, name, slug)
	org, ok := args.Get(0).(models.Organization)
	if !ok {
		return models.Organization{}, args.Error(1)
	}

	return org, args.Error(1)
}

// GetListRegions mocks getting list of regions.
func (m *MockClient) GetListRegions(ctx context.Context, orgID, projID string) ([]string, error) {
	args := m.Called(ctx, orgID, projID)
	regions, ok := args.Get(0).([]string)
	if !ok {
		return nil, args.Error(1)
	}

	return regions, args.Error(1)
}

// CheckProjectSlugIsTaken mocks checking if a project slug is taken.
func (m *MockClient) CheckProjectSlugIsTaken(ctx context.Context, orgID, projID, slug string) error {
	args := m.Called(ctx, orgID, projID, slug)
	return args.Error(0)
}

// CreateProject mocks creating a project.
func (m *MockClient) CreateProject(ctx context.Context, orgID string, project models.Project) (models.Project, error) {
	args := m.Called(ctx, orgID, project)
	project, ok := args.Get(0).(models.Project)
	if !ok {
		return models.Project{}, args.Error(1)
	}

	return project, args.Error(1)
}

// UpdateProject mocks updating a project.
func (m *MockClient) UpdateProject(
	ctx context.Context,
	orgID, projID string,
	project models.Project,
) (models.Project, error) {
	args := m.Called(ctx, orgID, projID, project)
	project, ok := args.Get(0).(models.Project)
	if !ok {
		return models.Project{}, args.Error(1)
	}

	return project, args.Error(1)
}

// DeleteProject mocks deleting a project.
func (m *MockClient) DeleteProject(ctx context.Context, orgID, projID string) error {
	args := m.Called(ctx, orgID, projID)
	return args.Error(0)
}

// PreviewDeployment mocks previewing a deployment.
func (m *MockClient) PreviewDeployment(
	ctx context.Context,
	orgID, projID, deployType string,
) (models.DeploymentPreview, error) {
	args := m.Called(ctx, orgID, projID, deployType)
	preview, ok := args.Get(0).(models.DeploymentPreview)
	if !ok {
		return models.DeploymentPreview{}, args.Error(1)
	}

	return preview, args.Error(1)
}

// DeployProject mocks deploying a project.
func (m *MockClient) DeployProject(
	ctx context.Context,
	orgID, projID, deployType string,
) error {
	args := m.Called(ctx, orgID, projID, deployType)
	return args.Error(0)
}

// GetTemporaryCredential mocks getting temporary credential.

func (m *MockClient) GetTemporaryCredential(
	ctx context.Context,
	orgID, projID string,
) (models.TemporaryCredential, error) {
	args := m.Called(ctx, orgID, projID)
	credential, ok := args.Get(0).(models.TemporaryCredential)
	if !ok {
		return models.TemporaryCredential{}, args.Error(1)
	}

	return credential, args.Error(1)
}

// GetDeploymentStatus mocks getting deployment status.
func (m *MockClient) GetDeploymentStatus(ctx context.Context, projID string) ([]byte, error) {
	args := m.Called(ctx, projID)
	healthStatus, ok := args.Get(0).([]byte)
	if !ok {
		return nil, args.Error(1)
	}

	return healthStatus, args.Error(1)
}

// GetHealthStatus mocks getting health status.
func (m *MockClient) GetHealthStatus(ctx context.Context, projID string) ([]byte, error) {
	args := m.Called(ctx, projID)
	healthStatus, ok := args.Get(0).([]byte)
	if !ok {
		return nil, args.Error(1)
	}

	return healthStatus, args.Error(1)
}

// GetDeploymentHealthStatus mocks getting deployment health status.
func (m *MockClient) GetDeploymentHealthStatus(
	ctx context.Context,
	projID string,
) (map[string]models.DeploymentHealthCheckResult, error) {
	args := m.Called(ctx, projID)
	healthChecks, ok := args.Get(0).(map[string]models.DeploymentHealthCheckResult)
	if !ok {
		return nil, args.Error(1)
	}

	return healthChecks, args.Error(1)
}

// GetOrganizationMembers mocks getting organization members.
func (m *MockClient) GetOrganizationMembers(ctx context.Context, orgID string) ([]models.OrganizationMember, error) {
	args := m.Called(ctx, orgID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	members, ok := args.Get(0).([]models.OrganizationMember)
	if !ok {
		return nil, args.Error(1)
	}

	return members, args.Error(1)
}

// Utility methods

// SetAuthToken mocks setting auth token.
func (m *MockClient) SetAuthToken(token string) {
	m.Called(token)
}

// GetRPCBaseURL mocks getting RPC base URL. TODO: Remove this once we have a proper RPC client
func (m *MockClient) GetRPCBaseURL() string {
	return m.Called().String(0)
}

// Authentication methods

// GetLoginLink mocks getting the login link.
func (m *MockClient) GetLoginLink(ctx context.Context) (LoginLinkResponse, error) {
	args := m.Called(ctx)
	response, ok := args.Get(0).(LoginLinkResponse)
	if !ok {
		return LoginLinkResponse{}, args.Error(1)
	}

	return response, args.Error(1)
}

// GetLoginToken mocks getting the login token.
func (m *MockClient) GetLoginToken(ctx context.Context, callbackURL string) (models.LoginToken, error) {
	args := m.Called(ctx, callbackURL)
	loginToken, ok := args.Get(0).(models.LoginToken)
	if !ok {
		return models.LoginToken{}, args.Error(1)
	}

	return loginToken, args.Error(1)
}

// MockHTTPClient is a mock implementation of HTTPClientInterface for testing.
type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	response, ok := args.Get(0).(*http.Response)
	if !ok {
		return nil, args.Error(1)
	}

	return response, args.Error(1)
}
