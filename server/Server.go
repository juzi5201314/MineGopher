package server

import (
	networkapi "github.com/juzi5201314/MineGopher/api/network"
	raknetapi "github.com/juzi5201314/MineGopher/api/network/raknet"
	"github.com/juzi5201314/MineGopher/api/player"
	api "github.com/juzi5201314/MineGopher/api/server"
	"github.com/juzi5201314/MineGopher/level"
	"github.com/juzi5201314/MineGopher/network"
	raknet "github.com/juzi5201314/MineGopher/network/raknet/server"
	"github.com/juzi5201314/MineGopher/network/webconsole"
	"github.com/juzi5201314/MineGopher/utils"
	"os"
	"strconv"
	"time"
	"github.com/juzi5201314/MineGopher/plugin"
)

const (
	ServerName    = "MineGopher"
	ServerVersion = "0.0.1"
)

type Server struct {
	isRunning         bool
	tick              int64
	logger            *utils.Logger
	pluginPath        string
	playersPath       string
	themePath         string
	worldsPath        string
	behaviorPacksPath string
	resourecePackPath string
	serverPath        string
	config            *utils.Config
	network           networkapi.NetWork
	ip                string
	port              int
	raknetServer      raknetapi.RaknetServer
	pluginLoader *plugin.PluginLoader

	levels       map[string]*level.Level
	defaultLevel string
}

func New(serverPath string, config *utils.Config, logger *utils.Logger) *Server {
	server := new(Server)
	api.SetServer(server)

	server.serverPath = serverPath
	server.config = config
	server.logger = logger
	server.pluginPath = serverPath + "/plugins/"
	server.themePath = serverPath + "/theme/"
	server.playersPath = serverPath + "/players/"
	server.worldsPath = serverPath + "/worlds/"
	server.resourecePackPath = serverPath + "/resoureces_pack/"
	server.levels = map[string]*level.Level{}

	server.ip = config.Get("server-ip", "0.0.0.0").(string)
	server.port = config.Get("server-port", 19132).(int)

	server.pluginLoader = plugin.NewLoader(server.pluginPath)
	//s.LevelManager = level.NewManager(serverPath)
	//server.CommandManager = commands.NewManager()
	//server.CommandReader = command.NewCommandReader(os.Stdin)
	/*
		s.SessionManager = packet.NewSessionManager()
		s.NetworkAdapter = packet.NewNetworkAdapter(s.SessionManager)
		s.NetworkAdapter.GetRakLibManager().PongData = s.GeneratePongData()
		s.NetworkAdapter.GetRakLibManager().RawPacketFunction = s.HandleRaw
		s.NetworkAdapter.GetRakLibManager().DisconnectFunction = s.HandleDisconnect

		s.RegisterDefaultProtocols()

		s.PackManager = packs.NewManager(serverPath)

		s.PermissionManager = permissions.NewManager()

		s.PluginManager = NewPluginManager(s)

		s.QueryManager = query.NewManager()

		if config.UseEncryption {
			var curve = elliptic.P384()

			var err error
			s.privateKey, err = ecdsa.GenerateKey(curve, rand.Reader)
			text.DefaultLogger.LogError(err)

			if !curve.IsOnCurve(s.privateKey.X, s.privateKey.Y) {
				text.DefaultLogger.Error("Invalid private key generated")
			}

			var token = make([]byte, 128)
			rand.Read(token)
			s.token = token
		}

		return s
	*/
	return server
}

func (server *Server) IsRunning() bool {
	return server.isRunning
}

func (server *Server) Start() {
	if server.isRunning {
		panic("The server has beem started!")
	}
	server.mkdirs()
	server.logger.Info("MineGopher " + ServerVersion + ", running on " + server.serverPath)
	server.isRunning = true

	server.defaultLevel = server.config.Get("level-name", "world").(string)
	dl := level.NewLevel(server.worldsPath+server.defaultLevel, server.defaultLevel)
	server.levels[server.defaultLevel] = dl

	server.network = network.New()
	server.network.SetName(server.config.Get("motd", "MineGopher Server For Minecraft: PE").(string))
	server.raknetServer = raknet.New(server.GetIp(), server.GetPort())
	server.raknetServer.Start()

	server.logger.Info("RakNetServer Listen " + server.GetIp() + ":" + strconv.Itoa(server.GetPort()))

	if server.config.Get("webconsole", true).(bool) {
		webconsole.Start()
	}

	server.pluginLoader.LoadPlugins()

	server.config.Save()
}

