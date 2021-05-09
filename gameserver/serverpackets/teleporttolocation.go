package serverpackets

import (
	"../../gameserver/models"
	"github.com/frostwind/l2go/packets"
)


func TeleportToLocation(coords models.Coord) []byte {
	c := coords
	buffer := packets.NewBuffer()
	buffer.WriteByte(0x38)   // Packet type: TeleportToLocation
	//38
	//XX XX XX XX	// ID чара
	//XX XX XX XX	// X
	//XX XX XX XX	// Y
	//XX XX XX XX	// Z

	buffer.WriteUInt32(268474300) // Object ID
	buffer.WriteUInt32(c.DstX) //
	buffer.WriteUInt32(c.DstY) //
	buffer.WriteUInt32(c.DstZ) //
	return buffer.Bytes()

}
