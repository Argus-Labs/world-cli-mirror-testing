package evm

import (
	"context"

	"github.com/argus-labs/go-ecs/internal/app/world-cli/common/config"
	"github.com/argus-labs/go-ecs/internal/app/world-cli/common/docker"
	"github.com/argus-labs/go-ecs/internal/app/world-cli/common/docker/service"
	"github.com/argus-labs/go-ecs/internal/app/world-cli/models"
	"github.com/argus-labs/go-ecs/internal/pkg/printer"
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
	defer dockerClient.Close()

	err = dockerClient.Stop(ctx, service.EVM, service.CelestiaDevNet)
	if err != nil {
		return err
	}

	printer.Infoln("EVM successfully stopped")
	return nil
}
