package chunk

import (
	"errors"
	"github.com/google/uuid"
	"github.com/juzi5201314/MineGopher/nbt"
	"github.com/juzi5201314/MineGopher/utils"
	"sync"
)

type Chunk struct {
	X, Z             int32
	LightPopulated   bool
	TerrainPopulated bool
	Biomes           []byte
	HeightMap        []int16

	InhabitedTime int64
	LastUpdate    int64

	*sync.RWMutex
	viewers   map[uuid.UUID]Viewer
	entities  map[uint64]IEntity
	blockNBT  map[int]*nbt.Compound
	subChunks map[byte]*SubChunk
}

func New(x, z int32) *Chunk {
	return &Chunk{x, z,
		true,
		true,
		make([]byte, 256),
		make([]int16, 256),
		0,
		0,
		&sync.RWMutex{},
		make(map[uuid.UUID]Viewer),
		make(map[uint64]IEntity),
		make(map[int]*nbt.Compound),
		make(map[byte]*SubChunk),
	}
}

func (chunk *Chunk) GetViewers() map[uuid.UUID]Viewer {
	return chunk.viewers
}

func (chunk *Chunk) AddViewer(player Viewer) {
	chunk.Lock()
	chunk.viewers[player.GetUUID()] = player
	chunk.Unlock()
}

func (chunk *Chunk) RemoveViewer(player Viewer) {
	chunk.Lock()
	delete(chunk.viewers, player.GetUUID())
	chunk.Unlock()
}

func (chunk *Chunk) GetBiome(x, z int) byte {
	return chunk.Biomes[chunk.GetBiomeIndex(x, z)]
}

func (chunk *Chunk) SetBiome(x, z int, biome byte) {
	chunk.Biomes[chunk.GetBiomeIndex(x, z)] = biome
}

func (chunk *Chunk) AddEntity(entity IEntity) error {
	if entity.IsClosed() {
		return errors.New("cannot add closed entity to chunk")
	}
	chunk.Lock()
	chunk.entities[entity.GetEid()] = entity
	chunk.Unlock()
	return nil
}

func (chunk *Chunk) RemoveEntity(runtimeId uint64) {
	chunk.Lock()
	delete(chunk.entities, runtimeId)
	chunk.Unlock()
}

func (chunk *Chunk) GetEntities() map[uint64]IEntity {
	return chunk.entities
}

func (chunk *Chunk) SetBlockNBTAt(x, y, z int, nbt *nbt.Compound) {
	chunk.Lock()
	if nbt == nil {
		delete(chunk.blockNBT, GetBlockNBTIndex(x, y, z))
	} else {
		chunk.blockNBT[GetBlockNBTIndex(x, y, z)] = nbt
	}
	chunk.Unlock()
}

func (chunk *Chunk) RemoveBlockNBTAt(x, y, z int) {
	chunk.Lock()
	delete(chunk.blockNBT, GetBlockNBTIndex(x, y, z))
	chunk.Unlock()
}

func (chunk *Chunk) BlockNBTExistsAt(x, y, z int) bool {
	chunk.RLock()
	var _, ok = chunk.blockNBT[GetBlockNBTIndex(x, y, z)]
	chunk.RUnlock()
	return ok
}

func (chunk *Chunk) GetBlockNBTAt(x, y, z int) (*nbt.Compound, bool) {
	chunk.RLock()
	var c, ok = chunk.blockNBT[GetBlockNBTIndex(x, y, z)]
	chunk.RUnlock()
	return c, ok
}

func (chunk *Chunk) GetBiomeIndex(x, z int) int {
	return (x << 4) | z
}

func (chunk *Chunk) GetIndex(x, y, z int) int {
	return (x << 12) | (z << 8) | y
}

func (chunk *Chunk) GetHeightMapIndex(x, z int) int {
	return (z << 4) | x
}

func (chunk *Chunk) SetBlockId(x, y, z int, blockId byte) {
	chunk.GetSubChunk(byte(y>>4)).SetBlockId(x, y&15, z, blockId)
}

func (chunk *Chunk) GetBlockId(x, y, z int) byte {
	return chunk.GetSubChunk(byte(y>>4)).GetBlockId(x, y&15, z)
}

func (chunk *Chunk) SetBlockData(x, y, z int, data byte) {
	chunk.GetSubChunk(byte(y>>4)).SetBlockData(x, y&15, z, data)
}

func (chunk *Chunk) GetBlockData(x, y, z int) byte {
	return chunk.GetSubChunk(byte(y>>4)).GetBlockData(x, y&15, z)
}

