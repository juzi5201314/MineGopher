package nbt

import "fmt"

const (
	EndTag byte = iota
	ByteTag
	ShortTag
	IntTag
	LongTag
	FloatTag
	DoubleTag
	ByteArrayTag
	StringTag
	ListTag
	CompoundTag
	IntArrayTag
	LongArrayTag
)

var TagName = map[byte]string{
	EndTag:       "EndTag",
	ByteTag:      "ByteTag",
	ShortTag:     "ShortTag",
	IntTag:       "IntTag",
	LongTag:      "LongTag",
	FloatTag:     "FloatTag",
	DoubleTag:    "DoubleTag",
	ByteArrayTag: "ByteArrayTag",
	StringTag:    "StringTag",
	ListTag:      "ListTag",
	CompoundTag:  "CompoundTag",
	IntArrayTag:  "IntArrayTag",
	LongArrayTag: "LongArrayTag",
}

func GetTagById(tagId byte, name string) INamedTag {
	switch tagId {
	case EndTag:
		return NewEnd(name)
	case ByteTag:
		return NewByte(name, 0)
	case ShortTag:
		return NewShort(name, 0)
	case IntTag:
		return NewInt(name, 0)
	case LongTag:
		return NewLong(name, 0)
	case FloatTag:
		return NewFloat(name, 0)
	case DoubleTag:
		return NewDouble(name, 0)
	case ByteArrayTag:
		return NewByteArray(name, []byte{})
	case StringTag:
		return NewString(name, "")
	case ListTag:
		return NewList(name, ByteTag, []INamedTag{})
	case CompoundTag:
		return NewCompound(name, map[string]INamedTag{})
	case IntArrayTag:
		return NewIntArray(name, []int32{})
	case LongArrayTag:
		return NewLongArray(name, []int64{})
	}
	return nil
}

// GetTagName returns the tag name associated with the given ID.
func GetTagName(id byte) string {
	return TagName[id]
}

type INamedTag interface {
	GetType() byte
	ToString() string
	GetName() string
	setName(string)
	Interface() interface{}
	setValue(interface{})
	IsCompatibleWith(INamedTag) bool
	IsOfType(byte) bool
	read(*Reader)
	write(*Writer)
}

type Tag struct {
	tagType byte
	value   interface{}
}

type NamedTag struct {
	*Tag
	name string
}

func NewNamedTag(name string, tagType byte, value interface{}) *NamedTag {
	return &NamedTag{&Tag{tagType, value}, name}
}

func (tag *Tag) GetType() byte {
	return tag.tagType
}

func (tag *Tag) IsOfType(tagType byte) bool {
	return tag.tagType == tagType
}

func (tag *Tag) IsCompatibleWith(namedTag INamedTag) bool {
	return tag.tagType == namedTag.GetType()
}

func (tag *Tag) Interface() interface{} {
	return tag.value
}

func (tag *Tag) setValue(value interface{}) {
	tag.value = value
}

func (tag *Tag) read(*Reader) {}

func (tag *Tag) write(*Writer) {}

func (tag *NamedTag) GetName() string {
	return tag.name
}

func (tag *NamedTag) setName(name string) {
	tag.name = name
}

func (tag *NamedTag) ToString() string {
	return GetTagName(tag.GetType()) + "('" + tag.GetName() + "'): " + fmt.Sprint(tag.value) + "\n"
}
