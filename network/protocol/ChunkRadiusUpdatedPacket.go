package protocol

type ChunkRadiusUpdatedPacket struct {
	*Packet
	Radius int32
}

func (pk *ChunkRadiusUpdatedPacket) Encode() {
	pk.PutVarInt(pk.Radius)
}

func (pk *ChunkRadiusUpdatedPacket) Decode() {

}
