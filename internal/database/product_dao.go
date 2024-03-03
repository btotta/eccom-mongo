package database

import (
	"context"
	"eccom-mongo/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductDAOInterface interface {
	CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error)
	GetProduct(ctx context.Context, productSku string) (*models.Product, error)
	SearchProduct(ctx context.Context, querys []string) ([]*models.Product, error)
}

type productDAO struct {
	collection *mongo.Collection
}

func NewProductDAO(mongoDB *mongo.Database) *productDAO {
	return &productDAO{
		collection: mongoDB.Collection("products"),
	}
}

func (dao *productDAO) CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {

	_, err := dao.collection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (dao *productDAO) GetProduct(ctx context.Context, productSku string) (*models.Product, error) {
	var product models.Product
	err := dao.collection.FindOne(ctx, models.Product{Sku: productSku}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (dao *productDAO) SearchProduct(ctx context.Context, querys []string) ([]*models.Product, error) {
	var products []*models.Product
	var searchConditions []bson.M

	for _, query := range querys {
		searchConditions = append(searchConditions, bson.M{"name": bson.M{"$regex": query, "$options": "i"}})
		searchConditions = append(searchConditions, bson.M{"description": bson.M{"$regex": query, "$options": "i"}})
	}

	filter := bson.M{"$or": searchConditions}

	cursor, err := dao.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product models.Product
		cursor.Decode(&product)
		products = append(products, &product)
	}

	return products, nil
}
