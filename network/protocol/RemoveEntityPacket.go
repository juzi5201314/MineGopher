package protocol

type RemoveEntityPacket struct {
	*Packet
	EntityUniqueId int64
}

func (pk *RemoveEntityPacket) Encode() {
	pk.PutEntityUniqueId(pk.EntityUniqueId)
}

func (pk *RemoveEntityPacket) Decode() {

}
