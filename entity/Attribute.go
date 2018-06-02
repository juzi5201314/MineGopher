package entity

import "math"

// AttributeName is the name used for an attribute in Minecraft.
type AttributeName string

const (
	AttributeHealth              AttributeName = "minecraft:health"
	AttributeMovementSpeed       AttributeName = "minecraft:movement"
	AttributeAttackDamage        AttributeName = "minecraft:attack_damage"
	AttributeAbsorption          AttributeName = "minecraft:absorption"
	AttributeHunger              AttributeName = "minecraft:hunger"
	AttributeSaturation          AttributeName = "minecraft:saturation"
	AttributeExhaustion          AttributeName = "minecraft:exhaustion"
	AttributeKnockBackResistance AttributeName = "minecraft:knockback_resistance"
	AttributeFollowRange         AttributeName = "minecraft:follow_range"
	AttributeExperienceLevel     AttributeName = "minecraft:player.level"
	AttributeExperience          AttributeName = "minecraft:player.experience"
	AttributeJumpStrength        AttributeName = "minecraft:horse.jump_strength"
)

// Attribute is a struct containing data of an entity property.
type Attribute struct {
	name         AttributeName
	MinValue     float32
	MaxValue     float32
	Value        float32
	DefaultValue float32
}

// AttributeMap is a struct containing an unlimited amount of attributes.
type AttributeMap map[AttributeName]*Attribute

// NewAttributeMap returns a new attribute map with default attributes.
func NewAttributeMap() AttributeMap {
	return AttributeMap{
		AttributeHealth:              NewAttribute(AttributeHealth, 20, 1024),
		AttributeMovementSpeed:       NewAttribute(AttributeMovementSpeed, 0.1, 1024),
		AttributeAttackDamage:        NewAttribute(AttributeAttackDamage, 2, 2048),
		AttributeAbsorption:          NewAttribute(AttributeAbsorption, 0, 1024),
		AttributeHunger:              NewAttribute(AttributeHunger, 20, 20),
		AttributeSaturation:          NewAttribute(AttributeSaturation, 20, 20),
		AttributeExhaustion:          NewAttribute(AttributeExhaustion, 0, 5),
		AttributeKnockBackResistance: NewAttribute(AttributeKnockBackResistance, 0, 1),
		AttributeFollowRange:         NewAttribute(AttributeFollowRange, 32, 2048),
		AttributeExperience:          NewAttribute(AttributeExperience, 0, 1),
		AttributeExperienceLevel:     NewAttribute(AttributeExperienceLevel, 0, math.MaxInt32),
		AttributeJumpStrength:        NewAttribute(AttributeJumpStrength, 0.7, 2),
	}
}

// NewAttribute returns a new Attribute with the given name.
func NewAttribute(name AttributeName, value, maxValue float32) *Attribute {
	return &Attribute{name, 0, maxValue, value, value}
}

// GetName returns the name of the attribute.
func (attribute *Attribute) GetName() AttributeName {
	return attribute.name
}

// Exists checks if an attribute with the given name exists.
func (attMap AttributeMap) Exists(name AttributeName) bool {
	var _, ok = attMap[name]
	return ok
}

// SetAttribute sets an attribute in this attribute map.
func (attMap AttributeMap) SetAttribute(attribute *Attribute) {
	attMap[attribute.GetName()] = attribute
}

// GetAttribute returns an attribute with the given name.
func (attMap AttributeMap) GetAttribute(name AttributeName) *Attribute {
	return attMap[name]
}
