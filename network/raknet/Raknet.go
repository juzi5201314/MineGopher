package raknet

import ()

var MAGIC []byte = []byte{
	0x00, 0xff, 0xff, 0x00,
	0xfe, 0xfe, 0xfe, 0xfe,
	0xfd, 0xfd, 0xfd, 0xfd,
	0x12, 0x34, 0x56, 0x78,
}

const (
	PRIORITY_NORMAL           = 0
	PRIORITY_IMMEDIATE        = 1
	FLAG_NEED_ACK             = 8
	PACKET_ENCAPSULATED       = 0x01
	PACKET_OPEN_SESSION       = 0x02
	PACKET_CLOSE_SESSION      = 0x03
	PACKET_SEND_QUEUE         = 0x05
	PACKET_ACK_NOTIFICATION   = 0x06
	PACKET_SET_OPTION         = 0x07
	PACKET_RAW                = 0x08
	PACKET_BLOCK_ADDRESS      = 0x09
	PACKET_UNBLOCK_ADDRESS    = 0x10
	PACKET_SHUTDOWN           = 0x7e
	PACKET_EMERGENCY_SHUTDOWN = 0x7f
)
