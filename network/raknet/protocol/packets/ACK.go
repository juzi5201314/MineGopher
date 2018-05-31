package packets

import (
	"github.com/juzi5201314/MineGopher/network/raknet/protocol"
)

type ACK struct {
	*AcknowledgementPacket
}

func NewACK() *ACK {
	return &ACK{
		&AcknowledgementPacket{
			protocol.NewPacket(FLAG_DATAGRAM_ACK),
			[]uint32{}}}
}
