package protocol

import "github.com/juzi5201314/MineGopher/network/protocol/types"

type ResourcePackInfoPacket struct {
	*Packet
	MustAccept    bool
	BehaviorPacks []types.ResourcePackInfoEntry
	ResourcePacks []types.ResourcePackInfoEntry
}

func (pk *ResourcePackInfoPacket) Encode() {
	pk.PutBool(pk.MustAccept)
	pk.PutPackInfo(pk.BehaviorPacks)
	pk.PutPackInfo(pk.ResourcePacks)
}

func (pk *ResourcePackInfoPacket) Decode() {

}