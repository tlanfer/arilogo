package wled

import (
	"api/internal/core"
)

func (c *Client) GetState() core.State {
	s := State{}
	return core.State{
		Fx:     s.Segments[0].Fx,
		Sx:     s.Segments[0].Sx,
		Ix:     s.Segments[0].Ix,
		O1:     s.Segments[0].O1,
		O2:     s.Segments[0].O2,
		O3:     s.Segments[0].O3,
		C1:     s.Segments[0].C1,
		C2:     s.Segments[0].C2,
		C3:     s.Segments[0].C3,
		Color1: core.Color{R: s.Segments[0].Colors[0][0], G: s.Segments[0].Colors[0][1], B: s.Segments[0].Colors[0][2]},
		Color2: core.Color{R: s.Segments[0].Colors[1][0], G: s.Segments[0].Colors[1][1], B: s.Segments[0].Colors[1][2]},
		Color3: core.Color{R: s.Segments[0].Colors[2][0], G: s.Segments[0].Colors[2][1], B: s.Segments[0].Colors[2][2]},
	}
}

func (c *Client) SetState(state core.State) {
	s := State{
		Transition: 2,
	}

	for i, segment := range state.Segments {
		s.Segments = append(s.Segments, Seg{
			Id:    i,
			Start: segment.Start,
			Stop:  segment.End,
			Colors: [][]int{
				{segment.Color1.R, segment.Color1.G, segment.Color1.B, 0},
				{segment.Color2.R, segment.Color2.G, segment.Color2.B, 0},
				{segment.Color3.R, segment.Color3.G, segment.Color3.B, 0},
			},
			Fx: segment.Fx,
			Sx: segment.Sx,
			Ix: segment.Ix,
			O1: segment.O1,
			O2: segment.O2,
			O3: segment.O3,
			C1: segment.C1,
			C2: segment.C2,
			C3: segment.C3,
		})
	}

	for i := len(s.Segments); i < 5; i++ {
		s.Segments = append(s.Segments, Seg{Id: i, Start: 0, Stop: 0})
	}

	_ = c.post("state", s, nil)
}

type State struct {
	Transition int   `json:"tt"`
	Segments   []Seg `json:"seg"`
}

type Seg struct {
	Id    int `json:"id"`
	Start int `json:"start"`
	Stop  int `json:"stop"`

	Fx int `json:"fx"`
	Sx int `json:"sx"`
	Ix int `json:"ix"`

	Colors [][]int `json:"col"`

	O1 bool `json:"o1"`
	O2 bool `json:"o2"`
	O3 bool `json:"o3"`
	C1 int  `json:"c1"`
	C2 int  `json:"c2"`
	C3 int  `json:"c3"`
}
