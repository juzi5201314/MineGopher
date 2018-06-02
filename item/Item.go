package item

import (
	"github.com/juzi5201314/MineGopher/nbt"
)

func Get(id int, damage int, count int8) *Item {
	return &Item{
		id:     id,
		damage: damage,
		count:  count,
	}
}

type Item struct {
	id         int
	damage     int
	count      int8
	name       string
	customName string
	lore       []string
	NBT        *nbt.Compound
}

func (item *Item) GetId() int {
	return item.id
}

func (item *Item) GetCount() int8 {
	return item.count
}

func (item *Item) GetDamage() int {
	return item.damage
}

func (item *Item) NBTParse(compound *nbt.Compound) {
	if compound.HasTagWithType("display", nbt.CompoundTag) {
		item.customName = compound.GetCompound("display").GetString("Name", item.name)
		for _, tag := range compound.GetCompound("display").GetList("Lore", nbt.StringTag).GetTags() {
			item.lore = append(item.lore, tag.Interface().(string))
		}
	}
	item.NBT = compound
}

func (item *Item) NBTEmit() *nbt.Compound {
	compound := item.NBT
	compound.SetCompound("display", make(map[string]nbt.INamedTag))
	if item.customName != "" {
		compound.GetCompound("display").SetString("Name", item.customName)
		var list []nbt.INamedTag
		for _, lore := range item.lore {
			list = append(list, nbt.NewString("", lore))
		}
		compound.GetCompound("display").SetList("Lore", nbt.StringTag, list)
	}
	return compound
}
