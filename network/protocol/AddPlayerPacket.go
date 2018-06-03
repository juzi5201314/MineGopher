package protocol

import (
	"github.com/golang/geo/r3"
	"github.com/google/uuid"
	"github.com/juzi5201314/MineGopher/entity/data"
)

type AddPlayerPacket struct {
	*Packet
	UUID            uuid.UUID
	Username        string
	DisplayName     string
	Platform        int32
	UnknownString   string
	EntityUniqueId  int64
	EntityRuntimeId uint64
	Position        r3.Vector
	Motion          r3.Vector
	Rotation        data.Rotation
	Metadata          map[uint32][]interface{}
	Flags             uint32
	CommandPermission uint32
	Flags2            uint32
	PlayerPermission  uint32
	CustomFlags       uint32
	Long1             int64
}


func (pk *AddPlayerPacket) Encode() {
	pk.PutUUID(pk.UUID)
	pk.PutString(pk.Username)
	pk.PutString(pk.DisplayName)
	pk.PutVarInt(pk.Platform)

	pk.PutEntityUniqueId(pk.EntityUniqueId)
	pk.PutEntityRuntimeId(pk.EntityRuntimeId)
	pk.PutString(pk.UnknownString)

	pk.PutVector(pk.Position)
	pk.PutVector(pk.Motion)
	pk.PutPlayerRotation(pk.Rotation)

	pk.PutVarInt(0) // TODO
	pk.PutEntityData(pk.Metadata)

	pk.PutUnsignedVarInt(pk.Flags)
	pk.PutUnsignedVarInt(pk.CommandPermission)
	pk.PutUnsignedVarInt(pk.Flags2)
	pk.PutUnsignedVarInt(pk.PlayerPermission)
	pk.PutUnsignedVarInt(pk.CustomFlags)

	pk.PutVarLong(pk.Long1)

	pk.PutUnsignedVarInt(0) // TODO
}

func (pk *AddPlayerPacket) Decode() {

}
