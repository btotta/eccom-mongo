package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// roles enum

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	LastName     string             `bson:"last_name"`
	Document     string             `bson:"document"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"password_hash"`
	Roles        []string           `bson:"roles"`
}
