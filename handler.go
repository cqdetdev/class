package class

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/armour"
	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/df-mc/dragonfly/server/player"
)

// Handler is the armour handler used to handle when a player changes class.
type Handler struct {
	inventory.NopHandler
	p *player.Player
}

// NewHandler returns a new *Handler.
func NewHandler(p *player.Player) *Handler { return &Handler{p: p} }

// Name ...
func (*Handler) Name() string { return "Class Handler" }

// HandlePlace handles when a piece of armour is placed in the armour inventory.
func (h *Handler) HandlePlace(ctx *event.Context, slot int, i item.Stack) {
	p := h.p
	if _, ok := i.Item().(armour.Armour); ok {
		fakeContainer := *h.p.Armour()
		fakeContainer.Inventory().AddItem(i)
		for _, class := range registeredClasses {
			if InClass(fakeContainer, class) {
				SetClass(p, class)
			}
		}
	}
}

// HandleTake handles when a piece of armour is removed from the armour inventory.
func (h *Handler) HandleTake(ctx *event.Context, slot int, i item.Stack) {
	p := h.p
	if _, ok := i.Item().(armour.Armour); ok {
		RemoveClass(p)
	}
}

// HandleDrop handles when a piece of armour is removed by drop cause from the armour inventory.
func (h *Handler) HandleDrop(ctx *event.Context, slot int, i item.Stack) {
	p := h.p
	if _, ok := i.Item().(armour.Armour); ok {
		RemoveClass(p)
	}
}
