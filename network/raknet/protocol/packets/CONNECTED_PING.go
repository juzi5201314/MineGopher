package packets

type ConnectedPing struct {
	*Packet
	PingSendTime int64
}


func NewConnectedPing() *ConnectedPing {
	return &ConnectedPing{NewPacket(CONNECTED_PING), 0}
}

func (ping *ConnectedPing) Encode() {

	ping.EncodeId()

	ping.PutLong(ping.PingSendTime)

}


func (ping *ConnectedPing) Decode() {

	ping.DecodeStep()

	ping.PingSendTime = ping.GetLong()
}