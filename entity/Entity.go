package entity

/*
import (
	"errors"
	"github.com/golang/geo/r3"
	"github.com/google/uuid"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/protocol"
	"github.com/irmine/gonbt"
	"github.com/irmine/worlds"
	"github.com/irmine/worlds/chunks"
	"math"
	"sync"
)

type Viewer interface {
	chunks.Viewer
	SendAddEntity(protocol.AddEntityEntry)
	SendAddPlayer(uuid.UUID, int32, protocol.AddPlayerEntry)
	SendPacket(packet packets.IPacket)
	SendRemoveEntity(int64)
}

type Entity struct {
	id           uint
	attributeMap AttributeMap

	Position r3.Vector
	Rotation Rotation
	Motion   r3.Vector
	OnGround bool

	Dimension *worlds.Dimension

	NameTag string

	runtimeId uint64
	closed    bool

	nbt *gonbt.Compound

	mutex      sync.RWMutex
	EntityData map[uint32][]interface{}
	SpawnedTo  map[uuid.UUID]Viewer
}

var UnloadedChunkMove = errors.New("tried to move entity in unloaded chunk")

func New(id uint) *Entity {
	ent := Entity{
		id,
		NewAttributeMap(),
		r3.Vector{},
		Rotation{},
		r3.Vector{},
		false,
		nil,
		"",
		0,
		true,
		gonbt.NewCompound("", make(map[string]gonbt.INamedTag)),
		sync.RWMutex{},
		make(map[uint32][]interface{}),
		make(map[uuid.UUID]Viewer),
	}
	return &ent
}

func (entity *Entity) GetNameTag() string {
	return entity.NameTag
}

func (entity *Entity) SetNameTag(nameTag string) {
	entity.NameTag = nameTag
}

func (entity *Entity) GetAttributeMap() AttributeMap {
	return entity.attributeMap
}

func (entity *Entity) SetAttributeMap(attMap AttributeMap) {
	entity.attributeMap = attMap
}

func (entity *Entity) GetEntityData() map[uint32][]interface{} {
	return entity.EntityData
}

func (entity *Entity) GetPosition() r3.Vector {
	return entity.Position
}

func (entity *Entity) SetPosition(v r3.Vector) error {
	var newChunkX = int32(math.Floor(float64(v.X))) >> 4
	var newChunkZ = int32(math.Floor(float64(v.Z))) >> 4

	var oldChunk = entity.GetChunk()
	var newChunk, ok = entity.Dimension.GetChunk(newChunkX, newChunkZ)
	if !ok {
		return UnloadedChunkMove
	}

	entity.Position = v

	if oldChunk != newChunk {
		newChunk.AddEntity(entity)
		entity.SpawnToAll()
		oldChunk.RemoveEntity(entity.runtimeId)
	}
	return nil
}

func (entity *Entity) IsOnGround() bool {
	return entity.OnGround
}

func (entity *Entity) GetChunk() *chunks.Chunk {
	var x = int32(math.Floor(float64(entity.Position.X))) >> 4
	var z = int32(math.Floor(float64(entity.Position.Z))) >> 4
	var chunk, _ = entity.Dimension.GetChunk(x, z)
	return chunk
}

func (entity *Entity) GetViewers() map[uuid.UUID]Viewer {
	return entity.SpawnedTo
}

func (entity *Entity) AddViewer(viewer Viewer) {
	entity.mutex.Lock()
	entity.SpawnedTo[viewer.GetUUID()] = viewer
	entity.mutex.Unlock()
}

func (entity *Entity) RemoveViewer(viewer Viewer) {
	entity.mutex.Lock()
	delete(entity.SpawnedTo, viewer.GetUUID())
	entity.mutex.Unlock()
}

func (entity *Entity) GetDimension() *worlds.Dimension {
	return entity.Dimension
}

func (entity *Entity) SetDimension(v interface {
	GetChunk(int32, int32) (*chunks.Chunk, bool)
}) {
	entity.Dimension = v.(*worlds.Dimension)
}

func (entity *Entity) GetRotation() Rotation {
	return entity.Rotation
}

func (entity *Entity) SetRotation(v Rotation) {
	entity.Rotation = v
}

func (entity *Entity) GetMotion() r3.Vector {
	return entity.Motion
}

func (entity *Entity) SetMotion(v r3.Vector) {
	entity.Motion = v
}

func (entity *Entity) GetRuntimeId() uint64 {
	return entity.runtimeId
}

func (entity *Entity) SetRuntimeId(id uint64) {
	entity.runtimeId = id
}

func (entity *Entity) GetUniqueId() int64 {
	return int64(entity.runtimeId)
}

func (entity *Entity) GetEntityId() uint {
	return entity.id
}

func (entity *Entity) IsClosed() bool {
	return entity.closed
}

func (entity *Entity) Close() {
	entity.closed = true
	entity.DespawnFromAll()

	entity.Dimension = nil
	entity.SpawnedTo = nil
}

func (entity *Entity) GetHealth() float32 {
	return entity.attributeMap.GetAttribute(AttributeHealth).Value
}

func (entity *Entity) SetHealth(health float32) {
	entity.attributeMap.GetAttribute(AttributeHealth).Value = health
}

func (entity *Entity) Kill() {
	entity.SetHealth(0)
}

func (entity *Entity) SpawnTo(viewer Viewer) {
	if entity.IsClosed() {
		return
	}
	entity.AddViewer(viewer)
	viewer.SendAddEntity(entity)
}

func (entity *Entity) DespawnFrom(viewer Viewer) {
	entity.RemoveViewer(viewer)
	viewer.SendRemoveEntity(entity.GetUniqueId())
}

func (entity *Entity) DespawnFromAll() {
	for _, viewer := range entity.SpawnedTo {
		entity.DespawnFrom(viewer)
	}
}

func (entity *Entity) SpawnToAll() {
	for _, v := range entity.GetChunk().GetViewers() {
		var (
			viewer Viewer
			ok     bool
		)
		if viewer, ok = v.(Viewer); !ok {
			continue
		}
		if _, ok := entity.SpawnedTo[viewer.GetUUID()]; !ok {
			entity.SpawnTo(viewer)
		}
	}
}

func (entity *Entity) GetNBT() *gonbt.Compound {
	return entity.nbt
}

func (entity *Entity) SetNBT(nbt *gonbt.Compound) {
	entity.nbt = nbt
}

func (entity *Entity) Tick() {

}
*/
