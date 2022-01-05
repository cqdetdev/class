package class

import (
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type EnergyUseItem interface {
	Energy() int
	Use(p *player.Player)
	Item() world.Item
}
type EnergyEffectItem interface {
	Energy() int
	Effect() effect.Effect
	Item() world.Item
}
