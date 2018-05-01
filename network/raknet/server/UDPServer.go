package server

import (
	"net"
	"strconv"
)

func NewUDPServer(ip string, port int) *UDPServer {
	server := &UDPServer{}
	server.queue = make(chan  string, 1024)
	addr, err := net.ResolveUDPAddr("udp", ip + ":" + strconv.Itoa(port))
	server.conn, err = net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	return server
}

type UDPServer struct {
	conn net.Conn
	queue chan string
}

func (server *UDPServer) close() {
	server.conn.Close()
}

func (server *UDPServer) ReadPacket() string {
	return <- server.queue
}

func (server *UDPServer) WritePacket(data []byte) int {
	server.conn.Write(data)
	return len(data)
}