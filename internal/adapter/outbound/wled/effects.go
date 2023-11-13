package wled

import (
	"api/internal/core"
	"strconv"
	"strings"
)

func (c *Client) Effects() []Effect {
	var effectNames []string
	var fxData []string

	var effects []Effect

	if err := c.get("eff", &effectNames); err != nil {
		return effects
	}

	if err := c.get("fxdata", &fxData); err != nil {
		return effects
	}

	for i := range effectNames {
		if effectNames[i] == "RSVD" || effectNames[i] == "-" {
			continue
		}
		effects = append(effects, parseEffect(i, effectNames[i], fxData[i]))
	}

	effects = core.Filter(effects, func(effect Effect) bool {
		return effect.Flags.Led1D == true
	})

	return effects
}

func parseEffect(id int, name, fxData string) Effect {
	parts := strings.SplitN(fxData, ";", 5)

	effect := Effect{
		Id:   id,
		Name: name,
		Colors: Colors{
			Color1: "Fx",
			Color2: "Bg",
			Color3: "Cs",
		},
		Palette: true,
		Flags:   Flags{Led1D: true},
	}

	if len(parts) > 0 {
		effect.Parameters = parseParamList(parts[0])
	}

	if len(parts) > 1 {
		effect.Colors = parseColorList(parts[1])
	}

	if len(parts) > 2 {
		effect.Palette = parts[2] == "!"
	}

	if len(parts) > 3 {
		effect.Flags = parseFlags(parts[3])
	}

	if len(parts) > 4 {
		parseDefaultsList(parts[4], &effect.Parameters)
	}

	return effect
}
func parseParamList(params string) Parameters {
	n := strings.SplitN(params, ",", 8)
	return Parameters{
		SX: parseParamSingle(n, 0, TypeSlider, "Effect speed"),
		IX: parseParamSingle(n, 1, TypeSlider, "Effect intensity"),
		C1: parseParamSingle(n, 2, TypeSlider, "Unknown"),
		C2: parseParamSingle(n, 3, TypeSlider, "Unknown"),
		C3: parseParamSingle(n, 4, TypeSlider, "Unknown"),
		O1: parseParamSingle(n, 5, TypeCheckBox, "Unknown"),
		O2: parseParamSingle(n, 6, TypeCheckBox, "Unknown"),
		O3: parseParamSingle(n, 7, TypeCheckBox, "Unknown"),
	}
}
func parseParamSingle(v []string, n int, t string, def string) Parameter {
	if len(v) < n+1 {
		return Parameter{Type: TypeNone}
	}

	switch v[n] {
	case "!":
		return Parameter{Type: t, Name: def}
	case "":
		return Parameter{Type: TypeNone}
	default:
		return Parameter{Type: t, Name: v[n]}
	}
}
func parseColorList(colors string) Colors {
	n := strings.SplitN(colors, ",", 3)
	return Colors{
		Color1: parseColorSimple(n, 0, "Fx"),
		Color2: parseColorSimple(n, 1, "Bg"),
		Color3: parseColorSimple(n, 2, "Cs"),
	}
}
func parseColorSimple(v []string, n int, def string) string {
	if len(v) < n+1 {
		return ""
	}
	if v[n] == "!" {
		return def
	}
	return v[n]
}
func parseFlags(flags string) Flags {
	return Flags{
		Led1D:                  strings.Contains(flags, "1"),
		Led2D:                  strings.Contains(flags, "2"),
		Led3D:                  strings.Contains(flags, "3"),
		AudioReactiveVolume:    strings.Contains(flags, "v"),
		AudioReactiveFrequency: strings.Contains(flags, "f"),
	}
}
func parseDefaultsList(defaults string, params *Parameters) {
	a := func(x string) int {
		i, _ := strconv.Atoi(x)
		return i
	}
	n := strings.Split(defaults, ",")
	for _, d := range n {
		key, value, _ := strings.Cut(d, "=")
		switch key {
		case "sx":
			params.SX.Def = a(value)
		case "ix":
			params.IX.Def = a(value)
		case "c1":
			params.C1.Def = a(value)
		case "c2":
			params.C2.Def = a(value)
		case "c3":
			params.C3.Def = a(value)
		case "o1":
			params.O1.Def = a(value)
		case "o2":
			params.O2.Def = a(value)
		case "o3":
			params.O3.Def = a(value)
		}
	}
}

type Effect struct {
	Id         int
	Name       string
	Parameters Parameters
	Colors     Colors
	Palette    bool
	Flags      Flags
}

type Parameters struct {
	SX, IX, C1, C2, C3, O1, O2, O3 Parameter
}

type Parameter struct {
	Type string
	Name string
	Def  int
}

const (
	TypeSlider   string = "SLIDER"
	TypeCheckBox        = "CHECKBOX"
	TypeNone            = "NONE"
)

type Colors struct {
	Color1 string `json:"color1"`
	Color2 string `json:"color2"`
	Color3 string `json:"color3"`
}

type Flags struct {
	Led1D                  bool
	Led2D                  bool
	Led3D                  bool
	AudioReactiveVolume    bool
	AudioReactiveFrequency bool
}
