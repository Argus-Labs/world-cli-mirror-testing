package main

import (
	"github.com/argus-labs/world-cli/v2/pkg/cardinal"
	"github.com/argus-labs/world-cli/v2/pkg/cardinal/examples/chat/system"
)

func main() {
	world := cardinal.NewWorld()

	cardinal.RegisterSystem(world, system.UserChatSystem)

	world.StartGame()
}
