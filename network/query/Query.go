package query

//emmmm

import (
	"github.com/juzi5201314/MineGopher/utils"
	"net"
)

var Header = []byte{0xfe, 0xfd}

const (
	HANDSHAKE  = 0x09
	STATISTICS = 0x00
)

type Query struct {
	*utils.Stream
	Address string
	Port    uint16

	Header  byte
	QueryId int32
	Token   []byte

	Statistics []byte

	IsShort bool
	Data    []byte
}

func New(buffer []byte, addr *net.UDPAddr) *Query {
	var stream = utils.NewStream()
	stream.Buffer = buffer
	return &Query{
		Stream:  stream,
		Address: addr.IP.String(),
		Port:    uint16(addr.Port),
	}
}

func (query *Query) Decode() {
	query.Offset = 2
	query.Header = query.GetByte()
	query.QueryId = query.GetInt()

	if query.Header == STATISTICS {
		query.Token = query.Get(4)
		var length = len(query.Get(-1)) + 4
		if length != 8 {
			query.IsShort = true
		}
	}
}
