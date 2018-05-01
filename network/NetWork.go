package network

import (
	"github.com/juzi5201314/minegopher/network/protocol"
	"github.com/juzi5201314/minegopher"
)

func NewNetWork(server *minegopher.Server) *NetWork {
	network := &NetWork{}
	network.server = server
	network.registerPackets()
	return network
}

type NetWork struct {
	packets []protocol.DataPacket
	interfaces []RaknetInterface
	upload float64
	download float64
	name string
	server *minegopher.Server
}

func (network *NetWork) registerInterface(i RaknetInterface) {
	network.interfaces = append(network.interfaces, i)
	i.SetNetWork(network)
	i.SetName(network.name)
}

func (network *NetWork) addStatistics(upload float64, download float64) {
	network.upload += upload
	network.download += download
}

func (network *NetWork) GetUpload() float64 {
	return network.upload
}

func (network *NetWork) GetDownload() float64 {
	return network.download
}

func (network *NetWork) GetName() string {
	return network.name
}

func (network *NetWork) SetName(name string) {
	network.name = name
}

func (network *NetWork) getPacket(id byte) *protocol.DataPacket {
	packet := network.packets[id & 0xff]
	if packet != nil {
		return packet.New()
	}
	return nil
}

func (network *NetWork) registerPacket(id byte, packet protocol.DataPacket) {
	network.packets[id & 0xff] = packet
}

func (network *NetWork) registerPackets() {

}