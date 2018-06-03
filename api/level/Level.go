package level

import (
	"github.com/golang/geo/r3"
	"github.com/juzi5201314/MineGopher/api/level/chunk"
	"github.com/juzi5201314/MineGopher/level"
	"github.com/juzi5201314/MineGopher/api/entity"
)

type Level interface {
	GetDimension() Dimension
	GetGameRule(gameRule level.GameRuleName) GameRule
	GetGameRules() map[level.GameRuleName]GameRule
	GetName() string
}

type GameRule interface {
	GetName() level.GameRuleName
	GetValue() interface{}
	SetValue(value interface{}) bool
}

type Dimension interface {
	AddEntity(entity.Entity, r3.Vector)
	LoadChunk(x, z int32, function func(chunk.Chunk))
	GetChunk(x,z int32) (chunk.Chunk, bool)
}