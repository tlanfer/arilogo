package core

import (
	twitchchat "api/internal/adapter/inbound/twitch"
	"context"
	"github.com/tlanfer/go-streamlabs"
	"log"
)

func NewEventManager(c *Controller, config AlertConfig, tc twitchchat.Twitch, sl streamlabs.Streamlabs) Manager {
	return Manager{
		tc:     tc,
		sl:     sl,
		c:      c,
		config: config,
	}
}

type Manager struct {
	config AlertConfig

	sl streamlabs.Streamlabs
	tc twitchchat.Twitch

	bits   chan twitchchat.Bits
	gift   chan twitchchat.SubGift
	resubs chan twitchchat.ReSub

	c *Controller
}

func (m Manager) Run() {
	bits := make(chan twitchchat.Bits)
	gift := make(chan twitchchat.SubGift)
	resubs := make(chan twitchchat.ReSub)
	donos := make(chan streamlabs.Donation)

	m.tc.OnBits(bits)
	m.tc.OnSubGift(gift)
	m.tc.OnReSub(resubs)
	m.sl.OnDonation(donos)

	for {
		var reaction *Reaction
		select {
		case alert := <-donos:
			reaction = m.getReaction(alert.Amount, m.config.GetAllDonationAlerts)
		case alert := <-bits:
			reaction = m.getReaction(alert.Amount, m.config.GetAllBitAlerts)
		case alert := <-gift:
			reaction = m.getReaction(alert.Amount, m.config.GetAllGifts)
		case alert := <-resubs:
			reaction = m.getReaction(alert.Months, m.config.GetAllReSubs)
		}

		m.c.AddToQueue(*reaction)
	}
}

func (m Manager) getReaction(amount int, getter func(ctx context.Context) (map[string]Reaction, error)) *Reaction {
	all, err := getter(context.Background())
	if err != nil {
		log.Printf("failed to load bit alerts: %v", err)
		return nil
	}
	return m.findReaction(amount, all)
}

func (m Manager) findReaction(amount int, all map[string]Reaction) *Reaction {
	var reaction *Reaction
	for _, r := range all {
		if r.Amount > amount {
			continue
		}

		if reaction == nil {
			reaction = &r
			continue
		}

		if reaction.Amount < r.Amount {
			reaction = &r
		}
	}

	return reaction
}
