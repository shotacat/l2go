package serverpackets

import (
	"github.com/frostwind/l2go/packets"
)

func UserInfoPacket() []byte {
	buffer := packets.NewBuffer()
	buffer.WriteByte(0x04)
	buffer.WriteInt32(11382) // x
	buffer.WriteInt32(16697) // y
	buffer.WriteInt32(-4666) // z
	buffer.WriteUInt32(0) // heading)
	buffer.WriteUInt32(268474300) // objectId
	//buffer.Write([]byte{'y', 0x00, 0x00, 0x00}) // name
	buffer.WriteString("nameЖопа") // name
	buffer.WriteUInt32(4)          // raceId
	buffer.WriteUInt32(1)          // gender
	buffer.WriteUInt32(0x35)       // classId
	buffer.WriteUInt32(1)          // level
	buffer.WriteUInt32(0)          // exp
	buffer.WriteUInt32(39)         // str
	buffer.WriteUInt32(29)         // dex
	buffer.WriteUInt32(45)         // con
	buffer.WriteUInt32(20)         // int
	buffer.WriteUInt32(10)         // wit
	buffer.WriteUInt32(27)         // men
	buffer.WriteUInt32(129)        // maximumHp
	buffer.WriteUInt32(129)        // hp
	buffer.WriteUInt32(39)         // maximumMp
	buffer.WriteUInt32(39)         // mp
	buffer.WriteUInt32(0)          // sp
	buffer.WriteUInt32(47801)      // getLoad()
	buffer.WriteUInt32(83000)      // maximumLoad

	buffer.WriteUInt32(0x28)

	buffer.WriteUInt32(0) // underwear.objectId
	buffer.WriteUInt32(0) // ear.right.objectId
	buffer.WriteUInt32(0) // ear.left.objectId
	buffer.WriteUInt32(0) // neck.objectId
	buffer.WriteUInt32(0) // finger.right.objectId
	buffer.WriteUInt32(0) // finger.left.objectId

	buffer.WriteUInt32(0) // head.objectId
	buffer.WriteUInt32(0) // hand.right.objectId
	buffer.WriteUInt32(0) // hand.left.objectId
	buffer.WriteUInt32(0) // gloves.objectId
	buffer.WriteUInt32(0) // chest.objectId
	buffer.WriteUInt32(0) // legs.objectId
	buffer.WriteUInt32(0) // feet.objectId
	buffer.WriteUInt32(0) // back.objectId
	buffer.WriteUInt32(0) // hand.leftAndRight.objectId

	buffer.WriteUInt32(0) // underwear.itemId
	buffer.WriteUInt32(0) // ear.right.itemId
	buffer.WriteUInt32(0) // ear.left.itemId
	buffer.WriteUInt32(0) // neck.itemId
	buffer.WriteUInt32(0) // finger.right.itemId
	buffer.WriteUInt32(0) // finger.left.itemId

	buffer.WriteUInt32(0) // head.itemId
	buffer.WriteUInt32(0) // hand.right.itemId
	buffer.WriteUInt32(0) // hand.left.itemId
	buffer.WriteUInt32(0) // gloves.itemId
	buffer.WriteUInt32(0) // chest.itemId
	buffer.WriteUInt32(0) // legs.itemId
	buffer.WriteUInt32(0) // feet.itemId
	buffer.WriteUInt32(0) // back.itemId
	buffer.WriteUInt32(0) // hand.leftAndRight.itemId


	buffer.WriteUInt32(72) // pAtk
	buffer.WriteUInt32(33) // pSpd
	buffer.WriteUInt32(33) // pDef
	buffer.WriteUInt32(43) // evasion
	buffer.WriteUInt32(3) // accuracy
	buffer.WriteUInt32(203) // critical

	buffer.WriteUInt32(327) // mAtk
	buffer.WriteUInt32(48) // mSpd
	buffer.WriteUInt32(0) // pSpd
	buffer.WriteUInt32(0) // mDef

	buffer.WriteUInt32(0) // getFlagDisplay()
	buffer.WriteUInt32(0) // karma

	buffer.WriteUInt32(115) // runSpeed
	buffer.WriteUInt32(80) // walkSpeed
	buffer.WriteUInt32(115) // swimSpeed
	buffer.WriteUInt32(80) // swimSpeed
	buffer.WriteUInt32(1000) // runSpeed) // getFloatingRunSpeed
	buffer.WriteUInt32(1000) // walkSpeed) // getFloatingWalkSpeed
	buffer.WriteUInt32(1000) // runSpeed) // getFlyingRunSpeed
	buffer.WriteUInt32(1000) // walkSpeed) // getFlyingWalkSpeed

	buffer.WriteFloat64(3.487196) // maleMovementMultiplier
	buffer.WriteFloat64(1) // maleAttackSpeedMultiplier
	buffer.WriteFloat64(19) // maleCollisionRadius
	buffer.WriteFloat64(19) // maleCollisionHeight


	buffer.WriteUInt32(0) // hairStyle
	buffer.WriteUInt32(0) // hairColor
	buffer.WriteUInt32(0) // face
	buffer.WriteUInt32(1) // gm
	buffer.Write([]byte{'c', 0x00, 0x00, 0x00}) // name
	buffer.WriteUInt32(0) // clanId
	buffer.WriteUInt32(0) // clanCrestId
	buffer.WriteUInt32(0) // allianceId
	buffer.WriteUInt32(0) // allianceCrestId
	buffer.WriteUInt32(0x00) // 0x60 ??? // siege-flags
	buffer.WriteByte(0x00)
	buffer.WriteByte(0) // privateStoreType
	buffer.WriteByte(0) // canCraft
	buffer.WriteUInt32(0) // pk
	buffer.WriteUInt32(0) // pvp
	buffer.Write([]byte{0x00, 0x00}) // name

	buffer.WriteByte(0x00) // 1-find party members

	return buffer.Bytes()
}
