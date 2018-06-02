package network

import (
	"fmt"
	"github.com/juzi5201314/MineGopher/api"
	"github.com/juzi5201314/MineGopher/network/protocol"
)

func New(s api.Server) *NetWork {
	network := new(NetWork)
	network.packets = map[byte]protocol.DataPacket{}
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
	return fmt.Sprint("MCPE;", network.name, ";", "201", ";", "1.2.10", ";", "0", ";", api.GetServer().GetConfig().Get("max-player", 20), ";", api.GetServer().GetRaknetServer().GetId(), ";", api.GetServer().GetName(), ";Creative;")

}

func (network *NetWork) SetName(name string) {
	network.name = name
}

func (network *NetWork) RaknetPacketToMinecraftPaket(buffer []byte) api.MinecraftPacket {
	packet := NewMinecraftPacket()
	packet.SetBuffer(buffer)
	packet.Decode()
	return packet
}

