package protocol

type DataPacket interface {
	Encode()
	EncodeHeader()
	Decode()
	DecodeHeader()
	SetBuffer([]byte)
	GetBuffer() []byte
}
