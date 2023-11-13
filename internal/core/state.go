package core

import "time"

type State struct {
	Fx int `json:"fx" yaml:"fx,omitempty"`
	Sx int `json:"SX" yaml:"SX,omitempty"`
	Ix int `json:"IX" yaml:"IX,omitempty"`

	O1 bool `json:"O1,omitempty" yaml:"O1,omitempty"`
	O2 bool `json:"O2,omitempty" yaml:"O2,omitempty"`
	O3 bool `json:"O3,omitempty" yaml:"O3,omitempty"`
	C1 int  `json:"C1,omitempty" yaml:"C1,omitempty"`
	C2 int  `json:"C2,omitempty" yaml:"C2,omitempty"`
	C3 int  `json:"C3,omitempty" yaml:"C3,omitempty"`

	Color1 Color `json:"color1,omitempty" yaml:"color1,omitempty"`
	Color2 Color `json:"color2,omitempty" yaml:"color2,omitempty"`
	Color3 Color `json:"color3,omitempty" yaml:"color3,omitempty"`

	Segments []Segment `json:"segments"`
}

type Action struct {
	State    State
	Duration time.Duration
}
