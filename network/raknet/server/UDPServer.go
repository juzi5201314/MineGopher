package server

import (
	"net"
)

func NewUDPServer(ip string, port int) *UDPServer {
	server := new(UDPServer)
	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP(ip), Port: port})
	server.conn = conn
	if err != nil {
		panic(err)
	}
	return server
}

type UDPServer struct {
	conn *net.UDPConn
}

func (server *UDPServer) close() {
	server.conn.Close()
}

func (server *UDPServer) Read(buffer []byte) (int, *net.UDPAddr, error) {
	return server.conn.ReadFromUDP(buffer)
}

func (server *UDPServer) Write(addr *net.UDPAddr, buffer []byte) (int, error) {
	return server.conn.WriteToUDP(buffer, addr)
}