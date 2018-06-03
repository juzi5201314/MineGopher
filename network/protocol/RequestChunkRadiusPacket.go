package protocol

type RequestChunkRadiusPacket struct {
	*Packet
	Radius int32
}


func (pk *RequestChunkRadiusPacket) Encode() {

}

func (pk *RequestChunkRadiusPacket) Decode() {
	pk.Radius = pk.GetVarInt()
}
