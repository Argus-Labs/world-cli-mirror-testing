package cloud

import (
	"context"

	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/app/world-cli/interfaces"
	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/app/world-cli/models"
	"github.com/stretchr/testify/mock"
)

// Interface guard.
var _ interfaces.CloudHandler = (*MockHandler)(nil)

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Deployment(
	ctx context.Context,
	organizationID string,
	project models.Project,
	deployType string,
) error {
	args := m.Called(ctx, organizationID, project, deployType)
	return args.Error(0)
}

func (m *MockHandler) Status(ctx context.Context, organization models.Organization, project models.Project) error {
	args := m.Called(ctx, organization, project)
	return args.Error(0)
}

func (m *MockHandler) TailLogs(ctx context.Context, region string, env string) error {
	args := m.Called(ctx, region, env)
	return args.Error(0)
}
