package protocol

import (
	"github.com/juzi5201314/MineGopher/network/raknet"
	"github.com/juzi5201314/MineGopher/utils"
	"strconv"
	"strings"
)

type DataPacket interface {
	New() *DataPacket
	SetBuffer([]byte)
	GetBuffer() []byte
	GetId() int
	HasMagic() bool
	Encode()
	Decode()
}

type Packet struct {
	id int
	*utils.Stream
}

func NewPacket(id int) *Packet {
	return &Packet{id, utils.NewStream()}
}

func (packet *Packet) GetId() int {
	return packet.id
}

func (packet *Packet) HasMagic() bool {
	magic := string(raknet.MAGIC)
	return strings.Contains(string(packet.Buffer), magic)
}

func (packet *Packet) DecodeStep() {
	packet.Offset = 1
}

func (packet *Packet) EncodeId() {
	packet.Buffer = append([]byte{}, byte(packet.id))
}

func (packet *Packet) ResetBase() {
	packet.ResetStream()
}

func (packet *Packet) GetAddress() (string, uint16, byte) {
	ipv := packet.GetByte()
	var port uint16
	var addr string
	switch ipv {
	case 4:
		var parts = []byte{(-packet.GetByte() - 1) & 0xff, (-packet.GetByte() - 1) & 0xff, (-packet.GetByte() - 1) & 0xff, (-packet.GetByte() - 1) & 0xff}
		var stringArr []string
		for _, part := range parts {
			stringArr = append(stringArr, strconv.Itoa(int(part)))
		}
		addr = strings.Join(stringArr, ".")
		port = packet.GetUnsignedShort()
	case 6:
		packet.GetLittleShort()
		port = packet.GetUnsignedShort()
		packet.GetInt()
		addr = string(packet.Get(16))
		packet.GetInt()
	}
	return addr, port, ipv
}

func (packet *Packet) PutAddress(add string, port uint16, ipv byte) {
	packet.PutByte(ipv)
	switch ipv {
	case 4:
		var stringArr = strings.Split(addr, ".")
		for _, str := range stringArr {
			var digit, _ = strconv.Atoi(str)
			packet.PutByte(byte(digit))
		}
		packet.PutUnsignedShort(port)
	case 6:
		packet.PutLittleShort(23)
		packet.PutUnsignedShort(port)
		packet.PutInt(0)
		packet.PutBytes([]byte(addr))
		packet.PutInt(0)
	}
}
