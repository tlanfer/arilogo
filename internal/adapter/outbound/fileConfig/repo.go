package fileConfig

import (
	"api/internal/core"
	"context"
	"gopkg.in/yaml.v3"
	"os"
)

var _ core.ReSubRepo = (*Repo)(nil)
var _ core.GlobalConfigRepo = (*Repo)(nil)
var _ core.PatternRepo = (*Repo)(nil)

func New(filename string) Repo {
	return Repo{
		filename: filename,
	}
}

type Repo struct {
	filename string
}

func (r Repo) load(_ context.Context) (*Dto, error) {

	file, err := os.OpenFile(r.filename, os.O_CREATE|os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	dto := &Dto{}
	err = yaml.NewDecoder(file).Decode(dto)
	if err != nil {
		return nil, err
	}
	return dto, nil
}

func (r Repo) save(_ context.Context, dto *Dto) error {

	file, err := os.OpenFile(r.filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	err = yaml.NewEncoder(file).Encode(&dto)
	if err != nil {
		return err
	}

	return nil
}

type Dto struct {
	IdleState core.State              `yaml:"idle_state"`
	Subs      map[string]ReSub        `yaml:"resubs"`
	Patterns  map[string]core.Pattern `yaml:"patterns"`
}
