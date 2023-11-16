package core

import (
	"context"
	"log"
	"time"
)

type Light interface {
	Presets() map[PresetId]string
	SetPreset(id PresetId) error
	SetAddr(addr string)
}

func NewController(repo GlobalConfigRepo, clients ...Light) *Controller {
	c := Controller{
		repo:          repo,
		clients:       clients,
		keepRunning:   true,
		Queue:         make(chan Reaction, 1000),
		configChanged: make(chan any),
	}
	return &c
}

type Controller struct {
	clients []Light

	Queue chan Reaction

	keepRunning   bool
	repo          GlobalConfigRepo
	configChanged chan any
}

func (c *Controller) ConfigChanged() {
	c.configChanged <- ""
}

func (c *Controller) AddToQueue(action Reaction) {
	c.Queue <- action
}

func (c *Controller) Run() {

	for c.keepRunning {
		c.idle()

		select {
		case action := <-c.Queue:
			c.runAction(action)
		case <-c.configChanged:
			continue
		}
	}
}

func (c *Controller) runAction(action Reaction) {
	c.setPreset(action.PresetId)
	time.Sleep(time.Duration(action.Duration) * time.Millisecond)
}

func (c *Controller) setPreset(presetId PresetId) {
	for _, client := range c.clients {
		if err := client.SetPreset(presetId); err != nil {
			log.Printf("Client %v failed: %v", client, err)
		}
	}
}

func (c *Controller) idle() {
	presetId, err := c.repo.GetIdleState(context.Background())
	if err != nil {
		return
	}
	c.setPreset(presetId)
}
