package system

import (
	"github.com/argus-labs/world-cli/v2/pkg/cardinal"
	"github.com/argus-labs/world-cli/v2/pkg/cardinal/examples/basic/component"
	"github.com/argus-labs/world-cli/v2/pkg/ecs"
)

type RegenSystemState struct {
	cardinal.BaseSystemState
	ecs.Contains[struct {
		ecs.Ref[component.Health]
	}]
}

func RegenSystem(state *RegenSystemState) error {
	for _, health := range state.Iter() { // Another shorthand
		health.Set(component.Health{HP: health.Get().HP + 10})
	}
	return nil
}
