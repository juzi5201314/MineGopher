package server

import (
	"fmt"
	"github.com/juzi5201314/MineGopher/api/server"
	"github.com/juzi5201314/MineGopher/network/raknet/protocol"
	"github.com/juzi5201314/MineGopher/network/raknet/protocol/packets"
	"net"
	"sync"
	"time"
	"github.com/juzi5201314/MineGopher/entity"
)

type Session struct {
	*net.UDPAddr
	Server        *RaknetServer
	ReceiveWindow *ReceiveWindow
	RecoveryQueue *RecoveryQueue

	MTUSize     int16
	Indexes     Indexes
	Queues      Queues
	ClientId    uint64
	CurrentPing int64
	LastUpdate  time.Time
	closed      bool
	player      *entity.Player
}

type Queues struct {
	Immediate *PriorityQueue
	High      *PriorityQueue
	Medium    *PriorityQueue
	Low       *PriorityQueue
}

type Indexes struct {
	sync.Mutex
	splits       map[int16][]*packets.EncapsulatedPacket
	splitCounts  map[int16]uint
	splitId      int16
	sendSequence uint32
	messageIndex uint32
	orderIndex   uint32
}

type RecoveryQueue struct {
	sync.Mutex
	datagrams map[uint32]*packets.Datagram
}

type TimestampedDatagram struct {
	*packets.Datagram
	Timestamp int64
}

type ReceiveWindow struct {
	DatagramHandleFunction func(datagram TimestampedDatagram)

	pendingDatagrams       chan TimestampedDatagram
	datagrams              map[uint32]TimestampedDatagram
	expectedSequenceNumber uint32
	highestSequenceNumber  uint32
}

func NewSession(addr *net.UDPAddr, mtuSize int16, server *RaknetServer) *Session {
	session := &Session{addr,
		server,
		NewReceiveWindow(),
		NewRecoveryQueue(),
		mtuSize,
		Indexes{sync.Mutex{}, make(map[int16][]*packets.EncapsulatedPacket), make(map[int16]uint), 0, 0, 0, 0},
		Queues{NewPriorityQueue(1), NewPriorityQueue(256), NewPriorityQueue(256), NewPriorityQueue(256)},
		0,
		0,
		time.Now(),
		false,
		nil,
	}
	session.player = &entity.Player{Entity: entity.New(entity.PLAYER), Session: session}
	session.ReceiveWindow.DatagramHandleFunction = func(datagram TimestampedDatagram) {
		session.LastUpdate = time.Now()
		session.SendACK(datagram.SequenceNumber)
		session.HandleDatagram(datagram)
	}

	return session
}

func (session *Session) Close() {
	session.closed = true
	session.UDPAddr = nil
	session.Server = nil
	session.ReceiveWindow = nil
	session.RecoveryQueue = nil
	session.Indexes = Indexes{}
	session.Queues = Queues{}
}

func (session *Session) IsClosed() bool {
	return session.closed
}

func (session *Session) Send(buffer []byte) (int, error) {
	return session.Server.udpServer.Write(session.UDPAddr, buffer)
}

func (session *Session) SendACK(sequenceNumber uint32) {
	ack := packets.NewACK()
	ack.Packets = []uint32{sequenceNumber}
	ack.Encode()
	session.Send(ack.Buffer)
}

func (session *Session) HandleDatagram(datagram TimestampedDatagram) {
	for _, packet := range *datagram.GetPackets() {
		if packet.HasSplit {
			session.HandleSplitEncapsulated(packet, datagram.Timestamp)
		} else {
			session.HandleEncapsulated(packet, datagram.Timestamp)
		}
	}
}

func (session *Session) HandleACK(ack *packets.ACK) {
	session.RecoveryQueue.RemoveRecovery(ack.Packets)
}

func (session *Session) HandleNACK(nack *packets.NACK) {
	for _, seq := range nack.Packets {
		if !session.RecoveryQueue.IsRecoverable(seq) {
			fmt.Println("Unrecoverable datagram:", seq, "(", nack.Packets, ")")
		}
	}
	datagrams, _ := session.RecoveryQueue.Recover(nack.Packets)
	for _, datagram := range datagrams {
		session.Server.udpServer.Write(session.UDPAddr, datagram.Buffer)
	}
}

func (session *Session) HandleEncapsulated(packet *packets.EncapsulatedPacket, timestamp int64) {
	session.LastUpdate = time.Now()
	switch packet.Buffer[0] {
	case packets.CONNECTION_REQUEST:
		session.HandleConnectionRequest(packet)
	case packets.INCOMING_CONNECTION:

	case packets.CONNECTED_PING:
		session.HandleConnectedPing(packet, timestamp)
	case packets.CONNECTED_PONG:
		session.HandleConnectedPong(packet, timestamp)
	case packets.DISCONNECT_NOTIFICATION:
		session.Close()
	default:
		session.player.HandlePacket(server.GetServer().GetNetWork().RaknetPacketToMinecraftPaket(packet.Buffer))
	}
}

func (session *Session) HandleConnectedPong(packet *packets.EncapsulatedPacket, timestamp int64) {
	pong := packets.NewConnectedPong()
	pong.Buffer = packet.Buffer
	pong.Decode()
	session.CurrentPing = (timestamp - pong.PongSendTime) / int64(time.Millisecond)
}

func (session *Session) HandleConnectedPing(packet *packets.EncapsulatedPacket, timestamp int64) {
	ping := packets.NewConnectedPing()
	ping.Buffer = packet.Buffer
	ping.Decode()

	pong := packets.NewConnectedPong()
	pong.PingSendTime = ping.PingSendTime
	pong.PongSendTime = timestamp

	session.SendPacket(pong, packets.ReliabilityUnreliable, PriorityLow)
}

