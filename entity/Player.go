package entity

import (
	"github.com/golang/geo/r3"
	"github.com/google/uuid"
	"github.com/juzi5201314/MineGopher/api"
	"github.com/juzi5201314/MineGopher/api/player"
	"github.com/juzi5201314/MineGopher/api/server"
	"github.com/juzi5201314/MineGopher/level"
	"github.com/juzi5201314/MineGopher/level/chunk"
	"github.com/juzi5201314/MineGopher/math"
	"github.com/juzi5201314/MineGopher/network"
	"github.com/juzi5201314/MineGopher/network/protocol"
	"github.com/juzi5201314/MineGopher/network/protocol/types"
	gfmath "math"
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
	nameTag     string
	platform    int

	viewDistance  int32
	needLoadChunk bool
	OnGround      bool

	chunkLoader *level.ChunkLoader
}

func (player *Player) HandlePacket(mpk interface {
	GetPackets() []protocol.DataPacket
}) {
	for _, pk := range mpk.GetPackets() {
		switch pk.(type) {
		case *protocol.LoginPacket:
			if _, in := protocol.Protocols[pk.(*protocol.LoginPacket).Protocol]; !in {
				server.GetServer().GetLogger().Debug("protocol: ", pk.(*protocol.LoginPacket).Protocol)
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
				server.GetServer().GetDefaultLevel().GetDimension().LoadChunk(0, 0, func(chunk *chunk.Chunk) {
					player.Entity.Dimension = server.GetServer().GetDefaultLevel().GetDimension()
					server.GetServer().GetDefaultLevel().GetDimension().AddEntity(player.Entity, r3.Vector{0, 40, 0})
					pk := &protocol.StartGamePacket{Packet: protocol.NewPacket(protocol.GetPacketId(protocol.START_GAME_PACKET))}
					pk.Generator = 1
					pk.LevelSeed = 1
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
			if !haschunk {
				for _, p := range server.GetServer().GetAllPlayer() {
					//p.SpawnTo(player)
					_ = p
				}
				//player.SpawnToAll()
				//player.UpdateAttributes()
				pk := &protocol.PlayStatusPacket{protocol.NewPacket(protocol.GetPacketId(protocol.PLAY_STATUS_PACKET)), 3}
				player.SendPacket(pk)
			}
			break
		case *protocol.MovePlayerPacket:
			movepk := pk.(*protocol.MovePlayerPacket)
			player.Move(movepk.Position.X, movepk.Position.Y, movepk.Position.Z, movepk.Rotation.Pitch, movepk.Rotation.Yaw, movepk.Rotation.HeadYaw, movepk.OnGround)
			for _, v := range player.GetViewers() {
				v.SendPacket(pk.(*protocol.MovePlayerPacket))
			}
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
	player.platform = packet.ClientData.DeviceOS

	player.chunkLoader = level.NewChunkLoader(nil, 0, 0)
	player.chunkLoader.OnLoad = func(chunk *chunk.Chunk) {
		cpk := &protocol.FullChunkDataPacket{protocol.NewPacket(protocol.GetPacketId(protocol.FULL_CHUNK_DATA_PACKET)), chunk.X, chunk.Z, chunk.ToBinary()}
		player.SendPacket(cpk)

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

func (player *Player) SpawnTo(p player.Player) {
	pk := &protocol.AddPlayerPacket{Packet: protocol.NewPacket(protocol.GetPacketId(protocol.ADD_PLAYER_PACKET))}
	pk.UUID = player.client_UUID
	pk.DisplayName = player.username
	pk.Username = player.username
	pk.EntityRuntimeId = player.GetEid()
	pk.EntityUniqueId = player.GetUniqueId()
	pk.Position = player.GetPosition()
	pk.Rotation = player.GetRotation()
	pk.Platform = int32(player.platform)
	pk.Motion = player.GetMotion()
	player.Entity.SpawnTo(p)
	player.SendPacket(pk)
}

func (playe *Player) SpawnToAll() {
	for _, p := range playe.viewers {
		if p.GetUUID() == playe.GetUUID() {
			continue
		}
		playe.SpawnTo(p.(player.Player))
	}
}

func (player *Player) Tick() {
	if player.needLoadChunk && !player.closed {
		player.chunkLoader.Warp(player.Dimension, int32(gfmath.Floor(player.Position.X))>>4, int32(gfmath.Floor(player.Position.Z))>>4)
		player.chunkLoader.Request(player.viewDistance, 3)
	}
	player.Entity.Tick()
}

func (player *Player) Move(x, y, z float64, pitch, yaw, headYaw float64, onGround bool) {
	newChunk, _ := player.Dimension.GetChunk(int32(gfmath.Floor(x))>>4, int32(gfmath.Floor(z))>>4)
	//println(int32(gfmath.Floor(x))>>4, int32(gfmath.Floor(z))>>4)
	if player.GetChunk() != newChunk {
		player.GetChunk().RemoveViewer(player)
		/*
			newChunk.AddViewer(player)
			for _, entity := range newChunk.GetEntities() {
				e := entity.(protocol.AddEntityEntry)
				pk := &protocol.AddEntityPacket{Packet: protocol.NewPacket(protocol.GetPacketId(protocol.ADD_ENTITY_PACKET))}
				pk.UniqueId = e.GetUniqueId()
				pk.RuntimeId = e.GetEid()
				pk.EntityType = e.GetId()
				pk.Position = e.GetPosition()
				pk.Motion = e.GetMotion()
				pk.Rotation = e.GetRotation()
				pk.Attributes = e.GetAttributeMap()
				pk.EntityData = e.GetEntityData()
				player.SendPacket(pk)
			}
		*/
	}
	player.SetPosition(r3.Vector{x, y, z})
	player.Rotation.Pitch += float64(pitch)
	player.Rotation.Yaw += float64(yaw)
	player.Rotation.HeadYaw = float64(headYaw)
	player.OnGround = onGround
}
