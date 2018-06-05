package network

import (
	"fmt"
	apinetwork "github.com/juzi5201314/MineGopher/api/network"
	"github.com/juzi5201314/MineGopher/api/server"
	"github.com/juzi5201314/MineGopher/network/protocol"
)

func New() *NetWork {
	network := new(NetWork)
	network.packets = map[byte]func() protocol.DataPacket{}
	network.RegisterPackets()
	return network
}

type NetWork struct {
	packets  map[byte]func() protocol.DataPacket
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
	return fmt.Sprint("MCPE;", network.name, ";", protocol.PROTOCOL, ";", protocol.VERSION, ";", "0", ";", server.GetServer().GetConfig().Get("max-player", 20), ";", server.GetServer().GetRaknetServer().GetId(), ";", server.GetServer().GetName(), ";Creative;")

}

func (network *NetWork) SetName(name string) {
	network.name = name
}

func (network *NetWork) RaknetPacketToMinecraftPaket(buffer []byte) apinetwork.MinecraftPacket {
	packet := NewMinecraftPacket()
	packet.SetBuffer(buffer)
	packet.Decode()
	return packet
}

func (network *NetWork) RegisterPacket(id byte, fn func() protocol.DataPacket) {
	network.packets[id] = fn
}

func (network *NetWork) GetPacket(id byte) (func() protocol.DataPacket, bool) {
	p, has := network.packets[id]
	return p, has
}

func (network *NetWork) RegisterPackets() {
	network.RegisterPacket(protocol.GetPacketId(protocol.LOGIN_PACKET), func() protocol.DataPacket {
		return &protocol.LoginPacket{Packet: protocol.NewPacket(protocol.GetPacketId(protocol.LOGIN_PACKET))}
	})
	network.RegisterPacket(protocol.GetPacketId(protocol.RESOURCE_PACK_CLIENT_RESPONSE_PACKET), func() protocol.DataPacket {
		return &protocol.ResourcePackClientResponsePacket{Packet: protocol.NewPacket(protocol.GetPacketId(protocol.RESOURCE_PACK_CLIENT_RESPONSE_PACKET))}
	})
	network.RegisterPacket(protocol.GetPacketId(protocol.ADD_ENTITY_PACKET), func() protocol.DataPacket {
		return &protocol.AddEntityPacket{Packet: protocol.NewPacket(protocol.GetPacketId(protocol.ADD_ENTITY_PACKET))}
	})
	network.RegisterPacket(protocol.GetPacketId(protocol.REMOVE_ENTITY_PACKET), func() protocol.DataPacket {
		return &protocol.RemoveEntityPacket{Packet: protocol.NewPacket(protocol.GetPacketId(protocol.REMOVE_ENTITY_PACKET))}
	})
	network.RegisterPacket(protocol.GetPacketId(protocol.START_GAME_PACKET), func() protocol.DataPacket {
		return &protocol.StartGamePacket{Packet: protocol.NewPacket(protocol.GetPacketId(protocol.START_GAME_PACKET))}
	})
	network.RegisterPacket(protocol.GetPacketId(protocol.CRAFTING_DATA_PACKET), func() protocol.DataPacket {
		return &protocol.CraftingDataPacket{Packet: protocol.NewPacket(protocol.GetPacketId(protocol.CRAFTING_DATA_PACKET))}
	})
	network.RegisterPacket(protocol.GetPacketId(protocol.CHUNK_RADIUS_UPDATED_PACKET), func() protocol.DataPacket {
		return &protocol.ChunkRadiusUpdatedPacket{Packet: protocol.NewPacket(protocol.GetPacketId(protocol.CHUNK_RADIUS_UPDATED_PACKET))}
	})
	network.RegisterPacket(protocol.GetPacketId(protocol.REQUEST_CHUNK_RADIUS_PACKET), func() protocol.DataPacket {
		return &protocol.RequestChunkRadiusPacket{Packet: protocol.NewPacket(protocol.GetPacketId(protocol.REQUEST_CHUNK_RADIUS_PACKET))}
	})
	network.RegisterPacket(protocol.GetPacketId(protocol.ADD_PLAYER_PACKET), func() protocol.DataPacket {
		return &protocol.AddPlayerPacket{Packet: protocol.NewPacket(protocol.GetPacketId(protocol.ADD_PLAYER_PACKET))}
	})
	network.RegisterPacket(protocol.GetPacketId(protocol.MOVE_PLAYER_PACKET), func() protocol.DataPacket {
		return &protocol.MovePlayerPacket{Packet: protocol.NewPacket(protocol.GetPacketId(protocol.MOVE_PLAYER_PACKET))}
	})
}
