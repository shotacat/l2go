package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Character struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Name      string        `bson:"name"`
	Race      uint32		`bson:"race"`
	Sex       uint32		`bson:"Sex"`
	ClassID   uint32		`bson:"ClassID"`
	STR       uint32		`bson:"STR"`
	CON       uint32		`bson:"CON"`
	DEX       uint32		`bson:"DEX"`
	INT       uint32		`bson:"INT"`
	MEN       uint32		`bson:"MEN"`
	WIT       uint32		`bson:"WIT"`
	HairStyle uint32		`bson:"HairStyle"`
	HairColor uint32		`bson:"HairColor"`
	Face      uint32		`bson:"Face"`
	//Coords	Coord `bson:"Coords"`
}