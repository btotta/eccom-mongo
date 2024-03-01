package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Logradouro  string             `bson:"logradouro"`
	Numero      int                `bson:"numero"`
	Bairro      string             `bson:"bairro"`
	Cidade      string             `bson:"cidade"`
	Estado      string             `bson:"estado"`
	CEP         string             `bson:"cep"`
	MainAddress bool               `bson:"main_address"`
}
