package raknet

import (
	"github.com/juzi5201314/MineGopher/api/player"
)

type RaknetServer interface {
	GetId() int64
	Start()
	GetPlayers() map[string]player.Player
}
