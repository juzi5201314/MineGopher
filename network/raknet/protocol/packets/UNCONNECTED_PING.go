package packets

import (
	"github.com/juzi5201314/MineGopher/network/raknet/protocol"
)

type UnconnectedPing struct {
	*UnconnectedMessage
	PingTime int64
}

func NewUnconnectedPing() *UnconnectedPing {
	return &UnconnectedPing{NewUnconnectedMessage(protocol.NewPacket(
		UNCONNECTED_PING,
	)), 0}
}

func (ping *UnconnectedPing) Encode() {
	ping.EncodeId()
	ping.PutLong(ping.PingTime)
	ping.PutMagic()
}

func (ping *UnconnectedPing) Decode() {
	ping.DecodeStep()
	ping.PingTime = ping.GetLong()
	ping.ReadMagic()
}
