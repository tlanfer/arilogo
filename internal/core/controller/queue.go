package controller

import (
	"api/internal/core"
)

func (c *Controller) AddToQueue(action core.Action) {
	c.Queue <- action
}
