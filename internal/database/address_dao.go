package database

import (
	"context"
	"eccom-mongo/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type AddressDAOInterface interface {
	CreateAddress(ctx context.Context, address *models.Address) error
	DeleteAddress(ctx context.Context, addressID string) error
	GetAddress(ctx context.Context, addressID string) (*models.Address, error)
	GetAllAddress(ctx context.Context) ([]models.Address, error)
}

type addressDAO struct {
	collection *mongo.Collection
}

func NewAddressDAO(db *mongo.Database) AddressDAOInterface {
	return &addressDAO{
		collection: db.Collection("address"),
	}
}

func (a *addressDAO) CreateAddress(ctx context.Context, address *models.Address) error {

	_, err := a.collection.InsertOne(ctx, address)
	if err != nil {
		return err
	}

	return nil

}

func (a *addressDAO) DeleteAddress(ctx context.Context, addressID string) error {

	_, err := a.collection.DeleteOne(ctx, map[string]string{"_id": addressID})
	if err != nil {
		return err
	}

	return nil

}

func (a *addressDAO) GetAddress(ctx context.Context, addressID string) (*models.Address, error) {

	var address models.Address
	err := a.collection.FindOne(ctx, map[string]string{"_id": addressID}).Decode(&address)
	if err != nil {
		return nil, err
	}

	return &address, nil

}

func (a *addressDAO) GetAllAddress(ctx context.Context) ([]models.Address, error) {

	var addresses []models.Address
	cursor, err := a.collection.Find(ctx, map[string]string{})
	if err != nil {
		return nil, err

	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var address models.Address
		cursor.Decode(&address)
		addresses = append(addresses, address)
	}

	return addresses, nil

}
