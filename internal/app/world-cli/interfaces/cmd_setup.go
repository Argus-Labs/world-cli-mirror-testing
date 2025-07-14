package interfaces

import (
	"context"

	"github.com/argus-labs/world-cli/v2/internal/app/world-cli/models"
)

// CommandSetupController defines the interface for command setup operations.
type CommandSetupController interface {
	// SetupCommandState performs the setup flow for commands and returns the established state.
	// Handles login, organization, and project setup based on requirements.
	SetupCommandState(ctx context.Context, req models.SetupRequest) (models.CommandState, error)
}
