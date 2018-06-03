package protocol

import (
	"github.com/golang/geo/r3"
	"github.com/juzi5201314/MineGopher/network/protocol/types"
	"encoding/base64"
	"github.com/juzi5201314/MineGopher/math"
)

type StartGamePacket struct {
	*Packet
	EntityUniqueId         int64
	EntityRuntimeId        uint64
	PlayerGameMode         int32
	PlayerPosition         r3.Vector
	Yaw                    float32
	Pitch                  float32
	LevelSeed              int32
	Dimension              int32
	Generator              int32
	LevelGameMode          int32
	Difficulty             int32
	LevelSpawnPosition     math.Position
	AchievementsDisabled   bool
	Time                   int32
	EduMode                bool
	RainLevel              float32
	LightningLevel         float32
	MultiPlayerGame        bool
	BroadcastToLan         bool
	BroadcastToXbox        bool
	CommandsEnabled        bool
	ForcedResourcePacks    bool
	GameRules              map[string]types.GameRuleEntry
	BonusChest             bool
	StartMap               bool
	TrustPlayers           bool
	DefaultPermissionLevel int32
	XBOXBroadcastMode      int32
	XBOXBroadcastIntent    bool
	LevelName              string
	IsTrial                bool
	CurrentTick            int64
	EnchantmentSeed        int32
	ServerChunkTickRange   int32
	HasPlatformBroadcast   bool
	PlatformBroadcastMode  uint32
}


func (pk *StartGamePacket) Encode() {
	pk.PutEntityUniqueId(pk.EntityUniqueId)
	pk.PutEntityRuntimeId(pk.EntityRuntimeId)

	pk.PutVarInt(pk.PlayerGameMode)

	pk.PutVector(pk.PlayerPosition)

	pk.PutLittleFloat(pk.Yaw)
	pk.PutLittleFloat(pk.Pitch)

	pk.PutVarInt(pk.LevelSeed)
	pk.PutVarInt(pk.Dimension)
	pk.PutVarInt(pk.Generator)
	pk.PutVarInt(pk.LevelGameMode)
	pk.PutVarInt(pk.Difficulty)

	pk.PutBlockPosition(pk.LevelSpawnPosition)
	pk.PutBool(pk.AchievementsDisabled)
	pk.PutVarInt(pk.Time)
	pk.PutBool(pk.EduMode)

	pk.PutLittleFloat(pk.RainLevel)
	pk.PutLittleFloat(pk.LightningLevel)

	pk.PutBool(pk.MultiPlayerGame)
	pk.PutBool(pk.BroadcastToLan)
	pk.PutBool(pk.BroadcastToXbox)
	pk.PutBool(pk.CommandsEnabled)
	pk.PutBool(pk.ForcedResourcePacks)

	pk.PutGameRules(pk.GameRules)

	pk.PutBool(pk.BonusChest)
	pk.PutBool(pk.StartMap)
	pk.PutBool(pk.TrustPlayers)
	pk.PutVarInt(pk.DefaultPermissionLevel)
	pk.PutVarInt(pk.XBOXBroadcastMode)
	pk.PutLittleInt(pk.ServerChunkTickRange)
	pk.PutBool(pk.HasPlatformBroadcast)
	pk.PutUnsignedVarInt(pk.PlatformBroadcastMode)
	pk.PutBool(pk.XBOXBroadcastIntent)

	pk.PutString(base64.RawStdEncoding.EncodeToString([]byte(pk.LevelName)))
	pk.PutString(pk.LevelName)
	pk.PutString("")
	pk.PutBool(pk.IsTrial)
	pk.PutLittleLong(pk.CurrentTick)
	pk.PutVarInt(pk.EnchantmentSeed)
}

func (pk *StartGamePacket) Decode() {

}