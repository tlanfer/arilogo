package twitchchat

import (
	"fmt"
	"github.com/gempir/go-twitch-irc/v3"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const Fdgt = "irc.fdgt.dev:6697"

type Twitch interface {
	Connect(channel string) error
	OnBits(chan<- Bits)
	OnSubGift(chan<- SubGift)
	OnReSub(chan<- ReSub)
}

type Opt func(chat *twitchChat)

func New(opts ...Opt) Twitch {
	t := &twitchChat{
		client: twitch.NewAnonymousClient(),
	}
	for _, opt := range opts {
		opt(t)
	}
	return t
}

func WithFdgt() Opt {
	return WithIrcServer(Fdgt)
}
func WithIrcServer(server string) Opt {
	return func(chat *twitchChat) {
		chat.client.IrcAddress = server
	}
}
func WithFaker(interval time.Duration) Opt {
	return func(chat *twitchChat) {
		go chat.FakerShaker(interval)
	}
}
func (t *twitchChat) FakerShaker(interval time.Duration) {

	if t.client.IrcAddress != Fdgt {
		log.Fatalf("Can only use fakes on %v!", Fdgt)
	}

	for {
		time.Sleep(interval)
		switch 4 {
		case 0:
			t.client.Say("channel", fmt.Sprintf("bits --bitscount %v Woohoo!", rand.Intn(500)))
		case 1:
			t.client.Say("channel", fmt.Sprintf("subgift --tier %v --username glEnd2", rand.Intn(3)+1))
		case 2:
			t.client.Say("channel", fmt.Sprintf("submysterygift --count %v --username zebiniasis", rand.Intn(15)+1))
		case 3:
			t.client.Say("channel", fmt.Sprintf("subscription --tier %v --username glEnd2", rand.Intn(3)+1))
		case 4:
			t.client.Say("channel", fmt.Sprintf("resubscription --tier %v --months %v", rand.Intn(3)+1, rand.Intn(20)+1))
		}

	}

}

type twitchChat struct {
	client         *twitch.Client
	currentChannel string

	chBits   chan<- Bits
	chGifts  chan<- SubGift
	chReSubs chan<- ReSub
}

func (t *twitchChat) OnBits(ch chan<- Bits) {
	t.chBits = ch
}

func (t *twitchChat) OnSubGift(ch chan<- SubGift) {
	t.chGifts = ch
}

func (t *twitchChat) OnReSub(ch chan<- ReSub) {
	t.chReSubs = ch
}

func (t *twitchChat) Connect(channel string) error {
	t.client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if message.Bits > 0 {
			t.chBits <- Bits{Amount: message.Bits}
		}
	})

	t.client.OnUserNoticeMessage(func(message twitch.UserNoticeMessage) {
		switch message.MsgID {
		case "submysterygift":
			count, _ := strconv.Atoi(message.MsgParams["msg-param-mass-gift-count"])
			t.chGifts <- SubGift{Amount: count}

		case "sub":
			fallthrough
		case "resub":
			months, _ := strconv.Atoi(message.MsgParams["msg-param-cumulative-months"])
			t.chReSubs <- ReSub{months}

		case "subgift":
			t.chGifts <- SubGift{Amount: 1}
		}
	})

	if t.currentChannel != channel {
		if t.currentChannel != "" {
			t.client.Depart(t.currentChannel)
		}
		t.client.Join(channel)
		t.currentChannel = channel
	}

	go func() {
		log.Println("Twitch connected:", t.currentChannel)
		_ = t.client.Connect()
	}()

	return nil
}
