package repo

import (
	"context"
	"strings"
	"regexp"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/DavidBalazic/SmartShopperApp/internal/models"
	"github.com/DavidBalazic/SmartShopperApp/config"
)

type MongoProductRepository struct {
	db *mongo.Collection
}

func NewMongoProductRepository() *MongoProductRepository {
	db := config.GetDB().Collection("products")
	return &MongoProductRepository{db: db}
}

func (r *MongoProductRepository) FindCheapestProduct(name string) (models.Product, error) {
	log.Printf("FindCheapestProduct called with name: %s", name)
	name = strings.ToLower(name)
	words := strings.Fields(name)

	var orConditions []bson.M
	for _, word := range words {
		orConditions = append(orConditions, bson.M{"name": bson.M{"$regex": regexp.QuoteMeta(word), "$options": "i"}})
	}

	var product models.Product
	err := r.db.FindOne(context.TODO(), bson.M{"$or": orConditions}).Decode(&product)
	if err != nil {
		log.Printf("FindCheapestProduct Error finding cheapest product: %v", err)
		return models.Product{}, err
	}
	
	return product, err
}

func (r *MongoProductRepository) FindCheapestProductByStore(name, store string) (models.Product, error) {
	log.Printf("FindCheapestProductByStore called with name: %s, store: %s", name, store)
	name = strings.ToLower(name)
	store = strings.ToLower(store)
	words := strings.Fields(name)

	var orConditions []bson.M
	for _, word := range words {
		orConditions = append(orConditions, bson.M{"name": bson.M{"$regex": regexp.QuoteMeta(word), "$options": "i"}})
	}

	var product models.Product
	err := r.db.FindOne(context.TODO(), bson.M{"$or": orConditions, "store": bson.M{"$regex": regexp.QuoteMeta(store), "$options": "i"}}).Decode(&product)
	if err != nil {
		log.Printf("FindCheapestProductByStore Error finding cheapest product in store: %v", err)
		return models.Product{}, err
	}
	return product, err
}

func (r *MongoProductRepository) FindAllProductPrices(name string) ([]models.Product, error) {
	log.Printf("FindAllProductPrices called with name: %s", name)
	name = strings.ToLower(name)
	words := strings.Fields(name)

	var orConditions []bson.M
	for _, word := range words {
		orConditions = append(orConditions, bson.M{"name": bson.M{"$regex": regexp.QuoteMeta(word), "$options": "i"}})
	}

	var products []models.Product
	cursor, err := r.db.Find(context.TODO(), bson.M{"$or": orConditions})
	if err != nil {
		log.Printf("FindAllProductPrices Error finding all product prices: %v", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var product models.Product
		cursor.Decode(&product)
		products = append(products, product)
	}
	return products, nil
}