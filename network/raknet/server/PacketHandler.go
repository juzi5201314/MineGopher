package server

import (
	"github.com/juzi5201314/MineGopher/network/raknet/protocol/packets"
	"net"
	"time"
	"fmt"
	"github.com/juzi5201314/MineGopher/network/raknet/protocol"
	"github.com/juzi5201314/MineGopher/api"
)

const (
	MaximumMTUSize = 1492
	MinimumMTUSize = 576
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

func handleUnconnectedPing(addr *net.UDPAddr, server *RaknetServer) {
	pong := packets.NewUnconnectedPong()
	pong.PingTime = time.Now().Unix()
	pong.ServerId = server.id
	pong.PongData = api.GetServer().GetNetWork().GetName()
	pong.Encode()
	server.udpServer.Write(addr, pong.Buffer)
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
	if request.MtuSize < MinimumMTUSize {
		request.MtuSize = MinimumMTUSize
	} else if request.MtuSize > MaximumMTUSize {
		request.MtuSize = MaximumMTUSize
	}
	reply.MtuSize = request.MtuSize
	reply.UseEncryption = false
	reply.ClientAddress = addr.IP.String()
	reply.ClientPort = uint16(addr.Port)
	reply.Encode()
	server.sessions[fmt.Sprint(addr)] = NewSession(addr, request.MtuSize, server)
	server.udpServer.Write(addr, reply.Buffer)
}