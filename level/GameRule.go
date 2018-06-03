package level

import (
	"reflect"
)

const (
	GameRuleCommandBlockOutput  GameRuleName = "commandblockoutput"
	GameRuleDoDaylightCycle     GameRuleName = "dodaylightcycle"
	GameRuleDoEntityDrops       GameRuleName = "doentitydrops"
	GameRuleDoFireTick          GameRuleName = "dofiretick"
	GameRuleDoMobLoot           GameRuleName = "domobloot"
	GameRuleDoMobSpawning       GameRuleName = "domobspawning"
	GameRuleDoTileDrops         GameRuleName = "dotiledrops"
	GameRuleDoWeatherCycle      GameRuleName = "doweathercycle"
	GameRuleDrowningDamage      GameRuleName = "drowningdamage"
	GameRuleFallDamage          GameRuleName = "falldamage"
	GameRuleFireDamage          GameRuleName = "firedamage"
	GameRuleKeepInventory       GameRuleName = "keepinventory"
	GameRuleMobGriefing         GameRuleName = "mobgriefing"
	GameRuleNaturalRegeneration GameRuleName = "naturalregeneration"
	GameRulePvp                 GameRuleName = "pvp"
	GameRuleSendCommandFeedback GameRuleName = "sendcommandfeedback"
	GameRuleShowCoordinates     GameRuleName = "showcoordinates"
	GameRuleRandomTickSpeed     GameRuleName = "randomtickspeed"
	GameRuleTntExplodes         GameRuleName = "tntexplodes"
)

type GameRuleName string

type GameRule struct {
	name  GameRuleName
	value interface{}
}

func NewGameRule(name GameRuleName, value interface{}) *GameRule {
	return &GameRule{name, value}
}

func (rule *GameRule) GetName() GameRuleName {
	return rule.name
}

func (rule *GameRule) GetValue() interface{} {
	return rule.value
}

func (rule *GameRule) SetValue(value interface{}) bool {
	if reflect.TypeOf(value).Kind() != reflect.TypeOf(rule.value).Kind() {
		return false
	}
	rule.value = value
	return true
}
