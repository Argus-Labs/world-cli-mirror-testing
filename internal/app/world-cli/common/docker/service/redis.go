package service

import (
	"fmt"
	"strconv"

	"github.com/argus-labs/go-ecs/internal/app/world-cli/common/config"
	"github.com/argus-labs/go-ecs/internal/pkg/logger"
	"github.com/docker/docker/api/types/container"
)

func getRedisContainerName(cfg *config.Config) string {
	return fmt.Sprintf("%s-redis", cfg.DockerEnv["CARDINAL_NAMESPACE"])
}

func Redis(cfg *config.Config) Service {
	// Check cardinal namespace
	checkCardinalNamespace(cfg)

	redisPort := cfg.DockerEnv["REDIS_PORT"]
	if redisPort == "" {
		redisPort = "6379"
	}

	intPort, err := strconv.Atoi(redisPort)
	if err != nil {
		logger.Error("Failed to convert redis port to int, defaulting to 6379", err)
		intPort = 6379
	}
	exposedPorts := []int{intPort}

	return Service{
		Name: getRedisContainerName(cfg),
		Config: container.Config{
			Image:        "redis:latest",
			ExposedPorts: getExposedPorts(exposedPorts),
		},
		HostConfig: container.HostConfig{
			PortBindings:  newPortMap(exposedPorts),
			RestartPolicy: container.RestartPolicy{Name: "unless-stopped"},
			NetworkMode:   container.NetworkMode(cfg.DockerEnv["CARDINAL_NAMESPACE"]),
		},
	}
}
