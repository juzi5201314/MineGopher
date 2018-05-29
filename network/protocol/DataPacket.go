package protocol

type DataPacket interface {
	New() *DataPacket
	Encode()
	Decode()
}
