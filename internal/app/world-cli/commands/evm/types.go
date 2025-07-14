package evm

import "github.com/argus-labs/go-ecs/internal/app/world-cli/interfaces"

var _ interfaces.EVMHandler = (*Handler)(nil)

type Handler struct {
}

func NewHandler() interfaces.EVMHandler {
	return &Handler{}
}
