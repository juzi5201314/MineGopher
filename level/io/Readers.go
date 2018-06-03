package io

import (
	"github.com/juzi5201314/MineGopher/nbt"
	"github.com/juzi5201314/MineGopher/level/chunk"
)

func GetAnvilChunkFromNBT(compound *nbt.Compound) *chunk.Chunk {
	var level = compound.GetCompound("Level")
	var c = chunk.New(level.GetInt("xPos", 0), level.GetInt("zPos", 0))
	c.LightPopulated = getBool(level.GetByte("LightPopulated", 0))
	c.TerrainPopulated = getBool(level.GetByte("TerrainPopulated", 0))
	c.Biomes = level.GetByteArray("Biomes", make([]byte, 256))
	c.InhabitedTime = level.GetLong("InhabitedTime", 0)
	c.LastUpdate = level.GetLong("LastUpdate", 0)
	for i, b := range level.GetByteArray("HeightMap", make([]byte, 256)) {
		c.HeightMap[i] = int16(b)
	}

	var sections = level.GetList("Sections", nbt.CompoundTag)
	if sections == nil {
		return c
	}
	for _, comp := range sections.GetTags() {
		section := comp.(*nbt.Compound)
		subChunk := chunk.NewSubChunk()
		subChunk.BlockLight = reorderNibbleArray(section.GetByteArray("BlockLight", make([]byte, 2048)))
		subChunk.BlockData = reorderNibbleArray(section.GetByteArray("Data", make([]byte, 2048)))
		subChunk.SkyLight = reorderNibbleArray(section.GetByteArray("SkyLight", make([]byte, 2048)))
		subChunk.BlockIds = reorderBlocks(section.GetByteArray("Blocks", make([]byte, 4096)))

		c.SetSubChunk(section.GetByte("Y", 0), subChunk)
	}

	return c
}

func reorderBlocks(blocks []byte) []byte {
	var data = make([]byte, 4096)
	var i = 0
	for x := 0; x < 16; x++ {
		var zM = x + 256
		for z := x; z < zM; z += 16 {
			var yM = z + 4096
			for y := z; y < yM; y += 256 {
				data[i] = blocks[y]
				i++
			}
		}
	}
	return data
}

func reorderNibbleArray(arr []byte) []byte {
	var data = make([]byte, 2048)
	var i = 0
	for x := 0; x < 8; x++ {
		for z := 0; z < 16; z++ {
			var zx = (z << 3) | x
			for y := 0; y < 8; y++ {
				var j = (y << 8) | zx
				var j80 = j | 0x80
				if arr[j] != 0 || arr[j80] != 0 {
					data[i] = (arr[j80] << 4) | (arr[j] & 0x0f)
					data[i|80] = (arr[j] >> 4) | (arr[j80] & 0xf0)
				}
				i++
			}
		}
		i += 128
	}
	return data
}

func getBool(value byte) bool {
	if value > 0 {
		return true
	}
	return false
}
