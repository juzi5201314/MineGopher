package server

import (
	"github.com/juzi5201314/MineGopher/api/network"
	"github.com/juzi5201314/MineGopher/api/network/raknet"
	"github.com/juzi5201314/MineGopher/api/player"
	"github.com/juzi5201314/MineGopher/level"
	"github.com/juzi5201314/MineGopher/utils"
	"time"
)

var server Server = nil

func SetServer(s Server) {
	if server != nil {
		return
	}
	server = s
}

func GetServer() Server {
	return server
}

type Server interface {
	IsRunning() bool
	GetName() string
	GetConfig() *utils.Config
	GetNetWork() network.NetWork
	GetRaknetServer() raknet.RaknetServer
	GetLogger() *utils.Logger
	GetLevels() map[string]*level.Level
	GetLevel(string) *level.Level
	GetDefaultLevel() *level.Level
	GetAllPlayer() map[string]player.Player
	GetPath() string
	GetIp() string
	GetPort() int
	ScheduleRepeatingTask(fn func(), d time.Duration) *time.Ticker
	ScheduleDelayedTask(fn func(), d time.Duration) *time.Timer
}
