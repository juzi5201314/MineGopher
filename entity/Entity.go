package entity

import (
	"github.com/golang/geo/r3"
	"github.com/google/uuid"
	"github.com/juzi5201314/MineGopher/level"
	"github.com/juzi5201314/MineGopher/nbt"
	"math"
	"sync"

	"github.com/juzi5201314/MineGopher/entity/data"
	"github.com/juzi5201314/MineGopher/level/chunk"
	"github.com/juzi5201314/MineGopher/network/protocol"
)

var EntityCount uint64 = 0

type Viewer interface {
	chunk.Viewer
	SendPacket(packet protocol.DataPacket)
}

type Entity struct {
	closed       bool
	id           uint32
	eid          uint64
	attributeMap data.AttributeMap
	Position     r3.Vector
	Rotation     data.Rotation
	Motion       r3.Vector
	nameTag      string
	Dimension    *level.Dimension
	nbt          *nbt.Compound
	mutex        sync.RWMutex
	EntityData   map[uint32][]interface{}
	viewers      map[uuid.UUID]Viewer
}

func New(id uint32) *Entity {
	defer func() {
		EntityCount++
	}()
	return &Entity{
		false,
		id,
		EntityCount,
		data.NewAttributeMap(),
		r3.Vector{},
		data.Rotation{},
		r3.Vector{},
		"",
		nil,
		nbt.NewCompound("", make(map[string]nbt.INamedTag)),
		sync.RWMutex{},
		make(map[uint32][]interface{}),
		make(map[uuid.UUID]Viewer),
	}
}

func (entity Entity) Tick() {

}

//这sb函数有毒
func (entity Entity) SetDimension(v interface {
	GetChunk(int32, int32) (*chunk.Chunk, bool)
}) {
	entity.Dimension = v.(*level.Dimension)
}

func (entity Entity) GetDimension() *level.Dimension {
	return entity.Dimension
}

func (entity Entity) Close() {
	entity.closed = true
}

func (entity Entity) IsClosed() bool {
	return entity.closed
}

func (entity Entity) GetId() uint32 {
	return entity.id
}

func (entity Entity) GetEid() uint64 {
	return entity.eid
}

func (entity *Entity) GetNameTag() string {
	return entity.nameTag
}

func (entity *Entity) SetNameTag(nameTag string) {
	entity.nameTag = nameTag
}

func (entity *Entity) GetAttributeMap() data.AttributeMap {
	return entity.attributeMap
}

func (entity *Entity) SetAttributeMap(attMap data.AttributeMap) {
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
	entity.GetDimension().GetName()
	var oldChunk = entity.GetChunk()
	var newChunk, ok = entity.Dimension.GetChunk(newChunkX, newChunkZ)
	if !ok {

	}

	entity.Position = v

	if oldChunk != newChunk {
		newChunk.AddEntity(entity)
		entity.SpawnToAll()
		oldChunk.RemoveEntity(entity.eid)
	}
	return nil
}

func (entity *Entity) GetChunk() *chunk.Chunk {
	var x = int32(math.Floor(float64(entity.Position.X))) >> 4
	var z = int32(math.Floor(float64(entity.Position.Z))) >> 4
	var chunk, _ = entity.Dimension.GetChunk(x, z)
	return chunk
}

func (entity *Entity) GetHealth() float32 {
	return entity.attributeMap.GetAttribute(data.AttributeHealth).Value
}

func (entity *Entity) SetHealth(health float32) {
	entity.attributeMap.GetAttribute(data.AttributeHealth).Value = health
}

func (entity *Entity) Kill() {
	entity.SetHealth(0)
}

func (entity *Entity) SpawnTo(viewer Viewer) {
	if entity.IsClosed() {
		return
	}
	entity.AddViewer(viewer)
	pk := &protocol.AddEntityPacket{Packet: protocol.NewPacket(protocol.GetPacketId(protocol.ADD_ENTITY_PACKET))}
	pk.UniqueId = entity.GetUniqueId()
	pk.RuntimeId = entity.GetEid()
	pk.EntityType = entity.GetId()
	pk.Position = entity.GetPosition()
	pk.Motion = entity.GetMotion()
	pk.Rotation = entity.GetRotation()
	pk.Attributes = entity.GetAttributeMap()
	pk.EntityData = entity.GetEntityData()
	viewer.SendPacket(pk)
}

func (entity *Entity) DespawnTo(viewer Viewer) {
	entity.RemoveViewer(viewer)
	pk := &protocol.RemoveEntityPacket{protocol.NewPacket(protocol.GetPacketId(protocol.REMOVE_ENTITY_PACKET)), entity.GetUniqueId()}
	viewer.SendPacket(pk)
}

func (entity *Entity) DespawnToAll() {
	for _, viewer := range entity.viewers {
		entity.DespawnTo(viewer)
	}
}

func (entity *Entity) SpawnToAll() {
	for _, v := range entity.GetChunk().GetViewers() {
		if _, in := entity.viewers[v.GetUUID()]; !in {
			entity.SpawnTo(v.(Viewer))
		}
	}
}

func (entity *Entity) GetNBT() *nbt.Compound {
	return entity.nbt
}

func (entity *Entity) SetNBT(nbt *nbt.Compound) {
	entity.nbt = nbt
}

func (entity *Entity) GetViewers() map[uuid.UUID]Viewer {
	return entity.viewers
}

func (entity *Entity) AddViewer(viewer Viewer) {
	entity.mutex.Lock()
	entity.viewers[viewer.GetUUID()] = viewer
	entity.mutex.Unlock()
}

func (entity *Entity) RemoveViewer(viewer Viewer) {
	entity.mutex.Lock()
	delete(entity.viewers, viewer.GetUUID())
	entity.mutex.Unlock()
}

func (entity *Entity) GetUniqueId() int64 {
	return int64(entity.eid)
}

func (entity *Entity) GetRotation() data.Rotation {
	return entity.Rotation
}

func (entity *Entity) SetRotation(v data.Rotation) {
	entity.Rotation = v
}

func (entity *Entity) GetMotion() r3.Vector {
	return entity.Motion
}

func (entity *Entity) SetMotion(v r3.Vector) {
	entity.Motion = v
}

func (entity *Entity) UpdateAttributes() {
	pk := &protocol.UpdateAttributesPacket{protocol.NewPacket(protocol.GetPacketId(protocol.UPDATE_ATTRIBUTES_PACKET)), entity.eid, entity.attributeMap}
	for _, p := range entity.viewers {
		p.SendPacket(pk)
	}
}

/*
import (
	"errors"
	"github.com/golang/geo/r3"
	"github.com/google/uuid"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/protocol"
	"github.com/irmine/nbt"
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

	attributeMap AttributeMap

	Position r3.Vector
	Rotation Rotation

	runtimeId uint64
	closed    bool

	nbt *nbt.Compound

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
		nbt.NewCompound("", make(map[string]nbt.INamedTag)),
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

func (entity *Entity) GetNBT() *nbt.Compound {
	return entity.nbt
}

func (entity *Entity) SetNBT(nbt *nbt.Compound) {
	entity.nbt = nbt
}

func (entity *Entity) Tick() {

}
*/
