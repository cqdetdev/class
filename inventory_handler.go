package class

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/armour"
	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/df-mc/dragonfly/server/player"
)

// Handler is the armour handler used to handle when a player changes class.
type InventoryHandler struct {
	inventory.NopHandler
	p *player.Player
}

// NewHandler returns a new *Handler.
func NewHandler(p *player.Player) *InventoryHandler { return &InventoryHandler{p: p} }

// Name ...
func (*InventoryHandler) Name() string { return "Class Handler" }

// HandlePlace handles when a piece of armour is placed in the armour inventory.
func (hl *InventoryHandler) HandlePlace(ctx *event.Context, slot int, i item.Stack) {
	p := hl.p
	if _, ok := i.Item().(armour.Armour); ok {
		fakeContainer := *p.Armour()
		fakeContainer.Inventory().AddItem(i)
		for _, class := range registeredClasses {
			if InClass(fakeContainer, class) {
				handler().HandleSetClass(ctx, p, class)
				if !ctx.Cancelled() {
					SetClass(p, class)
				}
			}
		}
	}
}

// HandleTake handles when a piece of armour is removed from the armour inventory.
func (hl *InventoryHandler) HandleTake(ctx *event.Context, slot int, i item.Stack) {
	handleRemove(ctx, hl)
}

// HandleDrop handles when a piece of armour is removed by drop cause from the armour inventory.
func (hl *InventoryHandler) HandleDrop(ctx *event.Context, slot int, i item.Stack) {
	handleRemove(ctx, hl)
}

func handleRemove(ctx *event.Context, hl *InventoryHandler) {
	p := hl.p
	class, _ := PlayerClass(p)
	handler().HandleRemoveClass(ctx, p, class)
	if !ctx.Cancelled() {
		RemoveClass(p)
	}
}
