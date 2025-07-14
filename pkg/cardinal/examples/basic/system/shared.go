package system

import (
	"github.com/argus-labs/go-ecs/pkg/cardinal/examples/basic/component"

	"github.com/argus-labs/go-ecs/pkg/ecs"
)

type PlayerSearch = ecs.Exact[struct {
	Tag    ecs.Ref[component.PlayerTag]
	Health ecs.Ref[component.Health]
}]

type GraveSearch = ecs.Exact[struct {
	Grave ecs.Ref[component.Gravestone]
}]
