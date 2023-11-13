package core

import (
	"fmt"
)

type Color struct {
	R int `json:"R,omitempty"`
	G int `json:"G,omitempty"`
	B int `json:"B,omitempty"`
}

func (c Color) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"#%02x%02x%02x\"", c.R, c.G, c.B)), nil
}

func (c *Color) UnmarshalJSON(bytes []byte) error {
	if _, err := fmt.Sscanf(string(bytes[1:8]), "#%02x%02x%02x", &c.R, &c.G, &c.B); err != nil {
		return err
	}
	return nil
}
