package packets

import (
	"github.com/juzi5201314/MineGopher/network/raknet/protocol"
)

type ConnectedPing struct {
	*protocol.Packet
	PingSendTime int64
}

func NewConnectedPing() *ConnectedPing {
	return &ConnectedPing{protocol.NewPacket(CONNECTED_PING), 0}
}

func (ping *ConnectedPing) Encode() {

	ping.EncodeId()

	ping.PutLong(ping.PingSendTime)

}

func (ping *ConnectedPing) Decode() {

	ping.DecodeStep()

	ping.PingSendTime = ping.GetLong()
}
