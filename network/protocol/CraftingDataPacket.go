package protocol

type CraftingDataPacket struct {
	*Packet
}

func (pk *CraftingDataPacket) Encode() {
	pk.PutUnsignedVarInt(0)
	pk.PutBool(true)
}

func (pk *CraftingDataPacket) Decode() {

}
