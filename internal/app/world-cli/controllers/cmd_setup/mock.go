package cmdsetup

import (
	"context"

	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/app/world-cli/interfaces"
	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/app/world-cli/models"
	"github.com/stretchr/testify/mock"
)

// Ensure MockController implements the interface.
var _ interfaces.CommandSetupController = (*MockController)(nil)

// MockController is a mock implementation of CommandSetupController.
type MockController struct {
	mock.Mock
}

// SetupCommandState mocks the setup command.
func (m *MockController) SetupCommandState(
	ctx context.Context,
	req models.SetupRequest,
) (models.CommandState, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return models.CommandState{}, args.Error(1)
	}

	if state, ok := args.Get(0).(models.CommandState); ok {
		return state, args.Error(1)
	}
	return models.CommandState{}, args.Error(1)
}
