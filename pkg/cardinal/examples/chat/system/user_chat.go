package system

import (
	"time"

	"github.com/argus-labs/go-ecs/pkg/cardinal"
	"github.com/argus-labs/go-ecs/pkg/cardinal/examples/chat/component"
	"github.com/argus-labs/go-ecs/pkg/cardinal/examples/chat/event"
)

type UserChatCommand struct {
	cardinal.BaseCommand
	ArgusAuthID   string `json:"argus_auth_id"`
	ArgusAuthName string `json:"argus_auth_name"`
	Message       string `json:"message"`
}

func (UserChatCommand) Name() string {
	return "user-chat"
}

type UserChatSystemState struct {
	cardinal.BaseSystemState
	UserChatCommands cardinal.WithCommand[UserChatCommand]
	UserChatEvent    cardinal.WithEvent[event.UserChat]
	ChatSearch       ChatSearch
}

func UserChatSystem(state *UserChatSystemState) error {
	for msg := range state.UserChatCommands.Iter() {
		timestamp := time.Now()

		id, err := state.ChatSearch.Create(
			component.UserTag{
				ArgusAuthID:   msg.ArgusAuthID,
				ArgusAuthName: msg.ArgusAuthName,
			},
			component.Chat{
				Message:   msg.Message,
				Timestamp: timestamp,
			},
		)

		if err != nil {
			// If we return the error, Cardinal will shutdown, so just log it.
			state.Logger().Error().Err(err).Msg("error creating entity")
			continue
		}

		state.Logger().Info().
			Uint32("entity", uint32(id)).
			Msgf("Created chat message %s (id: %s)", msg.Message, msg.ArgusAuthID)

		state.UserChatEvent.Emit(event.UserChat{
			ArgusAuthID:   msg.ArgusAuthID,
			ArgusAuthName: msg.ArgusAuthName,
			Message:       msg.Message,
			Timestamp:     timestamp,
		})
	}
	return nil
}
