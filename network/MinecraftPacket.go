package network

import (
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	
	"github.com/juzi5201314/MineGopher/network/protocol"
	"github.com/juzi5201314/MineGopher/utils"
	"io/ioutil"
	"github.com/juzi5201314/MineGopher/api/server"
)

type MinecraftPacket struct {
	*utils.Stream
	raw      []byte
	protocol int32
	id       byte
	packets  []protocol.DataPacket
}

func NewMinecraftPacket() *MinecraftPacket {
	packet := new(MinecraftPacket)
	packet.Stream = utils.NewStream()
	return packet
}

func (packet *MinecraftPacket) GetId() byte {
	return packet.id
}

func (packet *MinecraftPacket) GetBuffer() []byte {
	return packet.Buffer
}

func (packet *MinecraftPacket) SetBuffer(buffer []byte) {
	packet.Buffer = buffer
}

func (packet *MinecraftPacket) GetProtocol() int32 {
	return packet.protocol
}

func (packet *MinecraftPacket) Decode() {
	if packet.GetByte() != 0xFE {
		return
	}
	packet.raw = packet.Buffer[packet.Offset:]
	if err := packet.decompress(); err != nil {
		server.GetServer().GetLogger().PacicError(err)
	}

	packet.ResetStream()
	packet.SetBuffer(packet.raw)

	var packetData [][]byte
	for !packet.Feof() {
		packetData = append(packetData, packet.GetLengthPrefixedBytes())
	}
	packet.fetchPackets(packetData)
}

func (packet *MinecraftPacket) fetchPackets(packetData [][]byte) {
	for _, data := range packetData {
		if len(data) == 0 {
			continue
		}
		packet.id = data[0]
		packet.protocol = packet.peekProtocol(data)

		var pk protocol.DataPacket

		pkfn, has := server.GetServer().GetNetWork().GetPacket(packet.id)
		if !has {
			server.GetServer().GetLogger().Debug(fmt.Sprintf("0x%x \n", packet.id))
			return
		}
		pk = pkfn()
		pk.SetBuffer(data)
		pk.DecodeHeader()
		pk.Decode()
		packet.packets = append(packet.packets, pk)
	}
}

func (packet *MinecraftPacket) Encode() {
	packet.ResetStream()
	packet.PutByte(0xFE)

	stream := utils.NewStream()
	packet.putPackets(stream)

	var zlibData = packet.compress(stream)
	var data = zlibData
	packet.PutBytes(data)
}

func (packet *MinecraftPacket) putPackets(stream *utils.Stream) {
	for _, pk := range packet.GetPackets() {
		pk.EncodeHeader()
		pk.Encode()
		stream.PutLengthPrefixedBytes(pk.GetBuffer())
	}
}

func (packet *MinecraftPacket) AddPacket(pk protocol.DataPacket) {
	packet.packets = append(packet.packets, pk)
}

func (packet *MinecraftPacket) compress(stream *utils.Stream) []byte {
	var buff = bytes.Buffer{}
	var writer = zlib.NewWriter(&buff)
	writer.Write(stream.Buffer)
	writer.Close()

	return buff.Bytes()
}

func (packet *MinecraftPacket) GetPackets() []protocol.DataPacket {
	return packet.packets
}

func (packet *MinecraftPacket) peekProtocol(packetData []byte) int32 {
	if packetData[0] != 0x01 {
		return 0
	}
	var protocolBytes = packetData[1:5]
	var offset = 0
	var protocol = utils.ReadInt(&protocolBytes, &offset)
	if protocol == 0 {
		offset = 0
		protocolBytes = packetData[3:7]
		protocol = utils.ReadInt(&protocolBytes, &offset)
	}
	return protocol
}

func (packet *MinecraftPacket) decompress() error {
	var reader = bytes.NewReader(packet.raw)
	zlibReader, err := zlib.NewReader(reader)

	if err != nil {
		return err
	}
	if zlibReader == nil {
		return errors.New("an error occurred when decompressing zlib")
	}
	zlibReader.Close()

	packet.raw, err = ioutil.ReadAll(zlibReader)

	return err
}
