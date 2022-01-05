package class

import "github.com/df-mc/dragonfly/server/item/armour"

// Armour contains 4 armour tiers which defines the armour required for a class
type Armour struct {
	Helmet    armour.Tier
	Chestlate armour.Tier
	Leggings  armour.Tier
	Boots     armour.Tier
}
