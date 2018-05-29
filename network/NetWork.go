package network

import (
	"github.com/juzi5201314/MineGopher/api"
	"github.com/juzi5201314/MineGopher/network/protocol"
)

func New(s api.Server) *NetWork {
	network := new(NetWork)
	network.packets = map[byte]protocol.DataPacket{}
	network.registerPackets()
	return network
}

type NetWork struct {
	packets  map[byte]protocol.DataPacket
	upload   float64
	download float64
	name     string
}

func (network *NetWork) GetUpload() float64 {
	return network.upload
}

func (network *NetWork) GetDownload() float64 {
	return network.download
}

func (network *NetWork) AddUpload(b float64) {
	network.upload += b
}

func (network *NetWork) AddDownload(b float64) {
	network.download += b
}

func (network *NetWork) GetName() string {
	return network.name
}

func (network *NetWork) SetName(name string) {
	network.name = name
}

func (network *NetWork) GetPacket(id byte) *protocol.DataPacket {
	if packet, has := network.packets[id]; has {
		return packet.New()
	}
	return nil
}

func (network *NetWork) registerPacket(id byte, packet protocol.DataPacket) {
	network.packets[id] = packet
}

func (network *NetWork) registerPackets() {

}
