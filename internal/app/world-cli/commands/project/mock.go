package project

import (
	"context"

	"github.com/argus-labs/world-cli/v2/internal/app/world-cli/interfaces"
	"github.com/argus-labs/world-cli/v2/internal/app/world-cli/models"
	"github.com/stretchr/testify/mock"
)

// Interface guard.
var _ interfaces.ProjectHandler = (*MockHandler)(nil)

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Create(
	ctx context.Context,
	org models.Organization,
	flags models.CreateProjectFlags,
) (models.Project, error) {
	args := m.Called(ctx, org, flags)
	project, ok := args.Get(0).(models.Project)
	if !ok {
		return models.Project{}, args.Error(1)
	}

	return project, args.Error(1)
}

func (m *MockHandler) Switch(
	ctx context.Context,
	flags models.SwitchProjectFlags,
	org models.Organization,
	enableCreation bool,
) (models.Project, error) {
	args := m.Called(ctx, flags, org, enableCreation)
	project, ok := args.Get(0).(models.Project)
	if !ok {
		return models.Project{}, args.Error(1)
	}

	return project, args.Error(1)
}

func (m *MockHandler) HandleSwitch(ctx context.Context, org models.Organization) error {
	args := m.Called(ctx, org)
	return args.Error(0)
}

func (m *MockHandler) Update(
	ctx context.Context,
	project models.Project,
	org models.Organization,
	flags models.UpdateProjectFlags,
) error {
	args := m.Called(ctx, project, org, flags)
	return args.Error(0)
}

func (m *MockHandler) Delete(
	ctx context.Context,
	project models.Project,
) error {
	args := m.Called(ctx, project)
	return args.Error(0)
}

func (m *MockHandler) PreCreateUpdateValidation(printError bool) (string, string, error) {
	args := m.Called(printError)
	if args.Get(0) == nil {
		return "", "", args.Error(2)
	}

	slug, ok := args.Get(0).(string)
	if !ok {
		return "", "", args.Error(2)
	}

	projectName, ok := args.Get(1).(string)
	if !ok {
		return "", "", args.Error(2)
	}

	return slug, projectName, args.Error(2)
}

func (m *MockHandler) PrintNoProjectsInOrganization() {
	m.Called()
}
