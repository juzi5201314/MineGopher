package item

var ItemName = map[int]string{
	Air: "air",
}

func GetItemName(id int) string {
	return ItemName[id]
}
