package level

import (
	"github.com/golang/geo/r3"
	"math"
	"sync"
	"github.com/juzi5201314/MineGopher/level/providers"
	"github.com/google/uuid"
	"github.com/juzi5201314/MineGopher/level/generation"

	"github.com/juzi5201314/MineGopher/level/chunk"
)

type Dimension struct {
	level *Level

	chunkProvider providers.Provider

	mutex    sync.RWMutex
	entities map[uint64]chunk.IEntity
	viewers  map[uuid.UUID]chunk.Viewer
}


var EntityRuntimeId uint64

// NewDimension returns a new dimension with the given name, levelName, dimension ID and server path.
// Dimension data will be written in the `serverPath/worlds/levelName/name` path.
func NewDimension(level *Level) *Dimension {
	var dimension = &Dimension{level, nil, sync.RWMutex{}, make(map[uint64]chunk.IEntity), make(map[uuid.UUID]chunk.Viewer)}
	return dimension
}

// GetName returns the name of the dimension.
func (dimension *Dimension) GetName() string {
	return dimension.level.name
}

// GetLevel returns the level of the dimension.
func (dimension *Dimension) GetLevel() *Level {
	return dimension.level
}

// Close closes the dimension and saves it.
// If async is true, closes the dimension asynchronously.
func (dimension *Dimension) Close(async bool) {
	dimension.chunkProvider.Close(async)
}

// Save saves the dimension.
func (dimension *Dimension) Save() {
	dimension.chunkProvider.Save()
}

// GetEntities returns all loaded entities in this dimension in a runtime ID => entity map.
func (dimension *Dimension) GetEntities() map[uint64]chunk.IEntity {
	return dimension.entities
}

// GetViewers returns all entities considered as viewers in the dimension.
func (dimension *Dimension) GetViewers() map[uuid.UUID]chunk.Viewer {
	return dimension.viewers
}

// AddViewer adds a viewer to the dimension.
func (dimension *Dimension) AddViewer(viewer chunk.Viewer, position r3.Vector) {
	x, z := int32(math.Floor(position.X))>>4, int32(math.Floor(position.Z))>>4
	dimension.LoadChunk(x, z, func(chunk *chunk.Chunk) {
		dimension.mutex.Lock()
		dimension.viewers[viewer.GetUUID()] = viewer
		dimension.mutex.Unlock()
		chunk.AddViewer(viewer)
	})
}

// RemoveViewer removes a viewer from the dimension.
func (dimension *Dimension) RemoveViewer(uuid uuid.UUID) {
	dimension.mutex.Lock()
	delete(dimension.viewers, uuid)
	dimension.mutex.Unlock()
}

// GetViewer returns a viewer of a dimension by its UUID.
// A bool gets returned indicating whether the viewer was found or not.
func (dimension *Dimension) GetViewer(uuid uuid.UUID) (chunk.Viewer, bool) {
	dimension.mutex.RLock()
	viewer, ok := dimension.viewers[uuid]
	dimension.mutex.RUnlock()
	return viewer, ok
}

// AddEntity adds a new entity at the given position in the dimension.
func (dimension *Dimension) AddEntity(entity chunk.IEntity, position r3.Vector) {
	var x, z = int32(math.Floor(position.X)) >> 4, int32(math.Floor(position.Z)) >> 4
	dimension.LoadChunk(x, z, func(chunk *chunk.Chunk) {
		entity.SetPosition(position)
		entity.SpawnToAll()

		chunk.AddEntity(entity)
		dimension.mutex.Lock()
		dimension.entities[EntityRuntimeId] = entity
		dimension.mutex.Unlock()
	})
}

// RemoveEntity removes an entity in the dimension with the given runtime ID.
// The removed entity also gets closed if not yet done.
func (dimension *Dimension) RemoveEntity(runtimeId uint64) {
	dimension.mutex.Lock()
	if entity, ok := dimension.entities[runtimeId]; ok {
		if !entity.IsClosed() {
			entity.Close()
		}
		var x, z = int32(math.Floor(entity.GetPosition().X)), int32(math.Floor(entity.GetPosition().Z))
		if chunk, ok := dimension.GetChunk(x, z); ok {
			chunk.RemoveEntity(runtimeId)
		}
		delete(dimension.entities, runtimeId)
	}
	dimension.mutex.Unlock()
}

