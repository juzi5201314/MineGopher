package nbt

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type List struct {
	*NamedTag
	tags    []INamedTag
	tagType byte
}

func NewList(name string, tagType byte, tags []INamedTag) *List {
	return &List{NewNamedTag(name, ListTag, nil), tags, tagType}
}

func (list *List) read(reader *Reader) {
	list.tagType = reader.GetByte()
	var length = reader.GetInt()

	for i := int32(0); i < length && !reader.Feof(); i++ {
		var tag = GetTagById(list.tagType, "")
		tag.read(reader)
		list.tags = append(list.tags, tag)
	}
}

func (list *List) write(writer *Writer) {
	writer.PutByte(list.GetTagType())
	writer.PutInt(int32(len(list.tags)))

	for _, tag := range list.tags {
		tag.write(writer)
	}
}

func (list *List) GetTags() []INamedTag {
	return list.tags
}

func (list *List) GetTagType() byte {
	return list.tagType
}

func (list *List) GetTag(offset int) INamedTag {
	return list.tags[offset]
}

func (list *List) AddTag(tag INamedTag) error {
	if tag.GetType() != list.GetTagType() {
		return errors.New("invalid tag for list")
	}
	list.tags = append(list.tags, tag)
	return nil
}

func (list *List) Pop() INamedTag {
	var tag = list.tags[len(list.tags)-1]
	list.tags = list.tags[:len(list.tags)-2]
	return tag
}

func (list *List) Shift() INamedTag {
	var tag = list.tags[0]
	list.tags = list.tags[1:]
	return tag
}

func (list *List) DeleteAtOffset(offset int) {
	if offset > len(list.tags)-1 || offset < 0 {
		return
	}

	list.tags = append(list.tags[:offset], list.tags[offset+1:]...)
}

func (list *List) toString(nestingLevel int) string {
	var str = strings.Repeat(" ", nestingLevel*2)
	var entries = " entries"
	if len(list.tags) == 1 {
		entries = " entry"
	}

	str += "TAG_List('" + list.GetName() + " (" + GetTagName(list.tagType) + ")'): " + strconv.Itoa(len(list.tags)) + entries + "\n"
	str += strings.Repeat(" ", nestingLevel*2) + "{\n"

	for _, tag := range list.tags {
		if list, ok := tag.(*List); ok {
			str += list.toString(nestingLevel + 1)
		} else {
			if compound, ok := tag.(*Compound); ok {
				str += compound.toString(nestingLevel+1, true)
			} else {
				str += strings.Repeat(" ", (nestingLevel+1)*2)
				str += GetTagName(tag.GetType()) + "(None): " + fmt.Sprint(tag.Interface()) + "\n"
			}
		}
	}
	str += strings.Repeat(" ", nestingLevel*2) + "}\n"
	return str
}
