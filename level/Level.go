package level

import (
	"github.com/juzi5201314/MineGopher/level/generation"
	"github.com/juzi5201314/MineGopher/level/providers"
	"os"
)

type Level struct {
	path      string
	name      string
	dimension *Dimension
	gameRules map[GameRuleName]*GameRule
}

func NewLevel(path string, name string) *Level {
	os.MkdirAll(path, 0700)
	os.MkdirAll(path+"/region", 0700)
	level := new(Level)
	level.name = name
	level.path = path
	level.dimension = NewDimension(level)
	level.dimension.SetChunkProvider(providers.NewAnvil(path + "/region"))
	level.dimension.SetGenerator(generation.Flat{})
	level.gameRules = map[GameRuleName]*GameRule{}
	level.initGameRules()
	return level
}

func (level *Level) GetName() string {
	return level.name
}

func (level *Level) GetDimension() *Dimension {
	return level.dimension
}

func (level *Level) GetGameRule(gameRule GameRuleName) *GameRule {
	return level.gameRules[gameRule]
}

func (level *Level) GetGameRules() map[GameRuleName]*GameRule {
	return level.gameRules
}

func (level *Level) AddGameRule(rule *GameRule) {
	level.gameRules[rule.GetName()] = rule
}

func (level *Level) initGameRules() {
	level.AddGameRule(NewGameRule(GameRuleCommandBlockOutput, true))
	level.AddGameRule(NewGameRule(GameRuleDoDaylightCycle, true))
	level.AddGameRule(NewGameRule(GameRuleDoEntityDrops, true))
	level.AddGameRule(NewGameRule(GameRuleDoFireTick, true))
	level.AddGameRule(NewGameRule(GameRuleDoMobLoot, true))
	level.AddGameRule(NewGameRule(GameRuleDoMobSpawning, true))
	level.AddGameRule(NewGameRule(GameRuleDoTileDrops, true))
	level.AddGameRule(NewGameRule(GameRuleDoWeatherCycle, true))
	level.AddGameRule(NewGameRule(GameRuleDrowningDamage, true))
	level.AddGameRule(NewGameRule(GameRuleFallDamage, true))
	level.AddGameRule(NewGameRule(GameRuleFireDamage, true))
	level.AddGameRule(NewGameRule(GameRuleKeepInventory, false))
	level.AddGameRule(NewGameRule(GameRuleMobGriefing, true))
	level.AddGameRule(NewGameRule(GameRuleNaturalRegeneration, true))
	level.AddGameRule(NewGameRule(GameRulePvp, true))
	level.AddGameRule(NewGameRule(GameRuleSendCommandFeedback, true))
	level.AddGameRule(NewGameRule(GameRuleShowCoordinates, true))
	level.AddGameRule(NewGameRule(GameRuleRandomTickSpeed, uint32(3)))
	level.AddGameRule(NewGameRule(GameRuleTntExplodes, true))
}
