package protocol


import "github.com/juzi5201314/MineGopher/network/protocol/types"

type ResourcePackStackPacket struct {
	*Packet
	MustAccept    bool
	BehaviorPacks []types.ResourcePackInfoEntry
	ResourcePacks []types.ResourcePackInfoEntry
}

func (pk *ResourcePackStackPacket) Encode() {
	pk.PutBool(pk.MustAccept)
	pk.PutPackInfo(pk.BehaviorPacks)
	pk.PutPackInfo(pk.ResourcePacks)
}

func (pk *ResourcePackStackPacket) Decode() {

}