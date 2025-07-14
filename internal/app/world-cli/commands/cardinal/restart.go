package cardinal

import (
	"context"

	"github.com/argus-labs/world-cli/v2/internal/app/world-cli/models"
	"github.com/argus-labs/world-cli/v2/internal/app/world-cli/shared/config"
	"github.com/argus-labs/world-cli/v2/internal/app/world-cli/shared/docker"
	"github.com/argus-labs/world-cli/v2/internal/pkg/logger"
)

func (h *Handler) Restart(ctx context.Context, f models.RestartCardinalFlags) error {
	cfg, err := config.GetConfig(&f.Config)
	if err != nil {
		return err
	}
	cfg.Build = true
	cfg.Debug = f.Debug
	cfg.Detach = f.Detach

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

	err = dockerClient.Restart(ctx, getServices(cfg)...)
	if err != nil {
		return err
	}

	return nil
}
