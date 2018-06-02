package nbt

import (
	"bytes"
	"compress/gzip"
	"github.com/juzi5201314/MineGopher/utils"
)

type Writer struct {
	*BinaryStream
}

func NewWriter(network bool, endianType utils.EndianType) *Writer {
	return &Writer{NewStream([]byte{}, network, endianType)}
}

func (writer *Writer) WriteUncompressedCompound(compound *Compound) {
	writer.PutTag(compound)
	compound.write(writer)
}

func (writer *Writer) WriteCompressedCompound(compound *Compound) {
	writer.WriteUncompressedCompound(compound)

	var buffer = bytes.NewBuffer(writer.GetBuffer())
	var gz = gzip.NewWriter(buffer)
	gz.Write(writer.GetData())
	defer gz.Close()

	writer.PutBytes(buffer.Bytes())
}

func (writer *Writer) PutTag(tag INamedTag) {
	writer.PutByte(tag.GetType())
	if tag.GetType() != EndTag {
		writer.PutString(tag.GetName())
	}
}

func (writer *Writer) GetData() []byte {
	return writer.GetBuffer()
}
