package root

import (
	"context"

	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/app/world-cli/interfaces"
	"github.com/stretchr/testify/mock"
)

// Interface guard.
var _ interfaces.RootHandler = (*MockHandler)(nil)

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Create(directory string) error {
	args := m.Called(directory)
	return args.Error(0)
}

func (m *MockHandler) Doctor() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockHandler) Version(check bool) error {
	args := m.Called(check)
	return args.Error(0)
}

func (m *MockHandler) Login(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockHandler) SetAppVersion(version string) {
	m.Called(version)
}
