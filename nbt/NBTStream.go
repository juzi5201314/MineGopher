package nbt

import "github.com/juzi5201314/MineGopher/utils"

type IStream interface {
	SetBuffer([]byte)
	GetBuffer() []byte
	GetOffset() int
	SetOffset(int)
	Feof() bool
	Get(int) []byte
	PutBool(bool)
	GetBool() bool
	PutByte(byte)
	GetByte() byte
	PutShort(int16)
	GetShort() int16
	PutInt(int32)
	GetInt() int32
	PutLong(int64)
	GetLong() int64
	PutFloat(float32)
	GetFloat() float32
	PutDouble(float64)
	GetDouble() float64
	PutString(string)
	GetString() string
	PutBytes([]byte)
	ResetStream()
}

type BinaryStream struct {
	IStream
}

func NewStream(buffer []byte, network bool, endian utils.EndianType) *BinaryStream {
	var stream IStream
	if network && endian == utils.LittleEndian {
		stream = NewLittleEndianNetworkStream(buffer)
	} else {
		if endian == utils.LittleEndian {
			stream = NewLittleEndianStream(buffer)
		} else {
			stream = NewBigEndianStream(buffer)
		}
	}
	return &BinaryStream{stream}
}

type LittleEndianStream struct {
	Offset int
	Buffer []byte
}

func NewLittleEndianStream(buffer []byte) *LittleEndianStream {
	return &LittleEndianStream{0, buffer}
}

func (stream *LittleEndianStream) SetBuffer(Buffer []byte) { stream.Buffer = Buffer }

func (stream *LittleEndianStream) GetBuffer() []byte { return stream.Buffer }

func (stream *LittleEndianStream) GetOffset() int { return stream.Offset }

func (stream *LittleEndianStream) SetOffset(offset int) { stream.Offset = offset }

func (stream *LittleEndianStream) Feof() bool { return stream.Offset >= len(stream.Buffer)-1 }

func (stream *LittleEndianStream) Get(length int) []byte {
	return utils.Read(&stream.Buffer, &stream.Offset, length)
}

func (stream *LittleEndianStream) GetRemainder() []byte {
	return utils.Read(&stream.Buffer, &stream.Offset, len(stream.Buffer)-stream.Offset)
}

func (stream *LittleEndianStream) PutBool(v bool) { utils.WriteBool(&stream.Buffer, v) }

func (stream *LittleEndianStream) GetBool() bool {
	return utils.ReadBool(&stream.Buffer, &stream.Offset)
}

func (stream *LittleEndianStream) PutByte(v byte) { utils.WriteByte(&stream.Buffer, v) }

func (stream *LittleEndianStream) GetByte() byte {
	return utils.ReadByte(&stream.Buffer, &stream.Offset)
}

func (stream *LittleEndianStream) PutShort(v int16) { utils.WriteLittleShort(&stream.Buffer, v) }

func (stream *LittleEndianStream) GetShort() int16 {
	return utils.ReadLittleShort(&stream.Buffer, &stream.Offset)
}

func (stream *LittleEndianStream) PutInt(v int32) { utils.WriteLittleInt(&stream.Buffer, v) }

func (stream *LittleEndianStream) GetInt() int32 {
	return utils.ReadLittleInt(&stream.Buffer, &stream.Offset)
}

func (stream *LittleEndianStream) PutLong(v int64) { utils.WriteLittleLong(&stream.Buffer, v) }

func (stream *LittleEndianStream) GetLong() int64 {
	return utils.ReadLittleLong(&stream.Buffer, &stream.Offset)
}

func (stream *LittleEndianStream) PutFloat(v float32) { utils.WriteLittleFloat(&stream.Buffer, v) }

func (stream *LittleEndianStream) GetFloat() float32 {
	return utils.ReadLittleFloat(&stream.Buffer, &stream.Offset)
}

func (stream *LittleEndianStream) PutDouble(v float64) { utils.WriteLittleDouble(&stream.Buffer, v) }

func (stream *LittleEndianStream) GetDouble() float64 {
	return utils.ReadLittleDouble(&stream.Buffer, &stream.Offset)
}

