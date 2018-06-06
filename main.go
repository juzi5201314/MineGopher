package main

import (
	"fmt"
	"github.com/juzi5201314/MineGopher/server"
	"github.com/juzi5201314/MineGopher/utils"
"github.com/google/uuid"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

func main() {
	startTime := time.Now()
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(utils.ColorString(utils.Red + "[Error]" + e.(string)).ToANSI())
		}
	}()

	path := GetServerPath()
	os.Mkdir(path+"logs/", 0700)
	file, _ := os.OpenFile(path+"logs/"+time.Now().Format("2006-01-02_15.04.05")+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0700)
	logger := utils.NewLogger(file)
	config := utils.NewConfig(path+"minegopher.yml", utils.YMAL, map[string]interface{}{
		"motd":             "MineGopher Server For Minecraft: PE",
		"server-ip":        "0.0.0.0",
		"server-port":      19132,
		"max-player":       20,
		"gamemode":         0,
		"level-name":       "world",
		"level-seed":       "",
		"enable-query":     true,
		"auto-save-player": true,
		"auto-save-level":  true,
"webconsole": true,
"webconsole-static": true,
"webconsole-password": uuid.New(),
	})

	server := server.New(path, config, logger)
	server.Start()
	signalChan := make(chan os.Signal, 1)
	go func() {
		<-signalChan
		server.Shutdown()
	}()
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	server.GetLogger().Info("Server startup done! Use:", time.Now().Sub(startTime))

	for range time.NewTicker(time.Second / 20).C {
		if !server.IsRunning() {
			break
		}
		server.Tick()
	}
}

func GetServerPath() string {
	executable, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return strings.Replace(filepath.Dir(executable)+"/", `\`, "/", -1)
}
