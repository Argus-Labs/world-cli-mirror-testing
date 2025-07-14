package main

import (
	"github.com/argus-labs/go-ecs/pkg/cardinal"
	"github.com/argus-labs/go-ecs/pkg/cardinal/examples/chat/system"
)

func main() {
	world := cardinal.NewWorld()

	cardinal.RegisterSystem(world, system.UserChatSystem)

	world.StartGame()
}
