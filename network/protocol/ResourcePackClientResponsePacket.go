package protocol

type ResourcePackClientResponsePacket struct {
	*Packet
	Status    byte
	PackUUIDs []string
}

func (pk *ResourcePackClientResponsePacket) Encode() {

}

func (pk *ResourcePackClientResponsePacket) Decode() {
	pk.Status = pk.GetByte()
	var idCount = pk.GetLittleShort()
	for i := int16(0); i < idCount; i++ {
		pk.PackUUIDs = append(pk.PackUUIDs, pk.GetString())
	}
}