package class

import (
	"sync"

	"github.com/RestartFU/tickerFunc"
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/player"
)

func init() {
	registeredClasses = make([]Class, 0)
}

var registeredClasses []Class
var classMu sync.RWMutex

// Register adds the class provided to registeredClasses
func Register(class Class) {
	classMu.Lock()
	defer classMu.Unlock()
	registeredClasses = append(registeredClasses, class)
}

// Class is the class interface
type Class interface {
	Armour() Armour
	Handler(*player.Player) player.Handler
	New(*player.Player) Class
}

// EffectClass is an interface which may be used for classe with effects
type EffectClass interface {
	Effects() []effect.Effect
}

// EffectClass is an interface which may be used for classe with tickers
type TickerClass interface {
	Tickers(p *player.Player) []*tickerFunc.Ticker
}
