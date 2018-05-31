package packets

import (
	"github.com/juzi5201314/MineGopher/network/raknet/protocol"
)

type ConnectedPong struct {
	*protocol.Packet
	PingSendTime int64
	PongSendTime int64
}

func NewConnectedPong() *ConnectedPong {

	return &ConnectedPong{protocol.NewPacket(CONNECTED_PONG), 0, 0}
}

func (pong *ConnectedPong) Encode() {
	pong.EncodeId()
	pong.PutLong(pong.PingSendTime)
	pong.PutLong(pong.PongSendTime)
}

func (pong *ConnectedPong) Decode() {
	pong.DecodeStep()
	pong.PingSendTime = pong.GetLong()
	pong.PongSendTime = pong.GetLong()
}
