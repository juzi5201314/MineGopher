package protocol

type PlayStatusPacket struct {
	*Packet
	Status int32
}

func (packet *PlayStatusPacket) Encode() {
	packet.PutInt(packet.Status)
}

func (packet *PlayStatusPacket) Decode() {
	packet.Status = packet.GetInt()
}
