package controller

import (
	"api/internal/core"
	"context"
	"log"
	"time"
)

func (c *Controller) run() {

	for c.keepRunning {
		c.idle()

		select {
		case action := <-c.Queue:

			c.runAction(action)
		case <-c.configChanged:
			log.Println("Config changed")
			continue
		}
	}
}

func (c *Controller) runAction(action core.Action) {
	c.setState(action.State)
	time.Sleep(action.Duration)
}

func (c *Controller) setState(state core.State) {
	for _, client := range c.clients {
		client.SetState(state)
	}
}

func (c *Controller) idle() {
	idleState, err := c.repo.GetIdleState(context.Background())
	if err != nil {
		panic(err)
	}
	c.setState(*idleState)
}
