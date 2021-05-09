package serverpackets

import (
	"github.com/frostwind/l2go/packets"
)

func LogoutOkPacket() []byte {

	buffer := packets.NewBuffer()
	buffer.WriteByte(0x96) // Packet type: LogoutOk

	return buffer.Bytes()
}
