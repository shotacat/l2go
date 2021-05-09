package serverpackets

import (
	"github.com/frostwind/l2go/packets"
)

func CharCreateFailPacket(reason uint32) []byte {
	buffer := packets.NewBuffer()
	buffer.WriteByte(0x26)   // Packet type: CharCreateOk
	buffer.WriteUInt32(reason) // Name already exists

	return buffer.Bytes()
}
