package chunk

import (
	"github.com/golang/geo/r3"
	"github.com/juzi5201314/MineGopher/nbt"
)

type IEntity interface {
	GetEid() uint64
	IsClosed() bool
	GetId() uint32
	Close()
	GetPosition() r3.Vector
	SetPosition(r3.Vector) error
	GetNBT() *nbt.Compound
	SetNBT(*nbt.Compound)
	SetDimension(interface {
		GetChunk(int32, int32) (*Chunk, bool)
	})
	SpawnToAll()
	Tick()
}
