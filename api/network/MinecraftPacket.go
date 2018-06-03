package network

import "github.com/juzi5201314/MineGopher/network/protocol"

type MinecraftPacket interface {
	GetId() byte
	GetProtocol() int32
	Decode()
	Encode()
	GetPackets() []protocol.DataPacket
	GetBuffer() []byte
	SetBuffer([]byte)
}
