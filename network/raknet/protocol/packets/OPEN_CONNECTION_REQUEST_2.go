package packets

import (
	"github.com/juzi5201314/MineGopher/network/raknet/protocol"
)

type OpenConnectionRequest2 struct {
	*UnconnectedMessage
	ServerAddress string
	ServerPort    uint16
	MtuSize       int16
	ClientId      int64
}

func NewOpenConnectionRequest2() *OpenConnectionRequest2 {
	return &OpenConnectionRequest2{NewUnconnectedMessage(protocol.NewPacket(
		OPEN_CONNECTION_REQUEST_2,
	)), "", 0, 0, 0}
}

func (request *OpenConnectionRequest2) Encode() {
	request.EncodeId()
	request.PutMagic()
	request.PutAddress(request.ServerAddress, request.ServerPort, 4)
	request.PutShort(request.MtuSize)
	request.PutLong(request.ClientId)
}

func (request *OpenConnectionRequest2) Decode() {
	request.DecodeStep()
	request.ReadMagic()
	var address, port, _ = request.GetAddress()
	request.ServerAddress = address
	request.ServerPort = port
	request.MtuSize = request.GetShort()
	request.ClientId = request.GetLong()
}
