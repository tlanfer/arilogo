package fileConfig

import (
	"api/internal/core"
	"context"
	"fmt"
	"time"
)

type ReSub struct {
	Months   int           `yaml:"months"`
	State    core.State    `yaml:"state"`
	Duration time.Duration `yaml:"duration"`
}

func (r Repo) GetAllReSubs(ctx context.Context) (core.ReSubs, error) {
	loaded, err := r.load(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	var resubs []core.ReSub

	for id, sub := range loaded.Subs {
		resubs = append(resubs, core.ReSub{
			Id:       id,
			Months:   sub.Months,
			State:    sub.State,
			Duration: sub.Duration,
		})
	}

	return resubs, nil
}

func (r Repo) SaveSubAlert(ctx context.Context, rs core.ReSub) error {
	loaded, err := r.load(ctx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if loaded.Subs == nil {
		loaded.Subs = map[string]ReSub{}
	}

	loaded.Subs[rs.Id] = ReSub{
		State:    rs.State,
		Months:   rs.Months,
		Duration: rs.Duration,
	}

	if err := r.save(ctx, loaded); err != nil {
		return fmt.Errorf("failed to save sub alert: %w", err)
	}

	return nil
}

func (r Repo) DeleteSubAlert(ctx context.Context, id string) error {
	loaded, err := r.load(ctx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	if loaded.Subs == nil {
		return nil
	}

	delete(loaded.Subs, id)

	if err := r.save(ctx, loaded); err != nil {
		return fmt.Errorf("failed to delete sub alert: %w", err)
	}

	return nil
}
