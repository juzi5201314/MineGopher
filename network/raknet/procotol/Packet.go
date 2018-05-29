package procotol

import "github.com/juzi5201314/MineGopher/network/raknet/procotol/packets"

var PacketPool = map[byte]interface{} {
	0x00: new(packets.UNCONNECTED_PING_OPEN_CONNECTIONS),
}

func GetPacket(id byte) *DataPacket {
	if pk, exists := PacketPool[id]; exists {
		return pk.(DataPacket).New()
	}
	return nil
}

type DataPacket interface {
	New() *DataPacket
}