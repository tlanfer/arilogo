package fileConfig

import (
	"api/internal/core"
	"context"
	"fmt"
)

var _ core.ReSubAlertRepo = (*Repo)(nil)
var _ core.GiftSubAlertRepo = (*Repo)(nil)

func (r Repo) GetAllReSubs(ctx context.Context) (map[string]core.Reaction, error) {
	loaded, err := r.load(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return loaded.Resubs, nil
}

func (r Repo) GetResubById(ctx context.Context, id string) (*core.Reaction, error) {
	loaded, err := r.load(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	if loaded.Resubs == nil {
		return nil, nil
	}

	if reaction, ok := loaded.Resubs[id]; ok {
		return &reaction, nil
	} else {
		return nil, nil
	}
}

func (r Repo) SetResubById(ctx context.Context, id string, reaction core.Reaction) error {
	loaded, err := r.load(ctx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if loaded.Resubs == nil {
		loaded.Resubs = map[string]core.Reaction{}
	}

	loaded.Resubs[id] = reaction

	if err := r.save(ctx, loaded); err != nil {
		return fmt.Errorf("failed to save sub alert: %w", err)
	}

	return nil
}

func (r Repo) DeleteResubById(ctx context.Context, id string) error {
	loaded, err := r.load(ctx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if loaded.Resubs == nil {
		loaded.Resubs = map[string]core.Reaction{}
	}

	delete(loaded.Resubs, id)

	if err := r.save(ctx, loaded); err != nil {
		return fmt.Errorf("failed to save sub alert: %w", err)
	}

	return nil
}

func (r Repo) GetAllGifts(ctx context.Context) (map[string]core.Reaction, error) {
	loaded, err := r.load(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return loaded.Gifts, nil
}

func (r Repo) GetGiftById(ctx context.Context, id string) (*core.Reaction, error) {
	loaded, err := r.load(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	if loaded.Gifts == nil {
		return nil, nil
	}

	if reaction, ok := loaded.Gifts[id]; ok {
		return &reaction, nil
	} else {
		return nil, nil
	}
}

func (r Repo) SetGiftById(ctx context.Context, id string, reaction core.Reaction) error {
	loaded, err := r.load(ctx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if loaded.Gifts == nil {
		loaded.Gifts = map[string]core.Reaction{}
	}

	loaded.Gifts[id] = reaction

	if err := r.save(ctx, loaded); err != nil {
		return fmt.Errorf("failed to save sub alert: %w", err)
	}

	return nil
}

func (r Repo) DeleteGiftById(ctx context.Context, id string) error {
	loaded, err := r.load(ctx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if loaded.Resubs == nil {
		loaded.Resubs = map[string]core.Reaction{}
	}

	delete(loaded.Gifts, id)

	if err := r.save(ctx, loaded); err != nil {
		return fmt.Errorf("failed to save sub alert: %w", err)
	}

	return nil
}
