package class

import (
	"sync"

	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
)

var h Handler = NopHandler{}
var hMu sync.RWMutex

type Handler interface {
	HandleSetClass(ctx *event.Context, player *player.Player, class Class)
	HandleRemoveClass(ctx *event.Context, player *player.Player, oldClass Class)
}

type NopHandler struct{}

func (NopHandler) HandleSetClass(ctx *event.Context, player *player.Player, class Class)       {}
func (NopHandler) HandleRemoveClass(ctx *event.Context, player *player.Player, oldClass Class) {}

func handler() Handler {
	hMu.RLock()
	defer hMu.RUnlock()
	return h
}
func SetClassHandler(hl Handler) {
	if hl != nil {
		hMu.Lock()
		defer hMu.Unlock()
		h = hl
	}
}
