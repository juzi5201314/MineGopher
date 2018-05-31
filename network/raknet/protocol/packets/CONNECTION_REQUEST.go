package packets

import (
	"github.com/juzi5201314/MineGopher/network/raknet/protocol"
)

type ConnectionRequest struct {
	*protocol.Packet
	ClientId uint64
	PingSendTime uint64
	Security byte
}

func NewConnectionRequest() *ConnectionRequest {
	return &ConnectionRequest{protocol.NewPacket(CONNECTION_REQUEST), 0, 0, 0}
}

func (request *ConnectionRequest) Encode() {
	request.EncodeId()
	request.PutUnsignedLong(request.ClientId)
	request.PutUnsignedLong(request.PingSendTime)
}

func (request *ConnectionRequest) Decode() {
	request.DecodeStep()
	request.ClientId = request.GetUnsignedLong()
request.PingSendTime = request.GetUnsignedLong()
}

