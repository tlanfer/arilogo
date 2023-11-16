package fileConfig

import (
	"api/internal/core"
	"context"
	"fmt"
)

var _ core.BitAlertRepo = (*Repo)(nil)

func (r Repo) GetAllBitAlerts(ctx context.Context) (map[string]core.Reaction, error) {
	loaded, err := r.load(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	if loaded.BitAlerts == nil {
		return map[string]core.Reaction{}, nil
	}

	return loaded.BitAlerts, nil
}

func (r Repo) GetBitAlertById(ctx context.Context, id string) (*core.Reaction, error) {
	loaded, err := r.load(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	if loaded.BitAlerts == nil {
		return nil, nil
	}

	if reaction, ok := loaded.BitAlerts[id]; ok {
		return &reaction, nil
	} else {
		return nil, nil
	}
}

func (r Repo) SetBitAlertById(ctx context.Context, id string, reaction core.Reaction) error {
	loaded, err := r.load(ctx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if loaded.BitAlerts == nil {
		loaded.BitAlerts = map[string]core.Reaction{}
	}

	loaded.BitAlerts[id] = reaction

	if err := r.save(ctx, loaded); err != nil {
		return fmt.Errorf("failed to save sub alert: %w", err)
	}

	return nil
}

func (r Repo) DeleteBitAlertById(ctx context.Context, id string) error {
	loaded, err := r.load(ctx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if loaded.BitAlerts == nil {
		loaded.BitAlerts = map[string]core.Reaction{}
	}

	delete(loaded.BitAlerts, id)

	if err := r.save(ctx, loaded); err != nil {
		return fmt.Errorf("failed to save sub alert: %w", err)
	}

	return nil
}
