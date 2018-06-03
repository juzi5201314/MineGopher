package network

import (
	"github.com/juzi5201314/MineGopher/network/protocol"
)

type NetWork interface {
	GetName() string
	SetName(string)
	RaknetPacketToMinecraftPaket([]byte) MinecraftPacket
	GetPacket(id byte) (func() protocol.DataPacket, bool)
}
