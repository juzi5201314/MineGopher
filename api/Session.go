package api

import (
	"github.com/juzi5201314/MineGopher/network/raknet/protocol"
)

type Session interface {
	Close()
	IsClosed() bool
	SendPacket(protocol.ConnectedPacket, byte, byte)
}
