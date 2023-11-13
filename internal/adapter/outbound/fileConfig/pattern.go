package fileConfig

import (
	"api/internal/core"
	"context"
	"fmt"
)

func (r Repo) GetAllPattern(ctx context.Context) (map[string]core.Pattern, error) {
	dto, err := r.load(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return dto.Patterns, nil
}

func (r Repo) SavePattern(ctx context.Context, pattern core.Pattern) error {
	loaded, err := r.load(ctx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if loaded.Patterns == nil {
		loaded.Patterns = map[string]core.Pattern{}
	}

	loaded.Patterns[pattern.Id] = pattern
	if err := r.save(ctx, loaded); err != nil {
		return fmt.Errorf("failed to save sub alert: %w", err)
	}

	return nil
}

func (r Repo) DeletePattern(ctx context.Context, id string) error {
	loaded, err := r.load(ctx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	if loaded.Patterns == nil {
		return nil
	}

	delete(loaded.Patterns, id)

	if err := r.save(ctx, loaded); err != nil {
		return fmt.Errorf("failed to delete sub alert: %w", err)
	}

	return nil
}
