package core

import "context"

type PresetId int

type Reaction struct {
	Amount   int      `json:"amount" yaml:"amount"`
	PresetId PresetId `json:"presetId" yaml:"presetId"`
	Duration int      `json:"duration" yaml:"duration"`
}

type AlertConfig interface {
	ReSubAlertRepo
	GiftSubAlertRepo
	BitAlertRepo
	DonationAlertRepo
}

type GlobalConfigRepo interface {
	GetIdleState(ctx context.Context) (PresetId, error)
	SetIdleState(ctx context.Context, presetId PresetId) error

	GetChannel(ctx context.Context) (string, error)
	SetChannel(ctx context.Context, channel string) error

	GetStreamlabsToken(ctx context.Context) (string, error)
	SetStreamlabsToken(ctx context.Context, token string) error

	GetDeviceAddr(ctx context.Context) (string, error)
	SetDeviceAddr(ctx context.Context, token string) error
}

type BitAlertRepo interface {
	GetAllBitAlerts(ctx context.Context) (map[string]Reaction, error)
	GetBitAlertById(ctx context.Context, id string) (*Reaction, error)
	SetBitAlertById(ctx context.Context, id string, r Reaction) error
	DeleteBitAlertById(ctx context.Context, id string) error
}

type DonationAlertRepo interface {
	GetAllDonationAlerts(ctx context.Context) (map[string]Reaction, error)
	GetDonationAlertById(ctx context.Context, id string) (*Reaction, error)
	SetDonationAlertById(ctx context.Context, id string, r Reaction) error
	DeleteDonationAlertById(ctx context.Context, id string) error
}

type ReSubAlertRepo interface {
	GetAllReSubs(ctx context.Context) (map[string]Reaction, error)
	GetResubById(ctx context.Context, id string) (*Reaction, error)
	SetResubById(ctx context.Context, id string, reaction Reaction) error
	DeleteResubById(ctx context.Context, id string) error
}

type GiftSubAlertRepo interface {
	GetAllGifts(ctx context.Context) (map[string]Reaction, error)
	GetGiftById(ctx context.Context, id string) (*Reaction, error)
	SetGiftById(ctx context.Context, id string, reaction Reaction) error
	DeleteGiftById(ctx context.Context, id string) error
}
