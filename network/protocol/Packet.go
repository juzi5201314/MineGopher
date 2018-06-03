package protocol

type Packet struct {
	*PacketStream
	id byte
}

func NewPacket(id byte) *Packet {
	return &Packet{NewPacketStream(), id}
}

func (packet *Packet) EncodeHeader() {
	packet.PutUnsignedVarInt(uint32(packet.id))
	packet.PutByte(0)
	packet.PutByte(0)
}

func (packet *Packet) DecodeHeader() {
	packet.GetUnsignedVarInt()
	packet.GetByte()
	packet.GetByte()
}
