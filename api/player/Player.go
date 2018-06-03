package player

import (
	"github.com/google/uuid"
	"github.com/juzi5201314/MineGopher/network/protocol"
)

type Player interface {
	GetUUID() uuid.UUID
	GetXUID() string
	SendPacket(protocol.DataPacket)
}