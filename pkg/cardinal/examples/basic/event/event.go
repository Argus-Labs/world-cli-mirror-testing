package event

import "github.com/argus-labs/world-cli/v2/pkg/cardinal"

type PlayerDeath struct {
	cardinal.BaseEvent
	Nickname string
}

func (PlayerDeath) Name() string {
	return "player-death"
}

func (PlayerDeath) Group() string {
	return "rampage"
}

type NewPlayer struct {
	cardinal.BaseEvent
	Nickname string
}

func (NewPlayer) Name() string {
	return "new-player"
}
