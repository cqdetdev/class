package class

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type EnergyUseItem interface {
	Energy() int
	Use(p *player.Player)
	Item() world.Item
}
