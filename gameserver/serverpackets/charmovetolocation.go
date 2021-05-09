package serverpackets

import (
	"../../gameserver/models"
	"github.com/frostwind/l2go/packets"
)


func CharMoveToLocation(coords models.Coord) []byte {
	c := coords
	buffer := packets.NewBuffer()
	buffer.WriteByte(0x01)   // Packet type: CharCreateOk

	buffer.WriteUInt32(268474300) // Object ID
	buffer.WriteUInt32(c.DstX) //
	buffer.WriteUInt32(c.DstY) //
	buffer.WriteUInt32(c.DstZ) //
	buffer.WriteUInt32(c.CurX) // curr
	buffer.WriteUInt32(c.CurY) // curr
	buffer.WriteUInt32(c.CurZ) // curr
	return buffer.Bytes()

}
