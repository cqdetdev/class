package main

import (
	"time"

	"github.com/RestartFU/tickerFunc"
	"github.com/df-HCF/class"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/armour"
	"github.com/df-mc/dragonfly/server/player"
)

func main() {
	c := server.DefaultConfig()
	c.Players.SaveData = false
	s := server.New(&c, nil)
	s.Start()
	for {
		p, err := s.Accept()
		if err != nil {
			return
		}

		p.Armour().Inventory().Handle(class.NewHandler(p))
		c := &testClass{}
		class.Register(c)
		p.Inventory().AddItem(item.NewStack(item.Sugar{}, 64))
		p.Inventory().AddItem(item.NewStack(item.SpiderEye{}, 64))

		p.Inventory().AddItem(item.NewStack(item.Helmet{Tier: armour.TierGold}, 1))
		p.Inventory().AddItem(item.NewStack(item.Chestplate{Tier: armour.TierGold}, 1))
		p.Inventory().AddItem(item.NewStack(item.Leggings{Tier: armour.TierGold}, 1))
		p.Inventory().AddItem(item.NewStack(item.Boots{Tier: armour.TierGold}, 1))
	}
}

type testClass struct {
	tickers []*tickerFunc.Ticker
}

func (testClass) New(p *player.Player) class.Class {
	ticker := tickerFunc.NewTicker(5*time.Second, func() {
		p.Message("tickerFunc")
	})
	c := &testClass{tickers: []*tickerFunc.Ticker{ticker}}
	return c
}

func (*testClass) Armour() class.Armour {
	return class.Armour{
		armour.TierGold,
		armour.TierGold,
		armour.TierGold,
		armour.TierGold,
	}
}

func (*testClass) Effects() []effect.Effect {
	return []effect.Effect{
		effect.New(effect.Speed{}, 2, 15*time.Hour),
	}
}

func (c *testClass) Tickers(p *player.Player) []*tickerFunc.Ticker {
	return c.tickers
}

func (*testClass) Handler(p *player.Player) player.Handler { return nil }
