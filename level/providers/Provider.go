package providers

import (
	"github.com/juzi5201314/MineGopher/level/chunk"
	"github.com/juzi5201314/MineGopher/level/generation"
	"sync"
)

type Provider interface {
	Save()
	Close(bool)
	LoadChunk(int32, int32, func(*chunk.Chunk))
	IsChunkLoaded(int32, int32) bool
	UnloadChunk(int32, int32)
	SetChunk(int32, int32, *chunk.Chunk)
	GetChunk(int32, int32) (*chunk.Chunk, bool)
	SetGenerator(generation.Generator)
	GetGenerator() generation.Generator
	GenerateChunk(int32, int32)
}

type ChunkProvider struct {
	generator generation.Generator
	requests  chan ChunkRequest

	mutex sync.RWMutex
	chunk map[int]*chunk.Chunk
}

type ChunkRequest struct {
	function func(*chunk.Chunk)
	x        int32
	z        int32
}

func new() *ChunkProvider {
	return &ChunkProvider{requests: make(chan ChunkRequest, 4096), chunk: make(map[int]*chunk.Chunk)}
}

func (provider *ChunkProvider) LoadChunk(x, z int32, function func(*chunk.Chunk)) {
	if chunk, ok := provider.GetChunk(x, z); ok {
		function(chunk)
		return
	}
	provider.requests <- ChunkRequest{function, x, z}
}

func (provider *ChunkProvider) IsChunkLoaded(x, z int32) bool {
	provider.mutex.RLock()
	var _, ok = provider.chunk[provider.GetChunkIndex(x, z)]
	provider.mutex.RUnlock()
	return ok
}

func (provider *ChunkProvider) UnloadChunk(x, z int32) {
	provider.mutex.Lock()
	delete(provider.chunk, provider.GetChunkIndex(x, z))
	provider.mutex.Unlock()
}

func (provider *ChunkProvider) SetChunk(x, z int32, chunk *chunk.Chunk) {
	provider.mutex.Lock()
	provider.chunk[provider.GetChunkIndex(x, z)] = chunk
	provider.mutex.Unlock()
}

func (provider *ChunkProvider) GetChunk(x, z int32) (*chunk.Chunk, bool) {
	provider.mutex.RLock()
	var chunk, ok = provider.chunk[provider.GetChunkIndex(x, z)]
	provider.mutex.RUnlock()
	return chunk, ok
}

func (provider *ChunkProvider) SetGenerator(generator generation.Generator) {
	provider.generator = generator
}

func (provider *ChunkProvider) GetGenerator() generation.Generator {
	return provider.generator
}

func (provider *ChunkProvider) completeRequest(request ChunkRequest) {
	var chunk, ok = provider.GetChunk(request.x, request.z)
	if ok {
		request.function(chunk)
	}
}

func (provider *ChunkProvider) GenerateChunk(x, z int32) {
	var chunk = provider.generator.GenerateNewChunk(x, z)
	provider.SetChunk(x, z, chunk)
}

func (provider *ChunkProvider) GetChunkIndex(x, z int32) int {
	return int(((int64(x) & 0xffffffff) << 32) | (int64(z) & 0xffffffff))
}
