package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty"`
	Name         string               `bson:"name"`
	LastName     string               `bson:"last_name"`
	Document     string               `bson:"document"`
	Email        string               `bson:"email"`
	PasswordHash string               `bson:"password_hash"`
	Address      []primitive.ObjectID `bson:"address_id,omitempty"`
}
