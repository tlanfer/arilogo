package wled

import (
	"api/internal/core"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func (c *Client) Presets() map[core.PresetId]string {

	resp, err := c.client.Get(fmt.Sprintf("http://%v/presets.json", c.addr))
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	dto := map[string]Preset{}
	if err = json.NewDecoder(resp.Body).Decode(&dto); err != nil {
		log.Println("failed to query presets")
	}
	presets := map[core.PresetId]string{}

	for id, preset := range dto {
		i, _ := strconv.Atoi(id)
		presets[core.PresetId(i)] = preset.Name
	}

	return presets
}

func (c *Client) SetPreset(presetId core.PresetId) error {
	if err := c.post("state", State{Preset: presetId}, nil); err != nil {
		return fmt.Errorf("failed to apply preset: %w", err)
	}
	return nil
}

type State struct {
	Preset core.PresetId `json:"ps"`
}

type Preset struct {
	Name string `json:"n"`
}
