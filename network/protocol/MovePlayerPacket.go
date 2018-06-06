package protocol

import (
	"github.com/golang/geo/r3"
	"github.com/juzi5201314/MineGopher/entity/data"
)

const (
	MoveNormal   = 0
	MoveReset    = 1
	MoveTeleport = 2
	MovePitch    = 3
)

type MovePlayerPacket struct {
	*Packet
	RuntimeId            uint64
	Position             r3.Vector
	Rotation             data.Rotation
	Mode                 byte
	OnGround             bool
	RidingRuntimeId      uint64
	ExtraInt1, ExtraInt2 int32
}

func (pk *MovePlayerPacket) Encode() {
	pk.PutEntityRuntimeId(pk.RuntimeId)
	pk.PutVector(pk.Position)
	pk.PutPlayerRotation(pk.Rotation)
	pk.PutByte(pk.Mode)
	pk.PutBool(pk.OnGround)
	pk.PutEntityRuntimeId(pk.RidingRuntimeId)
	if pk.Mode == MoveTeleport {
		pk.PutLittleInt(pk.ExtraInt1)
		pk.PutLittleInt(pk.ExtraInt2)
	}
}

func (pk *MovePlayerPacket) Decode() {
	pk.RuntimeId = pk.GetEntityRuntimeId()
	pk.Position = pk.GetVector()
	pk.Rotation = pk.GetPlayerRotation()
	pk.Mode = pk.GetByte()
	pk.OnGround = pk.GetBool()
	pk.RidingRuntimeId = pk.GetEntityRuntimeId()
	if pk.Mode == MoveTeleport {
		pk.ExtraInt1 = pk.GetLittleInt()
		pk.ExtraInt2 = pk.GetLittleInt()
	}
}
