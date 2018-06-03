package player

import (
	"github.com/juzi5201314/MineGopher/network/protocol"
	"github.com/google/uuid"
)

type Player interface {
	GetUUID() uuid.UUID
	GetXUID() string
	SendPacket(protocol.DataPacket)
	SpawnTo(Player)
	Tick()
}