// GetEntity returns an entity in the dimension by its runtime ID.
// Returns UnavailableEntity error if no entity with that runtime ID was available in the dimension.
func (dimension *Dimension) GetEntity(runtimeId uint64) (chunk.IEntity, bool) {
	dimension.mutex.RLock()
	defer dimension.mutex.RUnlock()
	entity, ok := dimension.entities[runtimeId]
		return entity, ok
}

// HasEntity checks if the dimension has an entity available with the given runtime ID.
func (dimension *Dimension) HasEntity(runtimeId uint64) bool {
	dimension.mutex.RLock()
	var _, ok = dimension.entities[runtimeId]
	dimension.mutex.RUnlock()
	return ok
}

// IsChunkLoaded checks if a chunk at the given chunk X and Z is loaded.
func (dimension *Dimension) IsChunkLoaded(x, z int32) bool {
	return dimension.chunkProvider.IsChunkLoaded(x, z)
}

// UnloadChunk unloads a chunk at the given chunk X and Z.
func (dimension *Dimension) UnloadChunk(x, z int32) {
	dimension.chunkProvider.UnloadChunk(x, z)
}

// LoadChunk submits a request with the given chunk X and Z to get loaded.
// The function given gets run as soon as the chunk gets loaded.
func (dimension *Dimension) LoadChunk(x, z int32, function func(chunk *chunk.Chunk)) {
	dimension.chunkProvider.LoadChunk(x, z, function)
}

// SetChunk sets a new chunk at the given chunk X and Z.
func (dimension *Dimension) SetChunk(x, z int32, chunk *chunk.Chunk) {
	dimension.chunkProvider.SetChunk(x, z, chunk)
}

// GetChunk returns a chunk in the dimension at the given chunk X and Z.
func (dimension *Dimension) GetChunk(x, z int32) (*chunk.Chunk, bool) {
	if dimension == nil {
		println(414)
	}
	return dimension.chunkProvider.GetChunk(x, z)
}

// SetGenerator sets the generator of the dimension.
func (dimension *Dimension) SetGenerator(generator generation.Generator) {
	dimension.chunkProvider.SetGenerator(generator)
}

// GetGenerator returns the generator of the dimension.
func (dimension *Dimension) GetGenerator() generation.Generator {
	return dimension.chunkProvider.GetGenerator()
}

// GetChunkProvider returns the chunk provider of the dimension.
func (dimension *Dimension) GetChunkProvider() providers.Provider {
	return dimension.chunkProvider
}

// SetChunkProvider sets the chunk provider of the dimension.
func (dimension *Dimension) SetChunkProvider(provider providers.Provider) {
	dimension.chunkProvider = provider
}

/*
func (dimension *Dimension) GetBlockAt(vector r3.Vector) (blocks.Block, error) {
	var x, y, z = int(math.Floor(vector.X)), int(math.Floor(vector.Y)), int(math.Floor(vector.Z))
	var chunk, ok = dimension.GetChunk(int32(x>>4), int32(z>>4))
	if !ok {
		return nil, UnloadedChunk
	}
	var id, meta = chunk.GetBlockId(x&15, y, z&15), chunk.GetBlockData(x&15, y, z&15)
	var block, err = dimension.blockManager.Get(id, meta)
	if err != nil {
		if nbt, ok := chunk.GetBlockNBTAt(x&15, y, z&15); ok {
			block.SetNBT(nbt)
		}
	}
	return block, err
}

func (dimension *Dimension) SetBlockAt(vector r3.Vector, block blocks.Block) {
	var x, y, z = int(math.Floor(vector.X)), int(math.Floor(vector.Y)), int(math.Floor(vector.Z))
	dimension.LoadChunk(int32(x>>4), int32(z>>4), func(chunk *chunk.Chunk) {
		chunk.SetBlockId(x&15, y, z&15, block.GetId())
		chunk.SetBlockData(x&15, y, z&15, block.GetData())
		chunk.SetBlockNBTAt(x&15, y, z&15, block.GetNBT())
	})
}
*/

// Tick ticks the entire dimension, such as entities.
func (dimension *Dimension) Tick() {
	for runtimeId, entity := range dimension.entities {
		if entity.IsClosed() {
			dimension.RemoveEntity(runtimeId)
		} else {
			entity.Tick()
		}
	}
}