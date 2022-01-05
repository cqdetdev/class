package class

import (
	"sync"

	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/df-mc/dragonfly/server/player"
)

func init() {
	playerClasses = make(map[*player.Player]Class)
}

var playerClasses map[*player.Player]Class
var playerMu sync.RWMutex

// PlayerClass returns the current class of the player if they have one
func PlayerClass(p *player.Player) (Class, bool) {
	playerMu.RLock()
	defer playerMu.RUnlock()
	class, ok := playerClasses[p]
	return class, ok && class != nil
}

// stopAllTickers ...
func stopAllTickers(p *player.Player, class Class) {
	if tClass, ok := class.(TickerClass); ok {
		for _, ticker := range tClass.Tickers(p) {
			ticker.Stop()
		}
	}
}

// removeAllEffects ...
func removeAllEffects(p *player.Player, class Class) {
	if eClass, ok := class.(EffectClass); ok {
		for _, e := range eClass.Effects() {
			p.RemoveEffect(e.Type())
		}
	}
}

// startAllTickers ...
func startAllTickers(p *player.Player, class Class) {
	if tClass, ok := class.(TickerClass); ok {
		for _, ticker := range tClass.Tickers(p) {
			go ticker.Start()
		}
	}
}

// giveAllEffects ...
func giveAllEffects(p *player.Player, class Class) {
	if eClass, ok := class.(EffectClass); ok {
		for _, e := range eClass.Effects() {
			p.AddEffect(e)
		}
	}
}

// removeClass ...
func removeClass(p *player.Player) {
	playerMu.Lock()
	defer playerMu.Unlock()
	delete(playerClasses, p)
}

// RemoveClass removes the current class of the player if they have one.
func RemoveClass(p *player.Player) {
	if playerClass, ok := PlayerClass(p); ok {
		p.Handle(nil)
		removeClass(p)

		stopAllTickers(p, playerClass)
		removeAllEffects(p, playerClass)
	}
}

// setClass ...
func setClass(p *player.Player, class Class) {
	playerMu.Lock()
	playerClasses[p] = class
	playerMu.Unlock()
}

// SetClass Sets the class provided to the player.
func SetClass(p *player.Player, class Class) {
	RemoveClass(p)
	if class != nil {
		class = class.New(p)
		setClass(p, class)

		p.Handle(class.Handler(p))
		startAllTickers(p, class)
		giveAllEffects(p, class)
	}
}

// InClass returns a bool of if the inventory.Armour matches with the class provided.
func InClass(inv inventory.Armour, class Class) bool {
	tiers := class.Armour()

	helmet := item.Helmet{Tier: tiers.Helmet}
	chestplate := item.Chestplate{Tier: tiers.Chestplate}
	leggings := item.Leggings{Tier: tiers.Leggings}
	boots := item.Boots{Tier: tiers.Boots}
	return inv.Helmet().Item() == helmet && inv.Chestplate().Item() == chestplate && inv.Leggings().Item() == leggings && inv.Boots().Item() == boots
}
