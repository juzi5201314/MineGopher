package plugin

import (
	"io/ioutil"
	"github.com/juzi5201314/MineGopher/api/server"
	"strings"
	"plugin"
)

type PluginLoader struct {
	path string
}

func NewLoader(path string) *PluginLoader {
	return &PluginLoader{path}
}

func (loader *PluginLoader) LoadPlugins() {
	dir, err := ioutil.ReadDir(loader.path)
	if err != nil {
		server.GetServer().GetLogger().PanicError(err)
	}
	for _, fn := range dir {
		if strings.HasSuffix(fn.Name(), ".so") || strings.HasSuffix(fn.Name(), ".dll") {
			loader.LoadPlugin(fn.Name())
		}
	}
}

func (loader *PluginLoader) LoadPlugin(name string) {
	p, err := plugin.Open(loader.path + name)
	if err != nil {
		server.GetServer().GetLogger().PanicError(err)
	}
	pl := newPluginBase(p, name)
	pl.load()
}