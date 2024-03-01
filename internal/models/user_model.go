package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Nome         string             `bson:"nome"`
	Sobrenome    string             `bson:"sobrenome"`
	CPF          string             `bson:"cpf"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"password_hash"`
	AddressID    primitive.ObjectID `bson:"address_id,omitempty"`
}
