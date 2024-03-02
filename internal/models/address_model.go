package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Street       string             `bson:"street"`
	Number       string             `bson:"number"`
	Complement   string             `bson:"complement"`
	Neighborhood string             `bson:"neighborhood"`
	City         string             `bson:"city"`
	State        string             `bson:"state"`
	Country      string             `bson:"country"`
	ZipCode      string             `bson:"zip_code"`
	MainAddress  bool               `bson:"main_address"`
	UserID       primitive.ObjectID `bson:"user_id"`
}
