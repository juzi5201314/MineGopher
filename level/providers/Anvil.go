package providers

import (
	"github.com/juzi5201314/MineGopher/level/io"
	"github.com/juzi5201314/MineGopher/nbt"
	"github.com/juzi5201314/MineGopher/utils"
	"os"
	"strconv"
	"sync"
)

type Anvil struct {
	path string
	*ChunkProvider

	mutex   sync.RWMutex
	regions map[int]*io.Region
}

// NewAnvil returns an anvil chunk provider writing and reading regions from the given path.
func NewAnvil(path string) *Anvil {
	var provider = &Anvil{path, new(), sync.RWMutex{}, make(map[int]*io.Region)}
	go provider.Process()
	return provider
}

// Process continuously processes chunk requests for chunks that were not yet loaded when requested.
func (provider *Anvil) Process() {
	for {
		var request = <-provider.requests
		if provider.IsChunkLoaded(request.x, request.z) {
			provider.completeRequest(request)
			continue
		}

		go func() {
			var regionX, regionZ = request.x >> 5, request.z >> 5
			if provider.IsRegionLoaded(regionX, regionZ) {
				provider.load(request, regionX, regionZ)
			} else {
				var path = provider.path + "r." + strconv.Itoa(int(regionX)) + "." + strconv.Itoa(int(regionZ)) + ".mca"
				var _, err = os.Stat(path)
				if err != nil {
					os.Create(path)
				}
				provider.OpenRegion(regionX, regionZ, path)
				provider.load(request, regionX, regionZ)
			}
		}()
	}
}

// load loads a chunk at the given region X and Z for the given request.
func (provider *Anvil) load(request ChunkRequest, regionX, regionZ int32) {
	var region, _ = provider.GetRegion(regionX, regionZ)
	if !region.HasChunkGenerated(request.x, request.z) {
		provider.GenerateChunk(request.x, request.z)
		provider.completeRequest(request)
		return
	}

	var compression, data = region.GetChunkData(request.x, request.z)

	var reader = nbt.NewReader(data, false, utils.BigEndian)
	var c = reader.ReadIntoCompound(int(compression))

	if c == nil {
		provider.GenerateChunk(request.x, request.z)
		provider.completeRequest(request)
		return
	}

	provider.SetChunk(request.x, request.z, io.GetAnvilChunkFromNBT(c))
	provider.completeRequest(request)
}

// IsRegionLoaded checks if a region with the given region X and Z is loaded.
func (provider *Anvil) IsRegionLoaded(regionX, regionZ int32) bool {
	provider.mutex.RLock()
	var _, ok = provider.regions[provider.GetChunkIndex(regionX, regionZ)]
	provider.mutex.RUnlock()
	return ok
}

// GetRegion returns a region with the given region X and Z, or nil if it is not loaded, and a bool indicating success.
func (provider *Anvil) GetRegion(regionX, regionZ int32) (*io.Region, bool) {
	provider.mutex.RLock()
	var region, ok = provider.regions[provider.GetChunkIndex(regionX, regionZ)]
	provider.mutex.RUnlock()
	return region, ok
}

// OpenRegion opens a region file at the given region X and Z in the given path.
// OpenRegion creates a region file if it did not yet exist.
func (provider *Anvil) OpenRegion(regionX, regionZ int32, path string) {
	var region, _ = io.OpenRegion(path)
	provider.mutex.Lock()
	provider.regions[provider.GetChunkIndex(regionX, regionZ)] = region
	provider.mutex.Unlock()
}

// Close closes the provider and saves all chunks.
func (provider *Anvil) Close(async bool) {
	var c = func() {
		for index, region := range provider.regions {
			region.Close(true)
			delete(provider.regions, index)
		}
	}
	if async {
		go c()
	} else {
		c()
	}
}

// Save saves all regions in the provider.
func (provider *Anvil) Save() {
	for _, region := range provider.regions {
		region.Save()
	}
}
