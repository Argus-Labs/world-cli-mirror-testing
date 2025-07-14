package system

import (
	"github.com/argus-labs/go-ecs/pkg/cardinal"
	otherworld "github.com/argus-labs/go-ecs/pkg/cardinal/examples/basic/other_world"
)

// ExternalCommand should originate from another game shard.
type ExternalCommand struct {
	cardinal.BaseCommand
	Message string
}

func (ExternalCommand) Name() string {
	return "external"
}

func (ExternalCommand) Group() string {
	return "plugin"
}

type CallExternalCommand struct {
	cardinal.BaseCommand
	Message string
}

func (CallExternalCommand) Name() string {
	return "call-external"
}

type CallExternalSystemState struct {
	cardinal.BaseSystemState
	CallExternalCommands cardinal.WithCommand[CallExternalCommand]
}

func CallExternalSystem(state *CallExternalSystemState) error {
	for msg := range state.CallExternalCommands.Iter() {
		state.Logger().Info().Msg("Received call-external message")

		otherworld.Matchmaking.Send(&state.BaseSystemState, CreatePlayerCommand{
			Nickname: msg.Message,
		})
	}
	return nil
}