func (chunk *Chunk) SetBlockLight(x, y, z int, level byte) {
	chunk.GetSubChunk(byte(y>>4)).SetBlockLight(x, y&15, z, level)
}

func (chunk *Chunk) GetBlockLight(x, y, z int) byte {
	return chunk.GetSubChunk(byte(y>>4)).GetBlockLight(x, y&15, z)
}

func (chunk *Chunk) SetSkyLight(x, y, z int, level byte) {
	chunk.GetSubChunk(byte(y>>4)).SetSkyLight(x, y&15, z, level)
}

func (chunk *Chunk) GetSkyLight(x, y, z int) byte {
	return chunk.GetSubChunk(byte(y>>4)).GetSkyLight(x, y&15, z)
}

func (chunk *Chunk) SetSubChunk(y byte, subChunk *SubChunk) {
	chunk.Lock()
	chunk.subChunks[y] = subChunk
	chunk.Unlock()
}

func (chunk *Chunk) GetSubChunk(y byte) *SubChunk {
	chunk.RLock()
	if sub, ok := chunk.subChunks[y]; ok {
		chunk.RUnlock()
		return sub
	}
	chunk.RUnlock()
	chunk.Lock()
	defer chunk.Unlock()
	chunk.subChunks[y] = NewSubChunk()
	return chunk.subChunks[y]
}

func (chunk *Chunk) SubChunkExists(y byte) bool {
	chunk.RLock()
	var _, ok = chunk.subChunks[y]
	chunk.RUnlock()
	return ok
}

func (chunk *Chunk) GetSubChunks() map[byte]*SubChunk {
	return chunk.subChunks
}

func (chunk *Chunk) SetHeightMapAt(x, z int, value int16) {
	chunk.HeightMap[chunk.GetHeightMapIndex(x, z)] = value
}

func (chunk *Chunk) GetHeightMapAt(x, z int) int16 {
	return chunk.HeightMap[chunk.GetHeightMapIndex(x, z)]
}

func (chunk *Chunk) RecalculateHeightMap() {
	for x := 0; x < 16; x++ {
		for z := 0; z < 16; z++ {
			chunk.SetHeightMapAt(x, z, chunk.GetHighestSubChunk().GetHighestBlockY(x, z)+1)
		}
	}
}

func (chunk *Chunk) GetHighestSubChunk() *SubChunk {
	chunk.RLock()
	defer chunk.RUnlock()
	for y := 15; y >= 0; y-- {
		if _, ok := chunk.subChunks[byte(y)]; !ok {
			continue
		}
		if chunk.subChunks[byte(y)].IsAllAir() {
			continue
		}
		return chunk.subChunks[byte(y)]
	}
	return nil
}

func (chunk *Chunk) GetHighestBlockId(x, z int) byte {
	return chunk.GetHighestSubChunk().GetHighestBlockId(x, z)
}

func (chunk *Chunk) GetHighestBlockData(x, z int) byte {
	return chunk.GetHighestSubChunk().GetHighestBlockData(x, z)
}

func (chunk *Chunk) GetFilledSubChunks() int {

	return len(chunk.subChunks)
}

func (chunk *Chunk) PruneEmptySubChunks() {
	chunk.Lock()
	for y := 15; y >= 0; y-- {
		if _, ok := chunk.subChunks[byte(y)]; !ok {
			continue
		}
		if !chunk.subChunks[byte(y)].IsAllAir() {
			break
		}
		delete(chunk.subChunks, byte(y))
	}
	chunk.Unlock()
}

func (chunk *Chunk) ToBinary() []byte {
	var stream = utils.NewStream()
	var subChunkCount = chunk.GetFilledSubChunks()
	stream.PutByte(byte(subChunkCount))
	for i := 0; i < subChunkCount; i++ {
		if _, ok := chunk.subChunks[byte(i)]; !ok {
			//stream.PutBytes(make([]byte, 4096+2048+1))
		} else {
			stream.PutBytes(chunk.subChunks[byte(i)].ToBinary())
		}
	}
	for i := 255; i >= 0; i-- {
		stream.PutLittleShort(chunk.HeightMap[i])
	}
	for _, biome := range chunk.Biomes {
		stream.PutByte(byte(biome))
	}
	stream.PutByte(0)
	stream.PutVarInt(0)
	return stream.GetBuffer()
}

func GetBlockNBTIndex(x, y, z int) int {
	return ((y & 256) << 8) | ((x & 15) << 4) | (z & 15)
}