func (server *Server) Shutdown() {
	if !server.isRunning {
		return
	}
	for _, l := range server.levels {
		l.GetDimension().Save()
	}
	server.logger.Info("Server stopped.")
	server.isRunning = false
	server.logger.Close()
}

func (server *Server) GetConfig() *utils.Config {
	return server.config
}

func (server *Server) GetAllPlayer() map[string]player.Player {
	return server.raknetServer.GetPlayers()
}

func (server *Server) GetNetWork() networkapi.NetWork {
	return server.network
}

func (server *Server) GetRaknetServer() raknetapi.RaknetServer {
	return server.raknetServer
}

func (server *Server) GetName() string {
	return ServerName
}

func (server *Server) GetLogger() *utils.Logger {
	return server.logger
}

func (server *Server) GetLevels() map[string]*level.Level {
	return server.levels
}

func (server *Server) GetLevel(name string) *level.Level {
	return server.levels[name]
}

func (server *Server) GetDefaultLevel() *level.Level {
	return server.GetLevel(server.defaultLevel)
}

func (server *Server) GetPath() string {
	return server.serverPath
}

func (server *Server) ScheduleRepeatingTask(fn func(), d time.Duration) *time.Ticker {
	ticker := time.NewTicker(d)
	go func() {
		for range ticker.C {
			fn()
		}
	}()
	return ticker
}


