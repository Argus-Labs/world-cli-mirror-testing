package evm

import "github.com/argus-labs/world-cli/v2/internal/app/world-cli/interfaces"

var _ interfaces.EVMHandler = (*Handler)(nil)

type Handler struct {
}

func NewHandler() interfaces.EVMHandler {
	return &Handler{}
}