func (stream *LittleEndianStream) PutString(v string) {
	utils.WriteUnsignedShort(&stream.Buffer, uint16(len(v)))
	stream.Buffer = append(stream.Buffer, []byte(v)...)
}

func (stream *LittleEndianStream) GetString() string {
	var length = utils.ReadUnsignedShort(&stream.Buffer, &stream.Offset)
	return string(utils.Read(&stream.Buffer, &stream.Offset, int(length)))
}

func (stream *LittleEndianStream) PutBytes(bytes []byte) {
	stream.Buffer = append(stream.Buffer, bytes...)
}

func (stream *LittleEndianStream) ResetStream() {
	stream.Offset = 0
	stream.Buffer = []byte{}
}

type BigEndianStream struct {
	Offset int
	Buffer []byte
}

func NewBigEndianStream(buffer []byte) *BigEndianStream {
	return &BigEndianStream{0, buffer}
}

func (stream *BigEndianStream) SetBuffer(Buffer []byte) { stream.Buffer = Buffer }

func (stream *BigEndianStream) GetBuffer() []byte { return stream.Buffer }

func (stream *BigEndianStream) GetOffset() int { return stream.Offset }

func (stream *BigEndianStream) SetOffset(offset int) { stream.Offset = offset }

func (stream *BigEndianStream) Feof() bool { return stream.Offset >= len(stream.Buffer)-1 }

func (stream *BigEndianStream) Get(length int) []byte {
	return utils.Read(&stream.Buffer, &stream.Offset, length)
}

func (stream *BigEndianStream) GetRemainder() []byte {
	return utils.Read(&stream.Buffer, &stream.Offset, len(stream.Buffer)-stream.Offset)
}

func (stream *BigEndianStream) PutBool(v bool) { utils.WriteBool(&stream.Buffer, v) }

func (stream *BigEndianStream) GetBool() bool { return utils.ReadBool(&stream.Buffer, &stream.Offset) }

func (stream *BigEndianStream) PutByte(v byte) { utils.WriteByte(&stream.Buffer, v) }

func (stream *BigEndianStream) GetByte() byte { return utils.ReadByte(&stream.Buffer, &stream.Offset) }

func (stream *BigEndianStream) PutShort(v int16) { utils.WriteShort(&stream.Buffer, v) }

func (stream *BigEndianStream) GetShort() int16 {
	return utils.ReadShort(&stream.Buffer, &stream.Offset)
}

func (stream *BigEndianStream) PutInt(v int32) { utils.WriteInt(&stream.Buffer, v) }

func (stream *BigEndianStream) GetInt() int32 { return utils.ReadInt(&stream.Buffer, &stream.Offset) }

func (stream *BigEndianStream) PutLong(v int64) { utils.WriteLong(&stream.Buffer, v) }

func (stream *BigEndianStream) GetLong() int64 { return utils.ReadLong(&stream.Buffer, &stream.Offset) }

func (stream *BigEndianStream) PutFloat(v float32) { utils.WriteFloat(&stream.Buffer, v) }

func (stream *BigEndianStream) GetFloat() float32 {
	return utils.ReadFloat(&stream.Buffer, &stream.Offset)
}

func (stream *BigEndianStream) PutDouble(v float64) { utils.WriteDouble(&stream.Buffer, v) }

func (stream *BigEndianStream) GetDouble() float64 {
	return utils.ReadDouble(&stream.Buffer, &stream.Offset)
}

func (stream *BigEndianStream) PutString(v string) {
	utils.WriteUnsignedShort(&stream.Buffer, uint16(len(v)))
	stream.Buffer = append(stream.Buffer, []byte(v)...)
}

func (stream *BigEndianStream) GetString() string {
	var length = utils.ReadUnsignedShort(&stream.Buffer, &stream.Offset)
	return string(utils.Read(&stream.Buffer, &stream.Offset, int(length)))
}

func (stream *BigEndianStream) PutBytes(bytes []byte) { stream.Buffer = append(stream.Buffer, bytes...) }

