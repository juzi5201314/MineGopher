package utils

// Stream is a container of a byte array and an offset.

// Reading from the stream increments the offset.

type Stream struct {
	Offset int

	Buffer []byte
}

// NewStream returns a new stream.

func NewStream() *Stream {

	return &Stream{0, []byte{}}

}

// GetOffset returns the current stream offset.

func (stream *Stream) GetOffset() int {

	return stream.Offset

}

// SetOffset sets the offset of the stream.

func (stream *Stream) SetOffset(offset int) {

	stream.Offset = offset

}

// SetBuffer sets the buffer of the stream.

func (stream *Stream) SetBuffer(buffer []byte) {

	stream.Buffer = buffer

}

// GetBuffer returns the buffer of the stream.

func (stream *Stream) GetBuffer() []byte {

	return stream.Buffer

}

// Feof checks if the stream offset reached the end of its buffer.

func (stream *Stream) Feof() bool {

	return stream.Offset >= len(stream.Buffer)-1

}

// Get reads the given amount of bytes from the buffer.

// If length is negative, reads the leftover bytes.

func (stream *Stream) Get(length int) []byte {

	if length < 0 {

		length = len(stream.Buffer) - stream.Offset - 1

	}

	return Read(&stream.Buffer, &stream.Offset, length)

}

func (stream *Stream) PutBool(v bool) {

	WriteBool(&stream.Buffer, v)

}

func (stream *Stream) GetBool() bool {

	return ReadBool(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutByte(v byte) {

	WriteByte(&stream.Buffer, v)

}

func (stream *Stream) GetByte() byte {

	return ReadByte(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutUnsignedByte(v byte) {

	WriteUnsignedByte(&stream.Buffer, v)

}

func (stream *Stream) GetUnsignedByte() byte {

	return ReadUnsignedByte(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutShort(v int16) {

	WriteShort(&stream.Buffer, v)

}

func (stream *Stream) GetShort() int16 {

	return ReadShort(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutUnsignedShort(v uint16) {

	WriteUnsignedShort(&stream.Buffer, v)

}

func (stream *Stream) GetUnsignedShort() uint16 {

	return ReadUnsignedShort(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutInt(v int32) {

	WriteInt(&stream.Buffer, v)

}

func (stream *Stream) GetInt() int32 {

	return ReadInt(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutUnsignedInt(v uint32) {

	WriteUnsignedInt(&stream.Buffer, v)

}

func (stream *Stream) GetUnsignedInt() uint32 {

	return ReadUnsignedInt(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutLong(v int64) {

	WriteLong(&stream.Buffer, v)

}

func (stream *Stream) GetLong() int64 {

	return ReadLong(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutUnsignedLong(v uint64) {

	WriteUnsignedLong(&stream.Buffer, v)

}

func (stream *Stream) GetUnsignedLong() uint64 {

	return ReadUnsignedLong(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutFloat(v float32) {

	WriteFloat(&stream.Buffer, v)

}

func (stream *Stream) GetFloat() float32 {

	return ReadFloat(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutDouble(v float64) {

	WriteDouble(&stream.Buffer, v)

}

func (stream *Stream) GetDouble() float64 {

	return ReadDouble(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutVarInt(v int32) {

	WriteVarInt(&stream.Buffer, v)

}

func (stream *Stream) GetVarInt() int32 {

	return ReadVarInt(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutVarLong(v int64) {

	WriteVarLong(&stream.Buffer, v)

}

func (stream *Stream) GetVarLong() int64 {

	return ReadVarLong(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutUnsignedVarInt(v uint32) {

	WriteUnsignedVarInt(&stream.Buffer, v)

}

func (stream *Stream) GetUnsignedVarInt() uint32 {

	return ReadUnsignedVarInt(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutUnsignedVarLong(v uint64) {

	WriteUnsignedVarLong(&stream.Buffer, v)

}

func (stream *Stream) GetUnsignedVarLong() uint64 {

	return ReadUnsignedVarLong(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutString(v string) {

	WriteUnsignedVarInt(&stream.Buffer, uint32(len(v)))

	stream.Buffer = append(stream.Buffer, []byte(v)...)

}

func (stream *Stream) GetString() string {

	return string(Read(&stream.Buffer, &stream.Offset, int(stream.GetUnsignedVarInt())))

}

func (stream *Stream) PutLittleShort(v int16) {

	WriteLittleShort(&stream.Buffer, v)

}

func (stream *Stream) GetLittleShort() int16 {

	return ReadLittleShort(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutLittleUnsignedShort(v uint16) {

	WriteLittleUnsignedShort(&stream.Buffer, v)

}

func (stream *Stream) GetLittleUnsignedShort() uint16 {

	return ReadLittleUnsignedShort(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutLittleInt(v int32) {

	WriteLittleInt(&stream.Buffer, v)

}

func (stream *Stream) GetLittleInt() int32 {

	return ReadLittleInt(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutLittleUnsignedInt(v uint32) {

	WriteLittleUnsignedInt(&stream.Buffer, v)

}

func (stream *Stream) GetLittleUnsignedInt() uint32 {

	return ReadLittleUnsignedInt(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutLittleLong(v int64) {

	WriteLittleLong(&stream.Buffer, v)

}

func (stream *Stream) GetLittleLong() int64 {

	return ReadLittleLong(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutLittleUnsignedLong(v uint64) {

	WriteLittleUnsignedLong(&stream.Buffer, v)

}

func (stream *Stream) GetLittleUnsignedLong() uint64 {

	return ReadLittleUnsignedLong(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutLittleFloat(v float32) {

	WriteLittleFloat(&stream.Buffer, v)

}

func (stream *Stream) GetLittleFloat() float32 {

	return ReadLittleFloat(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutLittleDouble(v float64) {

	WriteLittleDouble(&stream.Buffer, v)

}

func (stream *Stream) GetLittleDouble() float64 {

	return ReadLittleDouble(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutTriad(v uint32) {

	WriteBigTriad(&stream.Buffer, v)

}

func (stream *Stream) GetTriad() uint32 {

	return ReadBigTriad(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutLittleTriad(v uint32) {

	WriteLittleTriad(&stream.Buffer, v)

}

func (stream *Stream) GetLittleTriad() uint32 {

	return ReadLittleTriad(&stream.Buffer, &stream.Offset)

}

func (stream *Stream) PutBytes(bytes []byte) {

	stream.Buffer = append(stream.Buffer, bytes...)

}

func (stream *Stream) PutLengthPrefixedBytes(bytes []byte) {

	stream.PutUnsignedVarInt(uint32(len(bytes)))

	stream.PutBytes(bytes)

}

func (stream *Stream) GetLengthPrefixedBytes() []byte {

	return []byte(stream.GetString())

}

func (stream *Stream) ResetStream() {

	stream.Offset = 0

	stream.Buffer = []byte{}

}
