package network

import "github.com/juzi5201314/minegopher"

type RaknetInterface struct {
	server *minegopher.Server
	network *NetWork
}

func (ri *RaknetInterface) SetNetWork(network *NetWork) {
	ri.network = network
}

func (ri *RaknetInterface) process() bool {

}

func (ri *RaknetInterface) SetName(name string) {

}