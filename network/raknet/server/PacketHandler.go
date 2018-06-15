package server

import (
	"fmt"
	"github.com/juzi5201314/MineGopher/api/server"
	"github.com/juzi5201314/MineGopher/network/raknet/protocol"
	"github.com/juzi5201314/MineGopher/network/raknet/protocol/packets"
	"net"
	"time"
)

func HandleUnconnectedMessage(packetInterface protocol.DataPacket, addr *net.UDPAddr, server *RaknetServer) {
	switch packet := packetInterface.(type) {
	case *packets.UnconnectedPing:
		handleUnconnectedPing(addr, server)
	case *packets.OpenConnectionRequest1:
		handleOpenConnectionRequest1(packet, addr, server)
	case *packets.OpenConnectionRequest2:
		handleOpenConnectionRequest2(packet, addr, server)
	}
}

func handleUnconnectedPing(addr *net.UDPAddr, rs *RaknetServer) {
	pong := packets.NewUnconnectedPong()
	pong.PingTime = time.Now().Unix()
	pong.ServerId = rs.id
	pong.PongData = server.GetServer().GetNetWork().GetName()
	pong.Encode()
	rs.udpServer.Write(addr, pong.Buffer)
}

func handleOpenConnectionRequest1(request *packets.OpenConnectionRequest1, addr *net.UDPAddr, server *RaknetServer) {
	reply := packets.NewOpenConnectionReply1()
	reply.ServerId = server.id
	reply.MtuSize = request.MtuSize
	reply.Security = false
	reply.Encode()
	server.udpServer.Write(addr, reply.Buffer)
}

func handleOpenConnectionRequest2(request *packets.OpenConnectionRequest2, addr *net.UDPAddr, server *RaknetServer) {
	reply := packets.NewOpenConnectionReply2()
	reply.ServerId = server.GetId()
	reply.MtuSize = request.MtuSize
	reply.UseEncryption = false
	reply.ClientAddress = addr.IP.String()
	reply.ClientPort = uint16(addr.Port)
	reply.Encode()
	server.sessions[fmt.Sprint(addr)] = NewSession(addr, request.MtuSize, server)
	server.udpServer.Write(addr, reply.Buffer)
}
