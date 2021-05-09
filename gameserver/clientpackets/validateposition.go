package clientpackets

import (
	"github.com/frostwind/l2go/packets"
)


type ValidateCoord struct {
	CurX uint32
	CurY uint32
	CurZ uint32
	Heading uint32
	Data uint32
}

func ValidatePosition(request []byte) ValidateCoord {
	var packet = packets.NewReader(request)
	var c ValidateCoord

	c.CurX = packet.ReadUInt32()
	c.CurY = packet.ReadUInt32()
	c.CurZ = packet.ReadUInt32()
	c.Heading = packet.ReadUInt32()
	c.Data = packet.ReadUInt32()

	return c
}
