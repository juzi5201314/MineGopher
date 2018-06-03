package protocol

type FullChunkDataPacket struct {
	*Packet
	ChunkX    int32
	ChunkZ    int32
	ChunkData []byte
}


func (pk *FullChunkDataPacket) Encode() {
	pk.PutVarInt(pk.ChunkX)
	pk.PutVarInt(pk.ChunkZ)
	pk.PutLengthPrefixedBytes(pk.ChunkData)
}

func (pk *FullChunkDataPacket) Decode() {

}
