package stream

import (
	"github.com/golang/geo/r3"
	"github.com/google/uuid"
	"github.com/juzi5201314/MineGopher/entity/data"
	"github.com/juzi5201314/MineGopher/item"
	"github.com/juzi5201314/MineGopher/math"
	"github.com/juzi5201314/MineGopher/nbt"
	"github.com/juzi5201314/MineGopher/network/protocol/types"
	"github.com/juzi5201314/MineGopher/utils"
)

type PacketStream struct {
	*utils.Stream
}

func NewPacketStream() *PacketStream {
	return &PacketStream{utils.NewStream()}
}

func (stream *PacketStream) PutEntityRuntimeId(id uint64) {
	stream.PutUnsignedVarLong(id)
}

func (stream *PacketStream) GetEntityRuntimeId() uint64 {
	return stream.GetUnsignedVarLong()
}

func (stream *PacketStream) PutEntityUniqueId(id int64) {
	stream.PutVarLong(id)
}

func (stream *PacketStream) GetEntityUniqueId() int64 {
	return stream.GetVarLong()
}

func (stream *PacketStream) PutVector(vector r3.Vector) {
	stream.PutLittleFloat(float32(vector.X))
	stream.PutLittleFloat(float32(vector.Y))
	stream.PutLittleFloat(float32(vector.Z))
}

func (stream *PacketStream) GetVector() r3.Vector {
	return r3.Vector{X: float64(stream.GetLittleFloat()), Y: float64(stream.GetLittleFloat()), Z: float64(stream.GetLittleFloat())}
}

func (stream *PacketStream) PutBlockPosition(position math.Position) {
	stream.PutVarInt(position.X)
	stream.PutUnsignedVarInt(position.Y)
	stream.PutVarInt(position.Z)
}

func (stream *PacketStream) GetBlockPosition() math.Position {
	return math.NewPosition(stream.GetVarInt(), stream.GetUnsignedVarInt(), stream.GetVarInt())
}

func (stream *PacketStream) PutEntityRotation(rotation data.Rotation) {
	stream.PutLittleFloat(float32(rotation.Pitch))
	stream.PutLittleFloat(float32(rotation.Yaw))
}

func (stream *PacketStream) GetEntityRotation() data.Rotation {
	return data.Rotation{Yaw: float64(stream.GetLittleFloat()), HeadYaw: 0, Pitch: float64(stream.GetLittleFloat())}
}

func (stream *PacketStream) PutPlayerRotation(rot data.Rotation) {
	stream.PutLittleFloat(float32(rot.Pitch))
	stream.PutLittleFloat(float32(rot.Yaw))
	stream.PutLittleFloat(float32(rot.HeadYaw))
}

func (stream *PacketStream) GetPlayerRotation() data.Rotation {
	return data.Rotation{Yaw: float64(stream.GetLittleFloat()), Pitch: float64(stream.GetLittleFloat()), HeadYaw: float64(stream.GetLittleFloat())}
}

func (stream *PacketStream) PutAttributeMap(m data.AttributeMap) {
	stream.PutUnsignedVarInt(uint32(len(m)))
	for _, v := range m {
		stream.PutLittleFloat(v.MinValue)
		stream.PutLittleFloat(v.MaxValue)
		stream.PutLittleFloat(v.Value)
		stream.PutLittleFloat(v.DefaultValue)
		stream.PutString(string(v.GetName()))
	}
}

func (stream *PacketStream) GetAttributeMap() data.AttributeMap {
	m := data.NewAttributeMap()
	c := stream.GetUnsignedVarInt()
	for i := uint32(0); i < c; i++ {
		min := stream.GetLittleFloat()
		max := stream.GetLittleFloat()
		value := stream.GetLittleFloat()
		defaultValue := stream.GetLittleFloat()
		name := data.AttributeName(stream.GetString())
		att := data.NewAttribute(name, value, max)
		att.DefaultValue = defaultValue
		att.MinValue = min
		m.SetAttribute(att)
	}
	return m
}

func (stream *PacketStream) PutItem(i *item.Item) {
	stream.PutVarInt(int32(i.GetId()))
	stream.PutVarInt(int32(((int16(i.GetDamage()) & 0x7fff) << 8) | int16(i.GetCount())))

	writer := nbt.NewWriter(true, utils.LittleEndian)
	writer.WriteUncompressedCompound(i.NBTEmit())
	d := writer.GetBuffer()
	stream.PutLittleShort(int16(len(d)))
	stream.PutBytes(d)
	stream.PutVarInt(0)
	stream.PutVarInt(0)
}

func (stream *PacketStream) GetItem() *item.Item {
	id := stream.GetVarInt()
	if id == 0 {
		i := item.Get(0, 0, 0)
		return i
	}
	aux := stream.GetVarInt()
	damage := aux >> 8
	count := aux & 0xff

	i := item.Get(int(id), int(damage), int8(count))
	nbtData := stream.Get(int(stream.GetLittleShort()))
	reader := nbt.NewReader(nbtData, true, utils.LittleEndian)
	i.NBTParse(reader.ReadUncompressedIntoCompound())
	stream.GetVarInt()
	stream.GetVarInt()
	return i
}

func (stream *PacketStream) PutEntityData(data map[uint32][]interface{}) {
	stream.PutUnsignedVarInt(0)
}

func (stream *PacketStream) GetEntityData() map[uint32][]interface{} {
	return make(map[uint32][]interface{})
}

func (stream *PacketStream) PutGameRules(gameRules map[string]types.GameRuleEntry) {
	stream.PutUnsignedVarInt(uint32(len(gameRules)))
	for _, gameRule := range gameRules {
		stream.PutString(gameRule.Name)
		switch value := gameRule.Value.(type) {
		case bool:
			stream.PutByte(1)
			stream.PutBool(value)
		case uint32:
			stream.PutByte(2)
			stream.PutUnsignedVarInt(value)
		case float32:
			stream.PutByte(3)
			stream.PutLittleFloat(value)
		}
	}
}

func (stream *PacketStream) PutPackInfo(packs []types.ResourcePackInfoEntry) {
	stream.PutLittleShort(int16(len(packs)))

	for _, pack := range packs {
		stream.PutString(pack.UUID)
		stream.PutString(pack.Version)
		stream.PutLittleLong(pack.PackSize)
		stream.PutString("")
		stream.PutString("")
	}
}

func (stream *PacketStream) PutPackStack(packs []types.ResourcePackStackEntry) {
	stream.PutUnsignedVarInt(uint32(len(packs)))
	for _, pack := range packs {
		stream.PutString(pack.UUID)
		stream.PutString(pack.Version)
		stream.PutString("")
	}
}

func (stream *PacketStream) PutUUID(uuid uuid.UUID) {
	b, err := uuid.MarshalBinary()
	if err != nil {
		panic(err)
	}
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	stream.PutBytes(b[8:])
	stream.PutBytes(b[:8])
}

func (stream *PacketStream) GetUUID() uuid.UUID {
	return uuid.Must(uuid.FromBytes(stream.Get(16)))
}
