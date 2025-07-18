package service

import (
	"fmt"
	"time"

	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/app/world-cli/shared/config"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

func getCelestiaDevNetContainerName(cfg *config.Config) string {
	return fmt.Sprintf("%s-celestia-devnet", cfg.DockerEnv["CARDINAL_NAMESPACE"])
}

func CelestiaDevNet(cfg *config.Config) Service {
	// Check cardinal namespace
	checkCardinalNamespace(cfg)

	return Service{
		Name: getCelestiaDevNetContainerName(cfg),
		Config: container.Config{
			Image:        "ghcr.io/rollkit/local-celestia-devnet:latest",
			ExposedPorts: getExposedPorts([]int{26658, 26659}),
			Healthcheck: &container.HealthConfig{
				Test:     []string{"CMD", "curl", "-f", "http://127.0.0.1:26659/head"},
				Interval: 1 * time.Second,
				Timeout:  1 * time.Second,
				Retries:  20,
			},
		},
		HostConfig: container.HostConfig{
			PortBindings: nat.PortMap{
				"26657/tcp": []nat.PortBinding{},
				"26658/tcp": []nat.PortBinding{{HostPort: "26658"}},
				"26659/tcp": []nat.PortBinding{{HostPort: "26659"}},
				"9090/tcp":  []nat.PortBinding{},
			},
			RestartPolicy: container.RestartPolicy{Name: "on-failure"},
			NetworkMode:   container.NetworkMode(cfg.DockerEnv["CARDINAL_NAMESPACE"]),
		},
	}
}
