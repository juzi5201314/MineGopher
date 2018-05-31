type NACK struct {
*AcknowledgementPacket
}

func NewNACK() *NACK {
return &NACK{
&AcknowledgementPacket{NewPacket(FlagDatagramNack), []uint32{}}
}
}
