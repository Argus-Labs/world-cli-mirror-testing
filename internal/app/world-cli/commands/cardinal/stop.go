package cardinal

import (
	"context"

	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/app/world-cli/models"
	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/app/world-cli/shared/config"
	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/app/world-cli/shared/docker"
	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/app/world-cli/shared/docker/service"
	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/pkg/logger"
	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/pkg/printer"
)

func (h *Handler) Stop(ctx context.Context, f models.StopCardinalFlags) error {
	cfg, err := config.GetConfig(&f.Config)
	if err != nil {
		return err
	}

	// Create docker client
	dockerClient, err := docker.NewClient(cfg)
	if err != nil {
		return err
	}
	defer func() {
		if err := dockerClient.Close(); err != nil {
			logger.Error("Failed to close docker client", "error", err)
		}
	}()

	err = dockerClient.Stop(ctx, service.Nakama, service.Cardinal,
		service.NakamaDB, service.Redis, service.Jaeger, service.Prometheus)
	if err != nil {
		return err
	}

	printer.Successln("Cardinal successfully stopped")

	return nil
}
