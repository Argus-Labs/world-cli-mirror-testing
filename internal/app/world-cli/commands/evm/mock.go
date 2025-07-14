package evm

import (
	"context"

	"github.com/argus-labs/go-ecs/internal/app/world-cli/interfaces"
	"github.com/argus-labs/go-ecs/internal/app/world-cli/models"
	"github.com/stretchr/testify/mock"
)

var _ interfaces.EVMHandler = (*MockHandler)(nil)

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Start(ctx context.Context, flags models.StartEVMFlags) error {
	args := m.Called(ctx, flags)
	return args.Error(0)
}

func (m *MockHandler) Stop(ctx context.Context, flags models.StopEVMFlags) error {
	args := m.Called(ctx, flags)
	return args.Error(0)
}
