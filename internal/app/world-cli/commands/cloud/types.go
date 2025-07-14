package cloud

import (
	"github.com/argus-labs/go-ecs/internal/app/world-cli/clients/api"
	"github.com/argus-labs/go-ecs/internal/app/world-cli/interfaces"
	"github.com/argus-labs/go-ecs/internal/app/world-cli/models"
	"github.com/argus-labs/go-ecs/internal/app/world-cli/services/config"
	"github.com/argus-labs/go-ecs/internal/app/world-cli/services/input"
)

// Interface guard.
var _ interfaces.CloudHandler = (*Handler)(nil)

type Handler struct {
	apiClient      api.ClientInterface
	configService  config.ServiceInterface
	projectHandler interfaces.ProjectHandler
	inputHandler   input.ServiceInterface
}

const (
	DeployStatusFailed  DeployStatus = "failed"
	DeployStatusCreated DeployStatus = "created"
	DeployStatusRemoved DeployStatus = "removed"

	DeployEnvPreview = "dev"
	DeployEnvLive    = "prod"
)

type HealthStatus string

const (
	HealthStatusHealthy   HealthStatus = "healthy"
	HealthStatusUnhealthy HealthStatus = "unhealthy"
	HealthStatusOffline   HealthStatus = "offline"
)

type DeployStatus string

type DeployInfo struct {
	DeployType    string
	DeployStatus  DeployStatus
	DeployDisplay string
}

type logParams struct {
	organization models.Organization
	project      models.Project
	region       string
	env          string
}

func NewHandler(
	apiClient api.ClientInterface,
	configService config.ServiceInterface,
	projectHandler interfaces.ProjectHandler,
	inputHandler input.ServiceInterface,
) *Handler {
	return &Handler{
		apiClient:      apiClient,
		configService:  configService,
		projectHandler: projectHandler,
		inputHandler:   inputHandler,
	}
}
