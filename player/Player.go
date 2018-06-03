package player

import (
	"github.com/google/uuid"
	"github.com/juzi5201314/MineGopher/api"
	"github.com/juzi5201314/MineGopher/network"
	"github.com/juzi5201314/MineGopher/network/protocol"
)

type Player struct {
	Session api.Session

	protocol    int32
	username    string
	skin_id     string
	skin_data   []byte
	cape_data   []byte
	client_UUID uuid.UUID
	client_id   int
	XUID        string
	language    string

	nameTag string
}

func (player *Player) HandlePacket(mpk api.MinecraftPacket) {
	for _, pk := range mpk.GetPackets() {
		switch pk.(type) {
		case *protocol.LoginPacket:
			if _, in := protocol.Protocols[pk.(*protocol.LoginPacket).Protocol]; !in {
				player.Close()
				return
			}
			player.onLogin(pk.(*protocol.LoginPacket))
			break
		}
	}
}

func (player *Player) onLogin(packet *protocol.LoginPacket) {
	player.username = packet.Username
	player.skin_id = packet.SkinId
	player.skin_data = packet.SkinData
	player.cape_data = packet.CapeData
	player.protocol = packet.Protocol
	player.client_id = packet.ClientId
	player.client_UUID = packet.ClientUUID
	player.XUID = packet.ClientXUID
	player.language = packet.Language
	player.nameTag = packet.Username
	player.Spawn()

	pk := &protocol.PlayStatusPacket{protocol.NewPacket(protocol.GetPacketId(protocol.PLAY_STATUS_PACKET)), 0}
	player.SendPacket(pk)
}

func (player *Player) SendPacket(packet protocol.DataPacket) {
	pk := network.NewMinecraftPacket()
	pk.AddPacket(packet)
	player.SendBatch(pk)
}

func (player *Player) SendBatch(packet *network.MinecraftPacket) {
	player.Session.SendPacket(packet, 2, 3)
}

func (player *Player) Spawn() {

}

func (player *Player) Close() {
	player.Session.Close()
}

func (player *Player) GetNameTag() string {
	return player.nameTag
}
