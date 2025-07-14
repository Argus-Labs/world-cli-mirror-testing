package cardinal

import "github.com/argus-labs/go-ecs/internal/app/world-cli/interfaces"

var _ interfaces.CardinalHandler = &Handler{}

type Handler struct {
}

func NewHandler() interfaces.CardinalHandler {
	return &Handler{}
}
