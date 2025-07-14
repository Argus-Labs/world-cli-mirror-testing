package system

import (
	"github.com/argus-labs/world-cli/v2/pkg/cardinal/examples/chat/component"
	"github.com/argus-labs/world-cli/v2/pkg/ecs"
)

type ChatSearch = ecs.Exact[struct {
	UserTag ecs.Ref[component.UserTag]
	Chat    ecs.Ref[component.Chat]
}]
