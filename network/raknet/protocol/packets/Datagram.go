package protocol

const (
	BITFLAG_VALID = 0x80
	BITFLAG_ACK   = 0x40
	BITFLAG_NAK   = 0x20

	BITFLAG_PACKET_PAIR     = 0x10
	BITFLAG_CONTINUOUS_SEND = 0x08
	BITFLAG_NEEDS_B_AND_AS  = 0x04
)

type Datagram struct {
	*Packet
	PacketPair     bool
	ContinuousSend bool
	NeedsBAndAs    bool
	SequenceNumber uint32
	packets        *[]*EncapsulatedPacket
}

func NewDatagram() *Datagram {
	datagram := &Datagram{NewPacket(0), false, false, false, 0, &[]*EncapsulatedPacket{}}
	datagram.ResetStream()
	return datagram
}

func (datagram *Datagram) GetPackets() *[]*EncapsulatedPacket {
	return datagram.packets
}

func (datagram *Datagram) Encode() {
	datagram.Buffer = []byte{}
	flags := BITFLAG_VALID
	if datagram.PacketPair {
		flags |= BITFLAG_PACKET_PAIR
	}
	if datagram.ContinuousSend {
		flags |= BITFLAG_CONTINUOUS_SEND
	}
	if datagram.NeedsBAndAs {
		flags |= BITFLAG_NEEDS_B_AND_AS
	}
	datagram.PutByte(byte(flags))
	datagram.PutLittleTriad(datagram.SequenceNumber)
	for _, packet := range *datagram.GetPackets() {
		packet.Encode()
		datagram.PutBytes(packet.Buffer)
	}
}

func (datagram *Datagram) Decode() {
	flags := datagram.GetByte()

	datagram.PacketPair = (flags & BitFlagPacketPair) != 0
	datagram.ContinuousSend = (flags & BitFlagContinuousSend) != 0
	datagram.NeedsBAndAs = (flags & BitFlagNeedsBAndAs) != 0

	datagram.SequenceNumber = datagram.GetLittleTriad()

	for !datagram.Feof() {
		epacket := NewEncapsulatedPacket()
		packet, err := epacket.GetFromBinary(datagram)
		if err != nil {
			return
		}
		datagram.packets = &(append(*datagram.packets, packet))
	}

}

func (datagram *Datagram) GetLength() int {
	length := 4
	for _, pk := range *datagram.GetPackets() {
		length += pk.GetLength()
	}
	return length
}

func (datagram *Datagram) AddPacket(packet *EncapsulatedPacket) {
	var packets = append(*datagram.packets, packet)
	datagram.packets = &packets
}
