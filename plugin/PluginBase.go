package plugin

import (
	"plugin"
	"github.com/juzi5201314/MineGopher/api/server"
	"strings"
)

type PluginBase struct {
	*plugin.Plugin
	name string
}

func newPluginBase(p *plugin.Plugin, name string) *PluginBase {
	return &PluginBase{p, strings.Replace(strings.Replace(name, ".dll", "", -1), ".so", "", -1)}
}

func (p *PluginBase) load() {
	fn, err := p.Lookup("OnLoad")
	if err != nil {
		server.GetServer().GetLogger().Error("plugin: ", p.name, ", is no \"OnLoad\" method")
	}
	fn.(func(base *PluginBase))(p)
}

func (p *PluginBase) GetName() string {
	return p.name
}