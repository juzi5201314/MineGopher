package network

import (
	"github.com/juzi5201314/minegopher/network/protocol"
	"github.com/juzi5201314/minegopher"
)

type SourceInterface interface {
PutPlayer(minegopher.Player, protocol.DataPacket) int
GetNetworkLatency(minegopher.Player) int
Close(minegopher.Player, string)
SetName(string)
process() bool
shutdown()
emergencyShutdown()
}