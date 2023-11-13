package controller

import (
	"api/internal/adapter/outbound/wled"
	"api/internal/core"
)

func New(repo core.GlobalConfigRepo, clients ...wled.Client) *Controller {
	c := Controller{
		repo:          repo,
		clients:       clients,
		keepRunning:   true,
		Queue:         make(chan core.Action, 1000),
		configChanged: make(chan any),
	}
	go c.run()
	return &c
}

type Controller struct {
	clients []wled.Client

	Queue chan core.Action

	keepRunning   bool
	repo          core.GlobalConfigRepo
	configChanged chan any
}

func (c *Controller) ConfigChanged() {
	c.configChanged <- ""
}
