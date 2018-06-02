package nbt

import (
	"strconv"
	"strings"
	"sync"
)

type Compound struct {
	*NamedTag
	tags  map[string]INamedTag
	mutex sync.RWMutex
}

func NewCompound(name string, tags map[string]INamedTag) *Compound {
	if tags == nil {
		tags = make(map[string]INamedTag)
	}
	return &Compound{NewNamedTag(name, CompoundTag, nil), tags, sync.RWMutex{}}
}

func (compound *Compound) read(reader *Reader) {
	for {
		var tag = reader.GetTag()
		if tag == nil || tag.GetType() == EndTag {
			return
		}
		tag.read(reader)

		compound.SetTag(tag)
	}
}

func (compound *Compound) write(writer *Writer) {
	compound.mutex.RLock()
	for _, tag := range compound.tags {
		writer.PutTag(tag)
		tag.write(writer)
	}
	compound.mutex.RUnlock()
	writer.PutTag(NewEnd(""))
}

func (compound *Compound) HasTag(name string) bool {
	compound.mutex.RLock()
	var _, exists = compound.tags[name]
	compound.mutex.RUnlock()
	return exists
}

func (compound *Compound) HasTagWithType(name string, tagType byte) bool {
	if !compound.HasTag(name) {
		return false
	}
	var tag = compound.GetTag(name)
	return tag.IsOfType(tagType)
}

func (compound *Compound) GetTag(name string) INamedTag {
	if !compound.HasTag(name) {
		return nil
	}
	compound.mutex.RLock()
	defer compound.mutex.RUnlock()
	return compound.tags[name]
}

func (compound *Compound) SetTag(tag INamedTag) {
	compound.mutex.Lock()
	compound.tags[tag.GetName()] = tag
	compound.mutex.Unlock()
}

func (compound *Compound) GetTags() map[string]INamedTag {
	compound.mutex.RLock()
	defer compound.mutex.RUnlock()
	return compound.tags
}

func (compound *Compound) SetByte(name string, value byte) {
	compound.SetTag(NewByte(name, value))
}

func (compound *Compound) GetByte(name string, defaultValue byte) byte {
	if !compound.HasTagWithType(name, ByteTag) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(byte)
}

func (compound *Compound) SetShort(name string, value int16) {
	compound.SetTag(NewShort(name, value))
}

func (compound *Compound) GetShort(name string, defaultValue int16) int16 {
	if !compound.HasTagWithType(name, ShortTag) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(int16)
}

func (compound *Compound) SetInt(name string, value int32) {
	compound.SetTag(NewInt(name, value))
}

func (compound *Compound) GetInt(name string, defaultValue int32) int32 {
	if !compound.HasTagWithType(name, IntTag) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(int32)
}

func (compound *Compound) SetLong(name string, value int64) {
	compound.SetTag(NewLong(name, value))
}

func (compound *Compound) GetLong(name string, defaultValue int64) int64 {
	if !compound.HasTagWithType(name, LongTag) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(int64)
}

func (compound *Compound) SetFloat(name string, value float32) {
	compound.SetTag(NewFloat(name, value))
}

func (compound *Compound) GetFloat(name string, defaultValue float32) float32 {
	if !compound.HasTagWithType(name, FloatTag) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(float32)
}

func (compound *Compound) SetDouble(name string, value float64) {
	compound.SetTag(NewDouble(name, value))
}

func (compound *Compound) GetDouble(name string, defaultValue float64) float64 {
	if !compound.HasTagWithType(name, DoubleTag) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(float64)
}

func (compound *Compound) SetString(name string, value string) {
	compound.SetTag(NewString(name, value))
}

func (compound *Compound) GetString(name string, defaultValue string) string {
	if !compound.HasTagWithType(name, StringTag) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(string)
}

func (compound *Compound) SetByteArray(name string, value []byte) {
	compound.SetTag(NewByteArray(name, value))
}

func (compound *Compound) GetByteArray(name string, defaultValue []byte) []byte {
	if !compound.HasTagWithType(name, ByteArrayTag) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().([]byte)
}

func (compound *Compound) SetIntArray(name string, value []int32) {
	compound.SetTag(NewIntArray(name, value))
}

func (compound *Compound) GetIntArray(name string, defaultValue []int32) []int32 {
	if !compound.HasTagWithType(name, IntArrayTag) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().([]int32)
}

func (compound *Compound) SetLongArray(name string, value []int64) {
	compound.SetTag(NewLongArray(name, value))
}

func (compound *Compound) GetLongArray(name string, defaultValue []int64) []int64 {
	if !compound.HasTagWithType(name, LongArrayTag) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().([]int64)
}

func (compound *Compound) SetList(name string, tagType byte, value []INamedTag) {
	compound.SetTag(NewList(name, tagType, value))
}

func (compound *Compound) GetList(name string, tagType byte) *List {
	if !compound.HasTagWithType(name, ListTag) {
		return nil
	}
	var list = compound.GetTag(name).(*List)
	if len(list.GetTags()) != 0 {
		if list.GetTagType() != tagType {
			return nil
		}
	}
	list.tagType = tagType

	return list
}

func (compound *Compound) SetCompound(name string, value map[string]INamedTag) {
	compound.SetTag(NewCompound(name, value))
}

func (compound *Compound) GetCompound(name string) *Compound {
	if !compound.HasTagWithType(name, CompoundTag) {
		return nil
	}
	return compound.tags[name].(*Compound)
}

func (compound *Compound) Interface() interface{} {
	compound.mutex.RLock()
	defer compound.mutex.RUnlock()
	return compound.tags
}

func (compound *Compound) toString(nestingLevel int, inList bool) string {
	var str = strings.Repeat(" ", nestingLevel*2)
	var entries = " entries"

	compound.mutex.RLock()
	if len(compound.tags) == 1 {
		entries = " entry"
	}
	compound.mutex.RUnlock()

	var name = "'" + compound.GetName() + "'"
	if inList {
		name = "None"
	}

	compound.mutex.RLock()
	str += "TAG_Compound(" + name + "): " + strconv.Itoa(len(compound.tags)) + entries + "\n"
	str += strings.Repeat(" ", nestingLevel*2) + "{\n"

	for _, tag := range compound.tags {
		if list, ok := tag.(*List); ok {
			str += list.toString(nestingLevel + 1)
		} else {
			if compound, ok := tag.(*Compound); ok {
				str += compound.toString(nestingLevel+1, false)
			} else {
				str += strings.Repeat(" ", (nestingLevel+1)*2)
				str += tag.ToString()
			}
		}
	}
	compound.mutex.RUnlock()
	str += strings.Repeat(" ", nestingLevel*2) + "}\n"
	return str
}

func (compound *Compound) ToString() string {
	return compound.toString(0, false)
}
