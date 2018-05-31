package server

import (
	"net"
)

func NewUDPServer(ip string, port int) *UDPServer {
	server := new(UDPServer)
	server.ip = ip
	server.port = port
	return server
}

type UDPServer struct {
	ip   string
	port int
	conn *net.UDPConn
}

func (server *UDPServer) Start() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP(server.ip), Port: server.port})
	server.conn = conn
	if err != nil {
		panic(err)
	}
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
