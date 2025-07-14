package organization

import (
	"github.com/argus-labs/go-ecs/internal/app/world-cli/clients/api"
	"github.com/argus-labs/go-ecs/internal/app/world-cli/interfaces"
	"github.com/argus-labs/go-ecs/internal/app/world-cli/services/config"
	"github.com/argus-labs/go-ecs/internal/app/world-cli/services/input"
)

// Interface guard.
var _ interfaces.OrganizationHandler = (*Handler)(nil)

type Handler struct {
	projectHandler interfaces.ProjectHandler
	inputService   input.ServiceInterface
	apiClient      api.ClientInterface
	configService  config.ServiceInterface
}

func NewHandler(
	projectHandler interfaces.ProjectHandler,
	inputService input.ServiceInterface,
	apiClient api.ClientInterface,
	configService config.ServiceInterface,
) interfaces.OrganizationHandler {
	return &Handler{
		projectHandler: projectHandler,
		inputService:   inputService,
		apiClient:      apiClient,
		configService:  configService,
	}
}
