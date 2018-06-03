package entity

import (
	"github.com/google/uuid"
	"github.com/juzi5201314/MineGopher/api/server"
	"github.com/juzi5201314/MineGopher/network"
	"github.com/juzi5201314/MineGopher/network/protocol"
	"github.com/juzi5201314/MineGopher/network/protocol/types"
	"github.com/juzi5201314/MineGopher/level/chunk"
	"github.com/golang/geo/r3"
	"github.com/juzi5201314/MineGopher/math"
	"github.com/juzi5201314/MineGopher/api"
	"github.com/juzi5201314/MineGopher/level"
)

type Player struct {
	*Entity
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

	viewDistance int32
	needLoadChunk bool

	chunkLoader *level.ChunkLoader
}

func (player *Player) HandlePacket(mpk interface{
	GetPackets() []protocol.DataPacket
}) {
	for _, pk := range mpk.GetPackets() {
		switch pk.(type) {
		case *protocol.LoginPacket:
			if _, in := protocol.Protocols[pk.(*protocol.LoginPacket).Protocol]; !in {
				player.Close()
				return
			}
			player.onLogin(pk.(*protocol.LoginPacket))
			break
		case *protocol.ResourcePackClientResponsePacket:
			switch pk.(*protocol.ResourcePackClientResponsePacket).Status {
			case 1:
				//Refused
				break
			case 2:
				//SendPacks
				break
			case 3:
				//HaveAllPacks
				pk := &protocol.ResourcePackStackPacket{protocol.NewPacket(protocol.GetPacketId(protocol.RESOURCE_PACK_STACK_PACKET)), false, []types.ResourcePackInfoEntry{}, []types.ResourcePackInfoEntry{}}
				player.SendPacket(pk)
				break
			case 4:
				//Completed
				println(2333)
				server.GetServer().GetDefaultLevel().GetDimension().LoadChunk(0, 0, func(chunk *chunk.Chunk) {
					server.GetServer().GetDefaultLevel().GetDimension().AddEntity(player.Entity, r3.Vector{0, 40, 0})
					pk := &protocol.StartGamePacket{Packet: protocol.NewPacket(protocol.GetPacketId(protocol.START_GAME_PACKET))}
					pk.Generator = 1
					pk.LevelSeed = 312402
					pk.TrustPlayers = true
					pk.DefaultPermissionLevel = 1
					pk.EntityRuntimeId = player.GetEid()
					pk.EntityUniqueId = player.GetUniqueId()
					pk.PlayerGameMode = 1
					pk.PlayerPosition = player.GetPosition()
					pk.LevelGameMode = 1
					pk.LevelSpawnPosition = math.NewPosition(0, 40, 0)
					pk.CommandsEnabled = true
					var gameRules = server.GetServer().GetDefaultLevel().GetGameRules()
					var gameRuleEntries = map[string]types.GameRuleEntry{}
					for name, gameRule := range gameRules {
						gameRuleEntries[string(name)] = types.GameRuleEntry{Name: string(gameRule.GetName()), Value: gameRule.GetValue()}
					}
					pk.GameRules = gameRuleEntries
					pk.LevelName = server.GetServer().GetDefaultLevel().GetName()
					pk.CurrentTick = 0
					pk.Time = 0
					pk.AchievementsDisabled = true
					pk.BroadcastToXbox = true
					pk.BroadcastToLan = true
					player.SendPacket(pk)
					pk2 := &protocol.CraftingDataPacket{protocol.NewPacket(protocol.GetPacketId(protocol.CRAFTING_DATA_PACKET))}
					player.SendPacket(pk2)

				})
				break
			}
			break
		case *protocol.RequestChunkRadiusPacket:
			player.viewDistance = pk.(*protocol.RequestChunkRadiusPacket).Radius
			cpk := &protocol.ChunkRadiusUpdatedPacket{protocol.NewPacket(protocol.GetPacketId(protocol.CHUNK_RADIUS_UPDATED_PACKET)), player.viewDistance}
			player.SendPacket(cpk)

			haschunk := player.needLoadChunk
			player.needLoadChunk = true

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

	player.chunkLoader = level.NewChunkLoader(nil, 0, 0)
	player.chunkLoader.OnLoad = func(chunk *chunk.Chunk) {

	}

	pk := &protocol.PlayStatusPacket{protocol.NewPacket(protocol.GetPacketId(protocol.PLAY_STATUS_PACKET)), 0}
	player.SendPacket(pk)
	repk := &protocol.ResourcePackInfoPacket{protocol.NewPacket(protocol.GetPacketId(protocol.RESOURCE_PACKS_INFO_PACKET)), false, []types.ResourcePackInfoEntry{}, []types.ResourcePackInfoEntry{}}
	player.SendPacket(repk)
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

func (player *Player) GetName() string {
	return player.username
}

func (player *Player) GetUUID() uuid.UUID {
	return player.client_UUID
}

func (player *Player) GetXUID() string {
	return player.XUID
}

func (player *Player) Tick() {

	player.Entity.Tick()
}
