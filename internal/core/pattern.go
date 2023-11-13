package core

import "context"

type PatternRepo interface {
	GetAllPattern(ctx context.Context) (map[string]Pattern, error)
	SavePattern(ctx context.Context, pattern Pattern) error
	DeletePattern(ctx context.Context, id string) error
}

type Pattern struct {
	Id       string    `json:"id" yaml:"id"`
	Name     string    `json:"name" yaml:"name"`
	Segments []Segment `json:"segments" yaml:"segments"`
}

type Segment struct {
	Start int `json:"start" yaml:"start"`
	End   int `json:"end" yaml:"end"`

	Fx int `json:"fx" yaml:"fx,omitempty"`
	Sx int `json:"sx" yaml:"sx,omitempty"`
	Ix int `json:"ix" yaml:"ix,omitempty"`

	O1 bool `json:"o1,omitempty" yaml:"o1,omitempty"`
	O2 bool `json:"o2,omitempty" yaml:"o2,omitempty"`
	O3 bool `json:"o3,omitempty" yaml:"o3,omitempty"`
	C1 int  `json:"c1,omitempty" yaml:"c1,omitempty"`
	C2 int  `json:"c2,omitempty" yaml:"c2,omitempty"`
	C3 int  `json:"c3,omitempty" yaml:"c3,omitempty"`

	Color1 Color `json:"color1,omitempty" yaml:"color1,omitempty"`
	Color2 Color `json:"color2,omitempty" yaml:"color2,omitempty"`
	Color3 Color `json:"color3,omitempty" yaml:"color3,omitempty"`
}
