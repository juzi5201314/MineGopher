package chunk

import (
	"github.com/google/uuid"
	"github.com/juzi5201314/MineGopher/api/player"
)

type Chunk interface {
	GetViewers() map[uuid.UUID]player.Player
	AddEntity(Entity) error
	RemoveEntity(uint64)
	AddViewer(Player)
}