package packets

import (
	"github.com/juzi5201314/MineGopher/network/raknet/protocol"
)

type NACK struct {
*AcknowledgementPacket
}

func NewNACK() *NACK {
return &NACK{
&AcknowledgementPacket{protocol.NewPacket(FLAG_DATAGRAM_NACK), []uint32{}}}
}
