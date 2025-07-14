package user

import (
	"github.com/argus-labs/go-ecs/internal/app/world-cli/clients/api"
	"github.com/argus-labs/go-ecs/internal/app/world-cli/interfaces"
	"github.com/argus-labs/go-ecs/internal/app/world-cli/services/input"
)

// Interface guard.
var _ interfaces.UserHandler = (*Handler)(nil)

type Handler struct {
	apiClient    api.ClientInterface
	inputService input.ServiceInterface
}

func NewHandler(apiClient api.ClientInterface, inputService input.ServiceInterface) interfaces.UserHandler {
	return &Handler{
		apiClient:    apiClient,
		inputService: inputService,
	}
}
