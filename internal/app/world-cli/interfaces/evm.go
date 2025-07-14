package interfaces

import (
	"context"

	"github.com/argus-labs/go-ecs/internal/app/world-cli/models"
)

type EVMHandler interface {
	Start(ctx context.Context, flags models.StartEVMFlags) error
	Stop(ctx context.Context, flags models.StopEVMFlags) error
}
