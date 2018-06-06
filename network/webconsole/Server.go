package webconsole

import (
	"strconv"
	"encoding/json"

	"github.com/juzi5201314/MineGopher/api/server"

	"github.com/labstack/echo"
	"github.com/gorilla/websocket"
	"io/ioutil"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

const (
	OK          = iota
	ERROR
	PWD_ERROR
	PWD_SUCCESS
)

func Start() {
	e := echo.New()
	e.Logger.SetOutput(ioutil.Discard)
	e.HideBanner = true
	if server.GetServer().GetConfig().Get("webconsole-static", true).(bool) {
		e.Static("/", server.GetServer().GetPath()+"/theme")
	}
	e.GET("/ws", handle)
	go e.Start(":" + strconv.Itoa(server.GetServer().GetPort()))
}

func handle(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		server.GetServer().GetLogger().PanicError(err)
		return err
	}
	testpwd := func(pwd string) bool {
		return pwd == server.GetServer().GetConfig().Get("webconsole-password", "").(string)
	}

	for {
		_, pwd, err := ws.ReadMessage()
		if err != nil {
			return nil
		}
		println(string(pwd))
		if testpwd(string(pwd)) {
			break
		}
		j, _ := json.Marshal(map[string]interface{}{
			"status": PWD_ERROR,
		})
		ws.WriteMessage(websocket.TextMessage, j)
	}
	j, _ := json.Marshal(map[string]interface{}{
		"status": PWD_SUCCESS,
	})
	ws.WriteMessage(websocket.TextMessage, j)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}
		println(msg)
	}
	defer ws.Close()
	return nil
}
