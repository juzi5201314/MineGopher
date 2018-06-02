package nbt

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"

	"github.com/juzi5201314/MineGopher/utils"
	"io/ioutil"
)

const (
	CompressionNone = 0
	CompressionGzip = 1
	CompressionZlib = 2
)

type Reader struct {
	*BinaryStream
}

func NewReader(buffer []byte, network bool, endianType utils.EndianType) *Reader {
	return &Reader{NewStream(buffer, network, endianType)}
}

func (reader *Reader) ReadUncompressedIntoCompound() *Compound {
	var tag = reader.GetTag()
	if compound, ok := tag.(*Compound); ok {
		compound.read(reader)
		return compound
	}
	return nil
}

func (reader *Reader) ReadIntoCompound(compression int) *Compound {
	if compression == CompressionNone || reader.GetOffset() != 0 {
		return reader.ReadUncompressedIntoCompound()
	}

	var data []byte
	if compression == CompressionGzip {
		var gz, _ = gzip.NewReader(bytes.NewBuffer(reader.GetBuffer()))
		data, _ = ioutil.ReadAll(gz)
		gz.Close()
	} else {
		var zl, _ = zlib.NewReader(bytes.NewBuffer(reader.GetBuffer()))
		data, _ = ioutil.ReadAll(zl)
		zl.Close()
	}

	reader.SetBuffer(data)
	return reader.ReadUncompressedIntoCompound()
}

func (reader *Reader) GetTag() INamedTag {
	if reader.Feof() {
		return NewEnd("")
	}

	var tagId = reader.GetByte()
	var tagCheck = GetTagById(tagId, "")

	if tagId == EndTag {
		return NewEnd("")
	}

	if tagCheck == nil {
		return nil
	}

	var name = reader.GetString()

	return GetTagById(tagId, name)
}
