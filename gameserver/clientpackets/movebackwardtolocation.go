package clientpackets

import (
	"../../gameserver/models"
	"github.com/frostwind/l2go/packets"
)


func MoveBackwardToLocation(request []byte) models.Coord {
	var packet = packets.NewReader(request)
	var c models.Coord

	c.DstX = packet.ReadUInt32()
	c.DstY = packet.ReadUInt32()
	c.DstZ = packet.ReadUInt32()
	c.CurX = packet.ReadUInt32()
	c.CurY = packet.ReadUInt32()
	c.CurZ = packet.ReadUInt32()
	c.MoveType = packet.ReadUInt32()

	return c
}
