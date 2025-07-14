package organization

import (
	"context"

	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/app/world-cli/interfaces"
	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/app/world-cli/models"
	"github.com/stretchr/testify/mock"
)

// Interface guard.
var _ interfaces.OrganizationHandler = (*MockHandler)(nil)

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Create(ctx context.Context, flags models.CreateOrganizationFlags) (models.Organization, error) {
	args := m.Called(ctx, flags)
	org, ok := args.Get(0).(models.Organization)
	if !ok {
		return models.Organization{}, args.Error(1)
	}

	return org, args.Error(1)
}

func (m *MockHandler) Switch(ctx context.Context, flags models.SwitchOrganizationFlags) (models.Organization, error) {
	args := m.Called(ctx, flags)
	org, ok := args.Get(0).(models.Organization)
	if !ok {
		return models.Organization{}, args.Error(1)
	}

	return org, args.Error(1)
}

func (m *MockHandler) MembersList(ctx context.Context, org models.Organization, flags models.MembersListFlags) error {
	args := m.Called(ctx, org, flags)
	return args.Error(0)
}

func (m *MockHandler) PromptForSwitch(ctx context.Context, orgs []models.Organization, enableCreation bool,
) (models.Organization, error) {
	args := m.Called(ctx, orgs, enableCreation)
	org, ok := args.Get(0).(models.Organization)
	if !ok {
		return models.Organization{}, args.Error(1)
	}

	return org, args.Error(1)
}

func (m *MockHandler) PrintNoOrganizations() {
	m.Called()
}
