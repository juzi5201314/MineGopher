package protocol

import "github.com/juzi5201314/MineGopher/entity/data"

type UpdateAttributesPacket struct {
	*Packet
	RuntimeId  uint64
	Attributes data.AttributeMap
}

func (pk *UpdateAttributesPacket) Encode() {
	pk.PutEntityRuntimeId(pk.RuntimeId)
	pk.PutAttributeMap(pk.Attributes)
}

func (pk *UpdateAttributesPacket) Decode() {
	pk.RuntimeId = pk.GetEntityRuntimeId()
	pk.Attributes = pk.GetAttributeMap()
}
