package mgplugin

import (
	"io/ioutil"
	"github.com/juzi5201314/MineGopher/api/server"
	"strings"
	"plugin"
)

type PluginManager struct{
	plugins map[string]*GopherPlugin
}


func NewPluginManager() *PluginManager{
	mgr := new(PluginManager)
	return mgr
}

func (mgr *PluginManager) LoadPlugins() {
	pluginPath := server.GetServer().GetPluginPath()
	dir,_ := ioutil.ReadDir(pluginPath)
	for _,file := range dir{
		if !file.IsDir() && strings.HasSuffix(file.Name(),".so"){
			if !mgr.LoadPlugin(pluginPath + file.Name()){
				server.GetServer().GetLogger().Error("无法加载插件" + file.Name())
			}
		}
	}
}

func (mgr *PluginManager) LoadPlugin(path string)  bool{
	mgplugin,err := plugin.Open(path)
	if err != nil{
		return false
	}

	sym,err := mgplugin.Lookup("onLoad")

	if err != nil{
		return false
	}
	sym.(func())()
	server.GetServer().GetLogger().Info("Loading Plugin....")

	return true

}