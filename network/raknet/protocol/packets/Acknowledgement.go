package packets

import (
	"github.com/juzi5201314/MineGopher/network/raknet/protocol"
	"github.com/juzi5201314/MineGopher/utils"
	"sort"
)

type AcknowledgementPacket struct {
	*protocol.Packet
Packets []uint32
}

func (packet *AcknowledgementPacket) Encode() {
	packet.EncodeId()

	sort.Slice(packet.Packets, func(i, j int) bool {
		return packet.Packets[i] < packet.Packets[j]
	})
	packetCount := len(packet.Packets)

	if packetCount == 0 {
		packet.PutShort(0)
		return
	}

	stream := utils.NewStream()
	stream.ResetStream()

	pointer := 1
	firstPacket := packet.Packets[0]
	lastPacket := packet.Packets[0]

	intervalCount := int16(1)

	for pointer < packetCount {
		currentPacket := packet.Packets[pointer]
		difference := currentPacket - lastPacket

		if difference == 1 {
			lastPacket = currentPacket
		} else {
			if firstPacket == lastPacket {
				stream.PutByte(01)
				stream.PutLittleTriad(lastPacket)
				currentPacket = lastPacket
			} else {
				stream.PutByte(0)
				stream.PutLittleTriad(firstPacket)
				stream.PutLittleTriad(lastPacket)

				lastPacket = currentPacket
				firstPacket = lastPacket
			}
			intervalCount++
		}
		pointer++
	}

	if firstPacket == lastPacket {
		stream.PutByte(01)
		stream.PutLittleTriad(firstPacket)
	} else {
		stream.PutByte(00)
		stream.PutLittleTriad(firstPacket)
		stream.PutLittleTriad(lastPacket)
	}

	packet.PutShort(intervalCount)
	packet.PutBytes(stream.Buffer)
}

func (packet *AcknowledgementPacket) Decode() {
	packet.DecodeStep()
	packet.Packets = []uint32{}
	packetCount := packet.GetShort()
	count := 0
	for i := int16(0); i < packetCount && !packet.Feof() && count < 4096; i++ {
		if packet.GetByte() == 0 {
			start := packet.GetLittleTriad()
			end := packet.GetLittleTriad()
			if (end - start) > 512 {
				end = start + 512
			}

			for pack := start; pack < end; pack++ {
				packet.Packets = append(packet.Packets, pack)
				count++
			}

		} else {
			packet.Packets = append(packet.Packets, packet.GetLittleTriad())
			count++
		}
	}
}