func (server *Server) ScheduleDelayedTask(fn func(), d time.Duration) *time.Timer {
	return time.AfterFunc(d, fn)
}
/*
// GetMinecraftVersion returns the latest Minecraft game version.
// It is prefixed with a 'v', for example: "v1.2.10.1"
func (server *Server) GetMinecraftVersion() string {
	return info.LatestGameVersion
}

// GetMinecraftNetworkVersion returns the latest Minecraft network version.
// For example: "1.2.10.1"
func (server *Server) GetMinecraftNetworkVersion() string {
	return info.LatestGameVersionNetwork
}

// HasPermission returns if the server has a given permission.
// Always returns true to satisfy the ICommandSender interface.
func (server *Server) HasPermission(string) bool {
	return true
}

// SendMessage sends a message to the server to satisfy the ICommandSender interface.
func (server *Server) SendMessage(message ...interface{}) {
	text.DefaultLogger.Notice(message)
}

// GetEngineName returns 'minegopher'.
func (server *Server) GetEngineName() string {
	return minegopherName
}

// GetName returns the LAN name of the server specified in the configuration.
func (server *Server) GetName() string {
	return server.Config.ServerName
}

// GetPort returns the port of the server specified in the configuration.
func (server *Server) GetPort() uint16 {
	return server.Config.ServerPort
}

// GetAddress returns the IP address specified in the configuration.
func (server *Server) GetAddress() string {
	return server.Config.ServerIp
}

// GetMaximumPlayers returns the maximum amount of players on the server.
func (server *Server) GetMaximumPlayers() uint {
	return server.Config.MaximumPlayers
}

// Returns the Message Of The Day of the server.
func (server *Server) GetMotd() string {
	return server.Config.ServerMotd
}

// GetCurrentTick returns the current tick the server is on.
func (server *Server) GetCurrentTick() int64 {
	return server.tick
}

// BroadcastMessageTo broadcasts a message to all receivers.
func (server *Server) BroadcastMessageTo(receivers []*packet.MinecraftSession, message ...interface{}) {
	for _, session := range receivers {
		session.SendMessage(message)
	}
	text.DefaultLogger.LogChat(message)
}

// Broadcast broadcasts a message to all players and the console in the server.
func (server *Server) BroadcastMessage(message ...interface{}) {
	for _, session := range server.SessionManager.GetSessions() {
		session.SendMessage(message)
	}
	text.DefaultLogger.LogChat(message)
}

// GetPrivateKey returns the ECDSA private key of the server.
func (server *Server) GetPrivateKey() *ecdsa.PrivateKey {
	return server.privateKey
}

// GetPublicKey returns the ECDSA public key of the private key of the server.
func (server *Server) GetPublicKey() *ecdsa.PublicKey {
	return &server.privateKey.PublicKey
}

// GetServerToken returns the server token byte sequence.
func (server *Server) GetServerToken() []byte {
	return server.token
}

// GenerateQueryResult returns the query data of the server in a byte array.
func (server *Server) GenerateQueryResult() query.Result {
	var plugs []string
	for _, plug := range server.PluginManager.GetPlugins() {
		plugs = append(plugs, plug.GetName()+" v"+plug.GetVersion())
	}

	var ps []string
	for name := range server.SessionManager.GetSessions() {
		ps = append(ps, name)
	}

	var result = query.Result{
		MOTD:           server.GetMotd(),
		ListPlugins:    server.Config.AllowPluginQuery,
		PluginNames:    plugs,
		PlayerNames:    ps,
		GameMode:       "SMP",
		Version:        server.GetMinecraftVersion(),
		ServerEngine:   server.GetEngineName(),
		WorldName:      server.LevelManager.GetDefaultLevel().GetName(),
		OnlinePlayers:  int(server.SessionManager.GetSessionCount()),
		MaximumPlayers: int(server.Config.MaximumPlayers),
		Whitelist:      "off",
		Port:           server.Config.ServerPort,
		Address:        server.Config.ServerIp,
	}

	return result
}

// HandleRaw handles a raw packet, for instance a query packet.
func (server *Server) HandleRaw(packet []byte, addr *net2.UDPAddr) {
	if string(packet[0:2]) == string(query.Header) {
		if !server.Config.AllowQuery {
			return
		}

		var q = query.NewFromRaw(packet, addr)
		q.DecodeServer()

		server.QueryManager.HandleQuery(q)
		return
	}
	text.DefaultLogger.Debug("Unhandled raw packet:", hex.EncodeToString(packet))
}

// HandleDisconnect handles a disconnection from a session.
func (server *Server) HandleDisconnect(s *server.Session) {
	text.DefaultLogger.Debug(s, "disconnected!")
	session, ok := server.SessionManager.GetSessionByRakNetSession(s)

	server.SessionManager.RemoveMinecraftSession(session)
	if !ok {
		return
	}

	if session.GetPlayer().Dimension != nil {
		for _, online := range server.SessionManager.GetSessions() {
			online.SendPlayerList(data.ListTypeRemove, map[string]protocol.PlayerListEntry{online.GetPlayer().GetName(): online.GetPlayer()})
		}

		session.GetPlayer().DespawnFromAll()

		session.GetPlayer().Close()

		server.BroadcastMessage(text.Yellow+session.GetDisplayName(), "has left the server")
	}
}

// GeneratePongData generates the raknet pong data for the UnconnectedPong RakNet packet.
func (server *Server) GeneratePongData() string {
	return fmt.Sprint("MCPE;", server.GetMotd(), ";", info.LatestProtocol, ";", server.GetMinecraftNetworkVersion(), ";", server.SessionManager.GetSessionCount(), ";", server.Config.MaximumPlayers, ";", server.NetworkAdapter.GetRakLibManager().ServerId, ";", server.GetEngineName(), ";Creative;")
}

// Tick ticks the entire server. (Levels, scheduler, raknet server etc.)
// Internal. Not to be used by plugins.
func (server *Server) Tick() {
	if !server.isRunning {
		return
	}
	if server.tick%20 == 0 {
		server.QueryManager.SetQueryResult(server.GenerateQueryResult())
		server.NetworkAdapter.GetRakLibManager().PongData = server.GeneratePongData()
	}

	for _, session := range server.SessionManager.GetSessions() {
		session.Tick()
	}

	for range server.LevelManager.GetLevels() {
		//level.Tick()
	}
	server.tick++
}


func (server *Server) GetCommandManager() command {
	return server.CommandManager
}

**/

func (server *Server) Tick() {
	for _, p := range server.GetAllPlayer() {
		p.Tick()
	}
}

func (server *Server) mkdirs() {
	os.Mkdir(server.playersPath, 0700)
	os.Mkdir(server.pluginPath, 0700)
	//os.Mkdir(server.behaviorPacksPath, 0700)
	os.Mkdir(server.resourecePackPath, 0700)
	os.Mkdir(server.worldsPath, 0700)
	os.Mkdir(server.themePath, 0700)
}

func (server *Server) GetIp() string {
	return server.ip
}

func (server *Server) GetPort() int {
	return server.port
}
