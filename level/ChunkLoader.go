package level

import (
	"sync"
	"github.com/juzi5201314/MineGopher/level/chunk"
)

type ChunkLoader struct {
	Dimension *Dimension
	ChunkX    int32
	ChunkZ    int32

	// UnloadFunction gets called for every chunk that gets unloaded in this ChunkLoader.
	// chunk that exceed the request range automatically get unloaded during request.
	OnUnload func(*chunk.Chunk)
	// LoadFunction gets called for every chunk loaded by this ChunkLoader.
	OnLoad func(*chunk.Chunk)

	mutex        sync.RWMutex
	loadedchunk map[int]*chunk.Chunk
}

// NewChunkLoader returns a new ChunkLoader on the given dimension with the given chunk X and Z.
func NewChunkLoader(dimension *Dimension, x, z int32) *ChunkLoader {
	return &ChunkLoader{dimension, x, z, func(chunk *chunk.Chunk) {}, func(chunk *chunk.Chunk) {}, sync.RWMutex{}, make(map[int]*chunk.Chunk)}
}

// Move moves the ChunkLoader to the given chunk X and Z.
func (ChunkLoader *ChunkLoader) Move(chunkX, chunkZ int32) {
	ChunkLoader.ChunkX = chunkX
	ChunkLoader.ChunkZ = chunkZ
}

// Warp warps the ChunkLoader to the given dimension and moves it to the given chunk X and Z.
func (ChunkLoader *ChunkLoader) Warp(dimension *Dimension, chunkX, chunkZ int32) {
	ChunkLoader.Dimension = dimension
	ChunkLoader.Move(chunkX, chunkZ)
}

// GetLoadedChunkCount returns the count of the loaded chunk.
func (ChunkLoader *ChunkLoader) GetLoadedChunkCount() int {
	return len(ChunkLoader.loadedchunk)
}

// HasChunkInUse checks if the ChunkLoader has a chunk with the given chunk X and Z in use.
func (ChunkLoader *ChunkLoader) HasChunkInUse(chunkX, chunkZ int32) bool {
	ChunkLoader.mutex.RLock()
	var _, ok = ChunkLoader.loadedchunk[GetChunkIndex(chunkX, chunkZ)]
	ChunkLoader.mutex.RUnlock()
	return ok
}

// setChunkInUse sets the given chunk in use.
func (ChunkLoader *ChunkLoader) setChunkInUse(chunkX, chunkZ int32, chunk *chunk.Chunk) {
	ChunkLoader.mutex.Lock()
	ChunkLoader.loadedchunk[GetChunkIndex(chunkX, chunkZ)] = chunk
	ChunkLoader.mutex.Unlock()
}

// Request requests all chunk within the given view distance from the current position.
// All chunk loaded will run the load function of this ChunkLoader.
// Request will also unload any unused chunk beyond the distance specified.
func (ChunkLoader *ChunkLoader) Request(distance int32, maximumchunk int) {
	var f = func(chunk *chunk.Chunk) {
		ChunkLoader.setChunkInUse(chunk.X, chunk.Z, chunk)
		ChunkLoader.OnLoad(chunk)
	}
	i := 0
	for x := -distance + ChunkLoader.ChunkX; x <= distance+ChunkLoader.ChunkX; x++ {
		for z := -distance + ChunkLoader.ChunkZ; z <= distance+ChunkLoader.ChunkZ; z++ {
			if i == maximumchunk {
				break
			}
			var xRel = x - ChunkLoader.ChunkX
			var zRel = z - ChunkLoader.ChunkZ
			if xRel*xRel+zRel*zRel <= distance*distance {
				if ChunkLoader.HasChunkInUse(x, z) {
					continue
				}

				i++
				ChunkLoader.Dimension.chunkProvider.LoadChunk(x, z, f)
			}
		}
	}
	ChunkLoader.unloadUnused(distance)
}

// unloadUnused unloads all unused chunk beyond the given distance.
func (ChunkLoader *ChunkLoader) unloadUnused(distance int32) {
	var rs = distance * distance
	ChunkLoader.mutex.Lock()
	for index, chunk := range ChunkLoader.loadedchunk {
		xDist := ChunkLoader.ChunkX - chunk.X
		zDist := ChunkLoader.ChunkZ - chunk.Z

		if xDist*xDist+zDist*zDist > rs {
			delete(ChunkLoader.loadedchunk, index)
			ChunkLoader.OnUnload(chunk)
		}
	}
	ChunkLoader.mutex.Unlock()
}

func GetChunkIndex(x, z int32) int {
	return int(((int64(x) & 0xffffffff) << 32) | (int64(z) & 0xffffffff))
}