package api

type NetWork interface {
	GetName() string
	SetName(string)
	RaknetPacketToMinecraftPaket([]byte) MinecraftPacket
}
