package server

import (
	"github.com/juzi5201314/MineGopher/network/raknet/protocol/packets"
	"math"
)

const (
	PriorityImmediate byte = iota
	PriorityHigh
	PriorityMedium
	PriorityLow
)

type PriorityQueue chan *packets.EncapsulatedPacket

func NewPriorityQueue(bufferingSize int) *PriorityQueue {
	queue := PriorityQueue(make(chan *packets.EncapsulatedPacket, bufferingSize))
	return &queue
}

func (queue *PriorityQueue) AddEncapsulated(packet *packets.EncapsulatedPacket, session *Session) {
	for _, encapsulated := range queue.Split(packet, session) {
		*queue <- encapsulated
	}
}

func (queue *PriorityQueue) Flush(session *Session) {
	if len(*queue) == 0 {
		return
	}
	ind := 0
	datagram := packets.NewDatagram()
	datagram.NeedsBAndAs = true
	datagrams := map[int]*packets.Datagram{0: datagram}
	datagram.SequenceNumber = session.Indexes.sendSequence
	session.Indexes.sendSequence++

	i := 0
	for len(*queue) > 0 && i < 16 {
		i++
		packet := <-*queue
		if datagrams[ind].GetLength()+packet.GetLength() > int(session.MTUSize-38) {
			ind++
			datagrams[ind] = packets.NewDatagram()
			datagrams[ind].NeedsBAndAs = true
			datagrams[ind].SequenceNumber = session.Indexes.sendSequence
			session.Indexes.sendSequence++
		}
		datagrams[ind].AddPacket(packet)
	}

	l := len(datagrams)
	for j := 0; j < l; j++ {
		datagram := datagrams[j]
		datagram.Encode()
		session.RecoveryQueue.AddRecovery(datagram)
		session.Server.udpServer.Write(session.UDPAddr, datagram.Buffer)
	}
}

func (queue *PriorityQueue) Split(packet *packets.EncapsulatedPacket, session *Session) []*packets.EncapsulatedPacket {
	mtuSize := int(session.MTUSize - 60)
	var pks []*packets.EncapsulatedPacket
	if packet.IsSequenced() {
		packet.OrderIndex = session.Indexes.orderIndex
		session.Indexes.orderIndex++
	}

	if packet.GetLength() > mtuSize {
		buffer := packet.GetBuffer()
		splitSize := mtuSize
		var b uint
		for i := 0; i < len(buffer)+splitSize; i += splitSize {
			if i+splitSize >= len(buffer) {
				splitSize = len(buffer) - i
				if splitSize == 0 {
					break
				}
			}
			split := buffer[i : i+splitSize]
			encapsulated := packets.NewEncapsulatedPacket()
			encapsulated.HasSplit = true
			encapsulated.SplitId = session.Indexes.splitId
			encapsulated.SplitIndex = b
			encapsulated.SplitCount = uint(math.Ceil(float64(len(buffer)) / float64(splitSize)))
			encapsulated.Reliability = packet.Reliability
			b++
			encapsulated.Buffer = split
			encapsulated.OrderIndex = packet.OrderIndex
			if packet.IsReliable() {
				encapsulated.MessageIndex = session.Indexes.messageIndex
				session.Indexes.messageIndex++
			}
			pks = append(pks, encapsulated)
		}
		session.Indexes.splitId++
	} else {
		if packet.IsReliable() {
			packet.MessageIndex = session.Indexes.messageIndex
			session.Indexes.messageIndex++
		}
		pks = append(pks, packet)
	}
	return pks
}
