package database

import (
	"context"
	"eccom-mongo/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ErrNoDocuments = "mongo: no documents in result"
)

type UserDAOInterface interface {
	CreateUser(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByCPF(ctx context.Context, cpf string) (*models.User, error)
	FindByID(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context) ([]models.User, error)
}

type userDAO struct {
	collection *mongo.Collection
}

func NewUserDAO(mongoDB *mongo.Database) *userDAO {
	return &userDAO{
		collection: mongoDB.Collection("users"),
	}
}

func (dao *userDAO) CreateUser(ctx context.Context, user *models.User) error {
	_, err := dao.collection.InsertOne(ctx, user)
	return err
}

func (dao *userDAO) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := dao.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil && err.Error() == ErrNoDocuments {
		return nil, nil
	}
	return &user, err
}

func (dao *userDAO) FindByCPF(ctx context.Context, cpf string) (*models.User, error) {
	var user models.User
	err := dao.collection.FindOne(ctx, bson.M{"cpf": cpf}).Decode(&user)
	if err != nil && err.Error() == ErrNoDocuments {
		return nil, nil
	}
	return &user, err
}

func (dao *userDAO) FindByID(ctx context.Context, id string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user models.User
	err = dao.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *userDAO) Update(ctx context.Context, user *models.User) error {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}
	_, err := dao.collection.UpdateOne(ctx, filter, update)
	return err
}

func (dao *userDAO) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objID}
	result, err := dao.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (dao *userDAO) FindAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	cursor, err := dao.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}
