package packets

const (
	FLAG_DATAGRAM_ACK  = 0xc0
	FLAG_DATAGRAM_NACK = 0xa0

	CONNECTED_PING            = 0x00
	CONNECTED_PONG            = 0x03
	UNCONNECTED_PING          = 0x01
	UNCONNECTED_PONG          = 0x1c
	OPEN_CONNECTION_REQUEST_1 = 0x05
	OPEN_CONNECTION_REPLY_1   = 0x06
	OPEN_CONNECTION_REQUEST_2 = 0x07
	OPEN_CONNECTION_REPLY_2   = 0x08
	CONNECTION_REQUEST        = 0x09
	CONNECTION_ACCEPT         = 0x10
	INCOMING_CONNECTION       = 0x13
	DISCONNECT_NOTIFICATION   = 0x15
)
