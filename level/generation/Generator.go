package generation

import (
	"github.com/juzi5201314/MineGopher/level/chunk"
	"errors"
)

type Generator interface {
	GetName() string
	GenerateNewChunk(x, z int32) *chunk.Chunk
}

type Generators map[string]Generator

func NewManager() Generators {
	return Generators{}
}

func (generators Generators) Register(generator Generator) {
	generators[generator.GetName()] = generator
}

func (generators Generators) Deregister(name string) {
	delete(generators, name)
}


func (generators Generators) Get(name string) (Generator, error) {
	if !generators.IsRegistered(name) {
		return nil, errors.New("generator is not registered")
	}
	return generators[name], nil
}

func (generators Generators) IsRegistered(name string) bool {
	var _, ok = generators[name]
	return ok
}
