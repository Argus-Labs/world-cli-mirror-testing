package evm

import (
	"context"

	"github.com/argus-labs/world-cli/v2/internal/app/world-cli/models"
	"github.com/argus-labs/world-cli/v2/internal/app/world-cli/shared/config"
	"github.com/argus-labs/world-cli/v2/internal/app/world-cli/shared/docker"
	"github.com/argus-labs/world-cli/v2/internal/app/world-cli/shared/docker/service"
	"github.com/argus-labs/world-cli/v2/internal/pkg/logger"
	"github.com/argus-labs/world-cli/v2/internal/pkg/printer"
)

func (h *Handler) Stop(ctx context.Context, flags models.StopEVMFlags) error {
	cfg, err := config.GetConfig(&flags.Config)
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

	err = dockerClient.Stop(ctx, service.EVM, service.CelestiaDevNet)
	if err != nil {
		return err
	}

	printer.Infoln("EVM successfully stopped")
	return nil
}
