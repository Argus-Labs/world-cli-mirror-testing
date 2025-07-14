package root

import (
	"github.com/argus-labs/go-ecs/internal/app/world-cli/clients/api"
	"github.com/argus-labs/go-ecs/internal/app/world-cli/clients/browser"
	"github.com/argus-labs/go-ecs/internal/app/world-cli/interfaces"
	"github.com/argus-labs/go-ecs/internal/app/world-cli/services/config"
)

// Interface guard.
var _ interfaces.RootHandler = (*Handler)(nil)

type Handler struct {
	AppVersion      string
	configService   config.ServiceInterface
	apiClient       api.ClientInterface
	setupController interfaces.CommandSetupController
	browserClient   browser.ClientInterface
}

func NewHandler(
	appVersion string,
	configService config.ServiceInterface,
	apiClient api.ClientInterface,
	setupController interfaces.CommandSetupController,
	browserClient browser.ClientInterface,
) *Handler {
	return &Handler{
		AppVersion:      appVersion,
		configService:   configService,
		apiClient:       apiClient,
		setupController: setupController,
		browserClient:   browserClient,
	}
}
