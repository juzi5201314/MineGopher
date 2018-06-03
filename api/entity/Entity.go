package entity

import (
	"github.com/golang/geo/r3"
	"github.com/juzi5201314/MineGopher/entity/data"
	"github.com/google/uuid"
	"github.com/juzi5201314/MineGopher/level"
	"github.com/juzi5201314/MineGopher/api/level/chunk"
	"github.com/juzi5201314/MineGopher/api/player"
)

type Entity interface {
	GetDimension() level.Dimension
	GetRotation() data.Rotation
	SetRotation(v data.Rotation)
	GetMotion() r3.Vector
	SetMotion(r3.Vector)
	Tick()
	SetDimension(level.Dimension)
	Close()
	IsClosed() bool
	GetId() uint32
	GetEid() uint64
	GetUniqueId() int64
	SetNameTag(string)
	GetNameTag() string
	GetAttributeMap() data.AttributeMap
	SetAttributeMap(attMap data.AttributeMap)
	GetPosition() r3.Vector
	SetPosition(r3.Vector)
	GetChunk() chunk.Chunk
	GetHealth() float32
	SetHealth(health float32)
	SpawnTo(player.Player)
	SpawnToAll()
	Kill()
	DespawnTo(player.Player)
	DespawnToAll()
	GetViewers() map[uuid.UUID]player.Player
	AddViewer(player.Player)
	RemoveViewer(player.Player)
}