func (stream *BigEndianStream) ResetStream() {
	stream.Offset = 0
	stream.Buffer = []byte{}
}

type LittleEndianNetworkStream struct {
	Offset int
	Buffer []byte
}

func NewLittleEndianNetworkStream(buffer []byte) *LittleEndianNetworkStream {
	return &LittleEndianNetworkStream{0, buffer}
}

func (stream *LittleEndianNetworkStream) SetBuffer(Buffer []byte) { stream.Buffer = Buffer }

func (stream *LittleEndianNetworkStream) GetBuffer() []byte { return stream.Buffer }

func (stream *LittleEndianNetworkStream) GetOffset() int { return stream.Offset }

func (stream *LittleEndianNetworkStream) SetOffset(offset int) { stream.Offset = offset }

func (stream *LittleEndianNetworkStream) Feof() bool { return stream.Offset >= len(stream.Buffer)-1 }

func (stream *LittleEndianNetworkStream) Get(length int) []byte {
	return utils.Read(&stream.Buffer, &stream.Offset, length)
}

func (stream *LittleEndianNetworkStream) GetRemainder() []byte {
	return utils.Read(&stream.Buffer, &stream.Offset, len(stream.Buffer)-stream.Offset)
}

func (stream *LittleEndianNetworkStream) PutBool(v bool) { utils.WriteBool(&stream.Buffer, v) }

func (stream *LittleEndianNetworkStream) GetBool() bool {
	return utils.ReadBool(&stream.Buffer, &stream.Offset)
}

func (stream *LittleEndianNetworkStream) PutByte(v byte) { utils.WriteByte(&stream.Buffer, v) }

func (stream *LittleEndianNetworkStream) GetByte() byte {
	return utils.ReadByte(&stream.Buffer, &stream.Offset)
}

func (stream *LittleEndianNetworkStream) PutShort(v int16) { utils.WriteLittleShort(&stream.Buffer, v) }

func (stream *LittleEndianNetworkStream) GetShort() int16 {
	return utils.ReadLittleShort(&stream.Buffer, &stream.Offset)
}

func (stream *LittleEndianNetworkStream) PutInt(v int32) { utils.WriteVarInt(&stream.Buffer, v) }

func (stream *LittleEndianNetworkStream) GetInt() int32 {
	return utils.ReadVarInt(&stream.Buffer, &stream.Offset)
}

func (stream *LittleEndianNetworkStream) PutLong(v int64) { utils.WriteVarLong(&stream.Buffer, v) }

func (stream *LittleEndianNetworkStream) GetLong() int64 {
	return utils.ReadVarLong(&stream.Buffer, &stream.Offset)
}

func (stream *LittleEndianNetworkStream) PutFloat(v float32) {
	utils.WriteLittleFloat(&stream.Buffer, v)
}

func (stream *LittleEndianNetworkStream) GetFloat() float32 {
	return utils.ReadLittleFloat(&stream.Buffer, &stream.Offset)
}

func (stream *LittleEndianNetworkStream) PutDouble(v float64) {
	utils.WriteLittleDouble(&stream.Buffer, v)
}

func (stream *LittleEndianNetworkStream) GetDouble() float64 {
	return utils.ReadLittleDouble(&stream.Buffer, &stream.Offset)
}

func (stream *LittleEndianNetworkStream) PutString(v string) {
	utils.WriteUnsignedVarInt(&stream.Buffer, uint32(len(v)))
	stream.Buffer = append(stream.Buffer, []byte(v)...)
}

func (stream *LittleEndianNetworkStream) GetString() string {
	var length = utils.ReadUnsignedVarInt(&stream.Buffer, &stream.Offset)
	return string(utils.Read(&stream.Buffer, &stream.Offset, int(length)))
}

func (stream *LittleEndianNetworkStream) PutBytes(bytes []byte) {
	stream.Buffer = append(stream.Buffer, bytes...)
}

func (stream *LittleEndianNetworkStream) ResetStream() {
	stream.Offset = 0
	stream.Buffer = []byte{}
}
