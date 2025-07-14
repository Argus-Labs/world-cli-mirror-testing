package cardinal

import (
	"context"

	"github.com/argus-labs/world-cli/v2/internal/app/world-cli/common/config"
	"github.com/argus-labs/world-cli/v2/internal/app/world-cli/common/docker"
	"github.com/argus-labs/world-cli/v2/internal/app/world-cli/common/docker/service"
	"github.com/argus-labs/world-cli/v2/internal/app/world-cli/models"
	"github.com/argus-labs/world-cli/v2/internal/pkg/printer"
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
	defer dockerClient.Close()

	err = dockerClient.Stop(ctx, service.Nakama, service.Cardinal,
		service.NakamaDB, service.Redis, service.Jaeger, service.Prometheus)
	if err != nil {
		return err
	}

	printer.Successln("Cardinal successfully stopped")

	return nil
}
