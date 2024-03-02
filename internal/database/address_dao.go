package database

import (
	"context"
	"eccom-mongo/internal/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AddressDAOInterface interface {
	CreateAddress(ctx context.Context, address *models.Address) (*models.Address, error)
	DeleteAddress(ctx context.Context, addressID string) error
	GetAddress(ctx context.Context, addressID string) (*models.Address, error)
	GetAllAddressByUserID(ctx context.Context, user *models.User) ([]models.Address, error)
	MarkAddressAsMain(ctx context.Context, addressID string, user *models.User) (*models.Address, error)
}

type addressDAO struct {
	collection *mongo.Collection
}

func NewAddressDAO(db *mongo.Database) AddressDAOInterface {
	return &addressDAO{
		collection: db.Collection("address"),
	}
}

func (a *addressDAO) CreateAddress(ctx context.Context, address *models.Address) (*models.Address, error) {

	result, err := a.collection.InsertOne(ctx, address)
	if err != nil {
		return nil, err
	}

	address.ID = result.InsertedID.(primitive.ObjectID)

	return address, nil
}

func (a *addressDAO) DeleteAddress(ctx context.Context, addressID string) error {

	objID, err := primitive.ObjectIDFromHex(addressID)
	if err != nil {
		return err
	}

	_, err = a.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil

}

func (a *addressDAO) GetAddress(ctx context.Context, addressID string) (*models.Address, error) {

	objID, err := primitive.ObjectIDFromHex(addressID)
	if err != nil {
		return nil, err
	}

	var address models.Address

	err = a.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&address)
	if err != nil {
		return nil, err
	}

	return &address, nil

}

func (a *addressDAO) GetAllAddressByUserID(ctx context.Context, user *models.User) ([]models.Address, error) {

	var addresses []models.Address
	cursor, err := a.collection.Find(ctx, bson.M{"user_id": user.ID})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &addresses); err != nil {
		return nil, err
	}

	return addresses, nil
}

func (a *addressDAO) MarkAddressAsMain(ctx context.Context, addressID string, user *models.User) (*models.Address, error) {

	objID, err := primitive.ObjectIDFromHex(addressID)
	if err != nil {
		return nil, err
	}

	addresses, err := a.GetAllAddressByUserID(ctx, user)
	if err != nil {
		return nil, err
	}

	var toBeMainAddress *models.Address

	for _, address := range addresses {
		if address.ID == objID {
			toBeMainAddress = &address
			break
		}
	}

	if toBeMainAddress == nil {
		return nil, errors.New("address not found")
	}

	for _, address := range addresses {
		if address.ID != objID {
			address.MainAddress = false
			_, err := a.collection.UpdateOne(ctx, bson.M{"_id": address.ID}, bson.M{"$set": bson.M{"main_address": false}})
			if err != nil {
				return nil, err
			}

		}
	}

	_, err = a.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bson.M{"main_address": true}})
	if err != nil {
		return nil, err
	}

	toBeMainAddress.MainAddress = true

	return toBeMainAddress, nil

}
