package main

import (
	"github.com/argus-labs/go-ecs/pkg/cardinal"
	"github.com/argus-labs/go-ecs/pkg/cardinal/examples/demo-game/system"
)

func main() {
	world := cardinal.NewWorld()

	cardinal.RegisterSystem(world, system.PlayerSetUpdater, cardinal.WithHook(cardinal.PreUpdate))
	cardinal.RegisterSystem(world, system.PlayerSpawnSystem)
	cardinal.RegisterSystem(world, system.MovePlayerSystem)
	cardinal.RegisterSystem(world, system.PlayerLeaveSystem)
	cardinal.RegisterSystem(world, system.OnlineStatusUpdater)

	world.StartGame()
}
