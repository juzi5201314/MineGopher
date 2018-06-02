package block

import (
	"github.com/juzi5201314/MineGopher/nbt"
)

type IBlock interface {
	GetId() byte
	GetData() byte
	SetData(byte)
	GetNBT() *nbt.Compound
	SetNBT(*nbt.Compound)
}

type Block struct {
	name      string
	runtimeId int32
	data      int32
	nbt       *nbt.Compound
}

func New(name string, runtimeId int32, data int32) *Block {
	return &Block{name: name, runtimeId: runtimeId, data: data}
}

func (block *Block) GetNBT() *nbt.Compound {
	return block.nbt
}

func (block *Block) SetNBT(nbt *nbt.Compound) {
	block.nbt = nbt

}

func (block *Block) GetName() string {
	return block.name
}

func (block *Block) GetRuntimeId() int32 {
	return block.runtimeId
}

func (block *Block) GetData() int32 {
	return block.data
}

func (block *Block) SetData(data int32) {
	block.data = data
}

func (block *Block) GetPersistentId() *nbt.Compound {
	return nbt.NewCompound("", map[string]nbt.INamedTag{
		"name": nbt.NewString("name", block.GetName()),
		"val":  nbt.NewInt("val", block.data),
	})
}