func (session *Session) HandleConnectionRequest(packet *packets.EncapsulatedPacket) {
	request := packets.NewConnectionRequest()
	request.Buffer = packet.GetBuffer()
	request.Decode()

	session.ClientId = request.ClientId

	accept := packets.NewConnectionAccept()
	accept.ClientAddress = session.UDPAddr.IP.String()
	accept.ClientPort = uint16(session.UDPAddr.Port)

	accept.PingSendTime = uint64(time.Now().Unix())
	accept.PongSendTime = uint64(time.Now().Unix())

	session.SendPacket(accept, packets.ReliabilityReliableOrdered, PriorityImmediate)
}

func (session *Session) HandleSplitEncapsulated(packet *packets.EncapsulatedPacket, timestamp int64) {
	id := packet.SplitId
	session.Indexes.Lock()
	if session.Indexes.splits[id] == nil {
		session.Indexes.splits[id] = make([]*packets.EncapsulatedPacket, packet.SplitCount)
		session.Indexes.splitCounts[id] = 0
	}
	if pk := session.Indexes.splits[id][packet.SplitIndex]; pk == nil {
		session.Indexes.splitCounts[id]++
	}
	session.Indexes.splits[id][packet.SplitIndex] = packet
	if session.Indexes.splitCounts[id] == packet.SplitCount {
		newPacket := packets.NewEncapsulatedPacket()
		for _, pk := range session.Indexes.splits[id] {
			newPacket.PutBytes(pk.Buffer)
		}
		session.HandleEncapsulated(newPacket, timestamp)
		delete(session.Indexes.splits, id)
	}
	session.Indexes.Unlock()
}

func (session *Session) Tick(currentTick int64) {
	if session.closed {
		return
	}
	session.Queues.High.Flush(session)
	if currentTick%400 == 0 {
		ping := packets.NewConnectedPing()
		ping.PingSendTime = time.Now().Unix()
		session.SendPacket(ping, packets.ReliabilityUnreliable, PriorityImmediate)
	}
	if currentTick%2 == 0 {
		session.ReceiveWindow.Tick()
		session.Queues.Medium.Flush(session)
	}
	if currentTick%4 == 0 {
		session.Queues.Low.Flush(session)
	}
}

func (session *Session) SendPacket(packet protocol.ConnectedPacket, reliability byte, priority byte) {
	packet.Encode()
	encapsulated := packets.NewEncapsulatedPacket()
	encapsulated.Reliability = reliability
	encapsulated.Buffer = packet.GetBuffer()
	session.Queues.AddEncapsulated(encapsulated, priority, session)
}

func (queues Queues) AddEncapsulated(packet *packets.EncapsulatedPacket, priority byte, session *Session) {
	if session.IsClosed() {
		return
	}
	var queue *PriorityQueue
	switch priority {
	case PriorityImmediate:
		queue = queues.Immediate
	case PriorityHigh:
		queue = queues.High
	case PriorityMedium:
		queue = queues.Medium
	case PriorityLow:
		queue = queues.Low
	}
	queue.AddEncapsulated(packet, session)
	if priority == PriorityImmediate {
		queue.Flush(session)
	}
}

func NewReceiveWindow() *ReceiveWindow {
	return &ReceiveWindow{func(datagram TimestampedDatagram) {}, make(chan TimestampedDatagram, 128), make(map[uint32]TimestampedDatagram), 0, 0}
}

func (window *ReceiveWindow) AddDatagram(datagram *packets.Datagram) {
	if datagram.SequenceNumber < window.expectedSequenceNumber {
		return
	}
	if datagram.SequenceNumber > window.highestSequenceNumber {
		window.highestSequenceNumber = datagram.SequenceNumber
	}
	window.pendingDatagrams <- TimestampedDatagram{datagram, time.Now().Unix()}
}

func (window *ReceiveWindow) Tick() {
	for len(window.pendingDatagrams) > 0 {
		datagram := <-window.pendingDatagrams
		window.datagrams[datagram.SequenceNumber] = datagram
	}
	for i := window.expectedSequenceNumber; ; i++ {
		if datagram, ok := window.datagrams[i]; ok {
			go window.DatagramHandleFunction(datagram)
			window.expectedSequenceNumber++
			delete(window.datagrams, i)
		} else {
			return
		}
	}
}

func NewRecoveryQueue() *RecoveryQueue {
	return &RecoveryQueue{sync.Mutex{}, make(map[uint32]*packets.Datagram)}
}

func (queue *RecoveryQueue) AddRecovery(datagram *packets.Datagram) {
	queue.Lock()
	queue.datagrams[datagram.SequenceNumber] = datagram
	queue.Unlock()
}

func (queue *RecoveryQueue) IsRecoverable(sequenceNumber uint32) bool {
	queue.Lock()
	_, ok := queue.datagrams[sequenceNumber]
	queue.Unlock()
	return ok
}

func (queue *RecoveryQueue) RemoveRecovery(sequenceNumbers []uint32) {
	queue.Lock()
	for _, sequenceNumber := range sequenceNumbers {
		delete(queue.datagrams, sequenceNumber)
	}
	queue.Unlock()
}

func (queue *RecoveryQueue) Recover(sequenceNumbers []uint32) ([]*packets.Datagram, []uint32) {
	var datagrams []*packets.Datagram
	var recoveredSequenceNumbers []uint32
	queue.Lock()
	for _, sequenceNumber := range sequenceNumbers {
		if datagram, ok := queue.datagrams[sequenceNumber]; ok {
			datagrams = append(datagrams, datagram)
			recoveredSequenceNumbers = append(recoveredSequenceNumbers, sequenceNumber)
		}
	}
	queue.Unlock()
	return datagrams, sequenceNumbers
}
