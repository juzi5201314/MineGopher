package protocol

import (
	"github.com/golang/geo/r3"
	"github.com/google/uuid"
	"github.com/juzi5201314/MineGopher/entity/data"
	"github.com/juzi5201314/MineGopher/level"
)

type AddEntityEntry interface {
	GetUniqueId() int64
	GetEid() uint64
	GetId() uint32
	GetPosition() r3.Vector
	GetMotion() r3.Vector
	GetRotation() data.Rotation
	GetAttributeMap() data.AttributeMap
	GetEntityData() map[uint32][]interface{}
}

type AddPlayerEntry interface {
	AddEntityEntry
	GetDisplayName() string
	GetName() string
}

type PlayerListEntry interface {
	AddPlayerEntry
	GetXUID() string
	GetUUID() uuid.UUID
	GetSkinId() string
	GetSkinData() []byte
	GetCapeData() []byte
	GetGeometryName() string
	GetGeometryData() string
	GetPlatform() int32
}

type StartGameEntry interface {
	GetRuntimeId() uint64
	GetUniqueId() int64
	GetPosition() r3.Vector
	GetDimension() *level.Dimension
}
