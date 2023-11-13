package fileConfig

import (
	"api/internal/core"
	"context"
	"fmt"
)

func (r Repo) GetIdleState(ctx context.Context) (*core.State, error) {
	dto, err := r.load(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &dto.IdleState, nil
}

func (r Repo) SetIdleState(ctx context.Context, state core.State) error {
	dto, err := r.load(ctx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	dto.IdleState = state

	if err := r.save(ctx, dto); err != nil {
		return fmt.Errorf("failed to save: %w", err)
	}

	return nil
}
