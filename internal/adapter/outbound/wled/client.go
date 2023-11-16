package wled

import (
	"api/internal/core"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func New(addr string) core.Light {
	return &Client{
		addr: addr,
		client: http.Client{
			Timeout: 2 * time.Second,
		},
	}
}

type Client struct {
	addr   string
	client http.Client
}

func (c *Client) SetAddr(addr string) {
	c.addr = addr
}

func (c *Client) get(path string, dto any) error {
	resp, err := c.client.Get(fmt.Sprintf("http://%v/json/%v", c.addr, path))
	if err != nil {
		return fmt.Errorf("failed to get %v: %w", path, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if err := json.NewDecoder(resp.Body).Decode(dto); err != nil {
		return fmt.Errorf("failed to parse %v: %w", path, err)
	}

	return nil
}

func (c *Client) post(path string, in, out any) error {

	buf, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf("failed to serialize %v: %w", in, err)
	}

	resp, err := c.client.Post(fmt.Sprintf("http://%v/json/%v", c.addr, path), "application/json", bytes.NewReader(buf))
	if err != nil {
		return fmt.Errorf("failed to get %v: %w", path, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if out != nil {
		if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
			return fmt.Errorf("failed to parse %v: %w", path, err)
		}
	}

	return nil
}
