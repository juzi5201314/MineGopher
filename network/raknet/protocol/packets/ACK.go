package packets

type ACK struct {
*AcknowledgementPacket
}

func NewACK() *ACK {	
return &ACK{
&AcknowledgementPacket{NewPacket(FlagDatagramAck), []uint32{}}
}
}