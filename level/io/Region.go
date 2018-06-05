package io

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
	"os"
	"time"
)

type CompressionType byte

const (
	HeaderSize = 8192
	SectorSize = 4096

	LengthOffset = 4

	CompressionGzip CompressionType = 1
	CompressionZlib CompressionType = 2
)

type RegionHeader struct {
	Locations  [1024]*Location
	Timestamps [1024]int32
}

type Location struct {
	Offset       int32
	SectorLength int32
}

type Region struct {
	Header RegionHeader
	File   *os.File
}

func NewRegion(path string) (*Region, error) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	return &Region{RegionHeader{}, file}, err
}

func OpenRegion(path string) (*Region, error) {
	var region, err = NewRegion(path)
	region.LoadHeader()
	return region, err
}

func GetChunkLocationIndex(x, z int32) int {
	return int((x & 31) + (z&31)*32)
}

func (r *Region) Close(save bool) {
	r.Save()
	r.File.Close()
}

func (r *Region) Save() {
	r.CleanGarbage()
	r.WriteHeader()
}

func (r *Region) LoadHeader() {
	var buff = make([]byte, 8192)
	r.File.Read(buff)

	var o int32 = 0

	for i := 0; i < 1024; i++ {
		index := i * LengthOffset
		b := bytes.NewBuffer(buff[index : index+4])

		binary.Read(b, binary.BigEndian, &o)
		offset := o >> 8

		r.Header.Locations[i] = &Location{offset << 12, o & 0xff}
	}

	var in int32 = 0
	var buffer *bytes.Buffer

	for i := 0; i < 1024; i++ {
		buffer = bytes.NewBuffer(buff[HeaderSize/2+i*LengthOffset : HeaderSize/2+i*LengthOffset+LengthOffset])
		binary.Read(buffer, binary.BigEndian, &in)
		r.Header.Timestamps[i] = in
	}
}

func (r *Region) CleanGarbage() {
	var l int32
	var lastOffset int64 = HeaderSize

	var fileBuffer = bytes.NewBuffer([]byte{})

	for i := 0; i < 1024; i++ {
		loc := r.Header.Locations[i]
		if loc.Offset == 0 {
			continue
		}

		b := make([]byte, 4)
		r.File.ReadAt(b, int64(loc.Offset))
		buffer := bytes.NewBuffer(b)
		binary.Read(buffer, binary.BigEndian, &l)

		if l <= 0 {
			loc.Offset = 0
			continue
		}

		data := make([]byte, l)
		r.File.ReadAt(data, int64(loc.Offset)+4)
		buffer.Write(data)
		paddingLength := int32(math.Ceil(float64(buffer.Len())/SectorSize) * SectorSize)
		buffer.Write(make([]byte, paddingLength-l))

		fileBuffer.Write(buffer.Bytes())
		lastOffset += int64(len(buffer.Bytes()))
	}

	io.Copy(r.File, fileBuffer)
	r.File.Truncate(lastOffset)
}

func (r *Region) WriteHeader() {
	var header = bytes.NewBuffer([]byte{})
	var offsets []int32
	for i := 0; i < 1024; i++ {
		offsetL := r.Header.Locations[i].SectorLength
		offsetI := r.Header.Locations[i].Offset >> 12 << 8
		offset := offsetI | offsetL
		offsets = append(offsets, offset)
	}
	binary.Write(header, binary.BigEndian, offsets)

	var timestamps = bytes.NewBuffer([]byte{})
	binary.Write(timestamps, binary.BigEndian, r.Header.Timestamps[:])

	var headerBytes = append(header.Bytes(), timestamps.Bytes()...)
	r.File.WriteAt(headerBytes, 0)
}

func (r *Region) GetLocation(x, z int32) *Location {
	return r.Header.Locations[GetChunkLocationIndex(x, z)]
}

func (r *Region) GetChunkData(x, z int32) (compressionType CompressionType, chunkData []byte) {
	var loc = r.GetLocation(x, z)
	if loc.Offset == 0 {
		return 0, []byte{}
	}
	var buff = make([]byte, 5)

	r.File.ReadAt(buff, int64(loc.Offset))

	var buffer = bytes.NewBuffer(buff[:4])

	var length int32
	binary.Read(buffer, binary.BigEndian, &length)
	compressionType = CompressionType(buff[4])

	chunkData = make([]byte, length)
	r.File.ReadAt(chunkData, int64(loc.Offset+5))
	return
}

func (r *Region) WriteChunkData(x, z int32, data []byte, compressionType byte) {
	var loc = r.GetLocation(x, z)
	var buff = make([]byte, LengthOffset)

	r.File.ReadAt(buff, int64(loc.Offset))

	var buffer = bytes.NewBuffer(buff[:4])

	var oldLength int32
	var newLength = int32(len(data)) + 5
	binary.Read(buffer, binary.BigEndian, &oldLength)
	oldLength += 4
	var oldLengthPadded = int32(math.Ceil(float64(oldLength)/SectorSize) * SectorSize)

	var sectorLength = int32(math.Ceil(float64(newLength) / SectorSize))
	var newLengthPadded = sectorLength

	var padding []byte
	if newLengthPadded-newLength > 0 {
		padding = make([]byte, newLengthPadded-newLength)
	}

	var offset = loc.Offset
	if newLengthPadded > oldLengthPadded {
		d, _ := r.File.Stat()
		offset = int32(d.Size())
		loc.Offset = offset
	}
	loc.SectorLength = sectorLength

	r.Header.Timestamps[GetChunkLocationIndex(x, z)] = int32(time.Now().Unix())

	buffer = bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.BigEndian, newLength)

	buffer.WriteByte(compressionType)
	buffer.Write(data)
	buffer.Write(padding)

	r.File.WriteAt(buffer.Bytes(), int64(offset))
}

func (r *Region) HasChunkGenerated(x, z int32) bool {
	return r.GetLocation(x, z).IsExistent()
}

func (location *Location) IsExistent() bool {
	return location.Offset >= HeaderSize && location.SectorLength != 0
}
