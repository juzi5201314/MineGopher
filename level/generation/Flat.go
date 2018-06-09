package generation

import (
	"github.com/juzi5201314/MineGopher/level/chunk"
)

type Flat struct{}

func (flat Flat) GetName() string {
	return "Flat"
}

func (flat Flat) GenerateNewChunk(x, z int32) *chunk.Chunk {
	var c = chunk.New(x, z)
	var y int
	for x := 0; x < 16; x++ {
		for z := 0; z < 16; z++ {
			y = 0
			c.SetBlockId(x, y, z, 7)
			y++
			c.SetBlockId(x, y, z, 3)
			y++
			c.SetBlockId(x, y, z, 3)
			y++
			c.SetBlockId(x, y, z, 57)

			for i := y - 1; i >= 0; i-- {
				c.SetSkyLight(x, y, z, 0)
			}
		}
	}
	c.RecalculateHeightMap()
	return c
}
