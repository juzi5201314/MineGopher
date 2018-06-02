package api

import (
	"github.com/juzi5201314/MineGopher/utils"
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
	GetNetWork() NetWork
	GetRaknetServer() RaknetServer
	GetLogger() *utils.Logger
}
