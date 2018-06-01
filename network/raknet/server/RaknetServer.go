package server

import (
	"fmt"
	"github.com/juzi5201314/MineGopher/network/raknet/protocol"
	"github.com/juzi5201314/MineGopher/network/raknet/protocol/packets"
	"net"
	"sync"
	"time"
	"math/rand"
	"bytes"
	"github.com/juzi5201314/MineGopher/network/query"
	"github.com/juzi5201314/MineGopher/api"
)

const (
	//如果客户端超过5s没有"报平安"，服务器就会认为他game over了
	TIMEOUT = 5
)

type RaknetServer struct {
	udpServer   *UDPServer
	sessions    Sessions
	Timeout     time.Duration
	ipBlockList map[string]*net.UDPAddr
	running     bool
	CurrentTick int64
	id int64
	*sync.RWMutex
}

type Sessions map[string]*Session

func New(ip string, port int) *RaknetServer {
	return &RaknetServer{
		udpServer:   NewUDPServer(ip, port),
		sessions:    Sessions{},
		Timeout:     time.Second * TIMEOUT,
		ipBlockList: map[string]*net.UDPAddr{},
		id: rand.NewSource(time.Now().Unix()).Int63(),
		RWMutex:     &sync.RWMutex{},
	}
}

func (server *RaknetServer) Start() {
	server.running = true
	server.udpServer.Start()
	go func() {
		for {
			server.processPacket()
		}
	}()
	go server.sessionTick()
}

func (server *RaknetServer) Close() {
	server.running = false
}

func (server *RaknetServer) BlockIp(addr *net.UDPAddr, t time.Duration) {
	server.Lock()
	server.ipBlockList[fmt.Sprint(addr.IP)] = addr
	server.Unlock()
	time.AfterFunc(t, func() {
		server.UnBlockIp(addr)
	})
}

func (server *RaknetServer) UnBlockIp(addr *net.UDPAddr) {
	server.Lock()
	delete(server.ipBlockList, fmt.Sprint(addr.IP))
	server.Unlock()
}

func (server *RaknetServer) IsBlockedIp(addr *net.UDPAddr) bool {
	server.RLock()
	_, is := server.ipBlockList[fmt.Sprint(addr.IP)]
	server.RUnlock()
	return is
}

func (server *RaknetServer) sessionTick() {
	ticker := time.NewTicker(time.Duration(time.Second / 80))
	for range ticker.C {
		if !server.running {
			return
		}
		for index, session := range server.sessions {
			server.updateSession(session, index)
		}
		server.CurrentTick++
	}
}

func (server *RaknetServer) updateSession(session *Session, index string) {
	session.Tick(server.CurrentTick)
	if time.Now().Sub(session.LastUpdate) > server.Timeout {
		//game over - timeout
		session.Close()
		delete(server.sessions, index)
	}
}

func (server *RaknetServer) processPacket() {
	buffer := make([]byte, 2048)
	n, addr, err := server.udpServer.Read(buffer)
	if err != nil {
		//蜜汁错误，吓到服务器了，先封他号3秒压压惊
		server.BlockIp(addr, time.Second*3)
		return
	}
	if server.IsBlockedIp(addr) {
		return
	}
	buffer = buffer[:n]
	pid := buffer[0]
	var packet protocol.DataPacket = nil

	if server.sessions.Exists(addr) {
		switch {
		case pid&packets.BITFLAG_ACK != 0:
			packet = packets.NewACK()
			break
		case pid&packets.BITFLAG_NAK != 0:
			packet = packets.NewNACK()
			break
		case pid&packets.BITFLAG_VALID != 0:
			packet = packets.NewDatagram()
			break
		}
	} else {
		switch pid {
		case packets.UNCONNECTED_PING:
			packet = packets.NewUnconnectedPing()
			println(pid)
		case packets.OPEN_CONNECTION_REQUEST_1:
			packet = packets.NewOpenConnectionRequest1()
		case packets.OPEN_CONNECTION_REQUEST_2:
			packet = packets.NewOpenConnectionRequest2()
		}
	}
	if packet == nil {
		println(buffer[0])
		if bytes.Equal(buffer[0:2], query.Header) {
			if !api.GetServer().GetConfig().Get("enable-query", true).(bool) {
				return
			}
			println("query")
		}
		return
	}
	packet.SetBuffer(buffer)
	packet.Decode()
	if packet.HasMagic() {
		HandleUnconnectedMessage(packet, addr, server)
	} else if session, exists := server.sessions.GetSession(addr); exists {

		if datagram, ok := packet.(*packets.Datagram); ok {
			session.ReceiveWindow.AddDatagram(datagram)
		} else if ack, ok := packet.(*packets.ACK); ok {
			session.HandleACK(ack)
		} else if nack, ok := packet.(*packets.NACK); ok {
			session.HandleNACK(nack)
		}
	}
}

func (server *RaknetServer) GetId() int64 {
	return server.id
}

func (sessions Sessions) Exists(addr *net.UDPAddr) bool {
	_, exists := sessions[fmt.Sprint(addr)]
	return exists
}

func (sessions Sessions) GetSession(addr *net.UDPAddr) (*Session, bool) {
	session, exists := sessions[fmt.Sprint(addr)]
	return session, exists
}
