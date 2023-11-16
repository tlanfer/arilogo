package fileConfig

import (
	"api/internal/core"
	"context"
	"fmt"
	"io"
	"strings"
)

func (r Repo) GetIdleState(ctx context.Context) (core.PresetId, error) {
	dto, err := r.load(ctx)
	if err != nil && err != io.EOF {
		return -1, fmt.Errorf("failed to load config: %w", err)
	}

	if dto == nil || dto.IdlePreset == -1 {
		err := r.SetIdleState(ctx, 0)
		if err != nil {
			return -1, fmt.Errorf("failed to save default config: %w", err)
		}
		return 0, nil
	}

	return dto.IdlePreset, nil
}

func (r Repo) SetIdleState(ctx context.Context, preset core.PresetId) error {
	return r.modify(ctx, func(dto *Dto) error {
		dto.IdlePreset = preset
		return nil
	})
}

func (r Repo) GetChannel(ctx context.Context) (string, error) {
	dto, err := r.load(ctx)
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	if dto == nil || dto.TwitchChannel == "" {
		return "", nil
	}

	return dto.TwitchChannel, nil
}

func (r Repo) SetChannel(ctx context.Context, channel string) error {
	return r.modify(ctx, func(dto *Dto) error {
		dto.TwitchChannel = channel
		return nil
	})
}

func (r Repo) GetStreamlabsToken(ctx context.Context) (string, error) {
	dto, err := r.load(ctx)
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	if dto == nil || dto.StreamlabsToken == "" {
		return "", nil
	}

	return dto.StreamlabsToken, nil
}
func (r Repo) SetStreamlabsToken(ctx context.Context, token string) error {
	return r.modify(ctx, func(dto *Dto) error {
		dto.StreamlabsToken = token
		return nil
	})
}

func (r Repo) GetDeviceAddr(ctx context.Context) (string, error) {
	dto, err := r.load(ctx)
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	if dto == nil || dto.DeviceAddr == "" {
		return "", nil
	}

	return strings.TrimSpace(dto.DeviceAddr), nil
}
func (r Repo) SetDeviceAddr(ctx context.Context, addr string) error {
	return r.modify(ctx, func(dto *Dto) error {
		dto.DeviceAddr = addr
		return nil
	})
}

func (r Repo) modify(ctx context.Context, mod func(dto *Dto) error) error {
	dto, err := r.load(ctx)
	if err != nil && err != io.EOF {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if dto == nil {
		dto = &Dto{}
	}

	if err := mod(dto); err != nil {
		return err
	}

	if err := r.save(ctx, dto); err != nil {
		return fmt.Errorf("failed to save: %w", err)
	}

	return nil
}
