package packets

import (
	"github.com/juzi5201314/MineGopher/network/raknet/protocol"
)

type OpenConnectionRequest1 struct {
	*UnconnectedMessage
	Protocol byte
	MtuSize  int16
}

func NewOpenConnectionRequest1() *OpenConnectionRequest1 {
	return &OpenConnectionRequest1{NewUnconnectedMessage(protocol.NewPacket(
		OPEN_CONNECTION_REQUEST_1,
	)), 0, 0}
}

func (request *OpenConnectionRequest1) Encode() {
	request.EncodeId()
	request.PutMagic()
	request.PutByte(request.Protocol)

	bytes := make([]byte, request.MtuSize+28)
	request.PutBytes(bytes)
}

func (request *OpenConnectionRequest1) Decode() {
	request.DecodeStep()
	request.ReadMagic()
	request.Protocol = request.GetByte()
	request.MtuSize = int16(len(request.Buffer)) + 28 // Account for UDP and IP headers.
}