package serverpackets

import (
	"github.com/frostwind/l2go/packets"
)

func CharSelectedPacket() []byte {
	buffer := packets.NewBuffer()
	buffer.WriteByte(0x21)
	buffer.WriteString("nameЖопа") // name
	buffer.WriteUInt32(268474300)          // objectId)
	buffer.WriteString("titleЖопа") // name
	buffer.WriteUInt32(0x55555555)         //
	buffer.WriteUInt32(0)          // clanId)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(1)          // gender)
	buffer.WriteUInt32(4)          // raceId)
	buffer.WriteUInt32(0x35)          // classId)
	buffer.WriteUInt32(0x01)
	buffer.WriteInt32(11382) // x
	buffer.WriteInt32(16697) // y
	buffer.WriteInt32(-4666) // z
	buffer.WriteFloat64(10) // (character.hp)
	buffer.WriteFloat64(10) //(character.mp)
	buffer.WriteUInt32(0)          // sp)
	buffer.WriteUInt32(1)          // exp)
	buffer.WriteUInt32(1)          // level)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0)          // int)
	buffer.WriteUInt32(0)          // str)
	buffer.WriteUInt32(0)          // con)
	buffer.WriteUInt32(0)          // men)
	buffer.WriteUInt32(0)          // dex)
	buffer.WriteUInt32(0)          // wit)

	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)

	buffer.WriteUInt32(0x00) // in-game time

	return buffer.Bytes()
}
