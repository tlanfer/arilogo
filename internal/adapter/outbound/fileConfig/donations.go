package fileConfig

import (
	"api/internal/core"
	"context"
	"fmt"
)

var _ core.DonationAlertRepo = (*Repo)(nil)

func (r Repo) GetAllDonationAlerts(ctx context.Context) (map[string]core.Reaction, error) {
	loaded, err := r.load(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	if loaded.DonationAlerts == nil {
		return map[string]core.Reaction{}, nil
	}

	return loaded.DonationAlerts, nil
}

func (r Repo) GetDonationAlertById(ctx context.Context, id string) (*core.Reaction, error) {
	loaded, err := r.load(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	if loaded.DonationAlerts == nil {
		return nil, nil
	}

	if reaction, ok := loaded.DonationAlerts[id]; ok {
		return &reaction, nil
	} else {
		return nil, nil
	}
}

func (r Repo) SetDonationAlertById(ctx context.Context, id string, reaction core.Reaction) error {
	loaded, err := r.load(ctx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if loaded.DonationAlerts == nil {
		loaded.DonationAlerts = map[string]core.Reaction{}
	}

	loaded.DonationAlerts[id] = reaction

	if err := r.save(ctx, loaded); err != nil {
		return fmt.Errorf("failed to save sub alert: %w", err)
	}

	return nil
}

func (r Repo) DeleteDonationAlertById(ctx context.Context, id string) error {
	loaded, err := r.load(ctx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if loaded.DonationAlerts == nil {
		loaded.DonationAlerts = map[string]core.Reaction{}
	}

	delete(loaded.DonationAlerts, id)

	if err := r.save(ctx, loaded); err != nil {
		return fmt.Errorf("failed to save sub alert: %w", err)
	}

	return nil
}
