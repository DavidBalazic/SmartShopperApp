package repo

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/DavidBalazic/SmartShopperApp/internal/models"
    "github.com/DavidBalazic/SmartShopperApp/internal/database"
)

type ProductRepository interface {
	FindProductById(ctx context.Context, id string) (models.Product, error)
	AddProduct(ctx context.Context, product models.Product) (models.Product, error)
	FindProductsByIds(ctx context.Context, ids []string) ([]models.Product, error)
	AddProducts(ctx context.Context, products []models.Product) ([]models.Product, error)
}

type MongoProductRepository struct {
    Db *mongo.Collection
}

func NewMongoProductRepository() *MongoProductRepository {
    collection := db.GetDB().Collection("products")
	return &MongoProductRepository{Db: collection}
}

func (r *MongoProductRepository) FindProductById(ctx context.Context, id string) (models.Product, error) {
	log.Printf("FindProductByID called with ID: %s", id)

	objectId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        log.Printf("Invalid ObjectId format: %v", err)
        return models.Product{}, err
    }

    var product models.Product
    err = r.Db.FindOne(ctx, bson.M{"_id": objectId}).Decode(&product)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            log.Printf("FindProductByID: No product found with ID: %s", id)
        } else {
            log.Printf("FindProductByID Error: %v", err)
        }
        return models.Product{}, err
    }

	return product, nil
}

func (r *MongoProductRepository) AddProduct(ctx context.Context, product models.Product) (models.Product, error) {
	result, err := r.Db.InsertOne(ctx, product)
	if err != nil {
		log.Printf("AddProduct error: %v", err)
		return models.Product{}, err
	}
	product.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return product, nil
}

func (r *MongoProductRepository) FindProductsByIds(ctx context.Context, ids []string) ([]models.Product, error) {
	log.Printf("FindProductsByIDs called with ID: %s", ids)

    var objectIDs []primitive.ObjectID
    for _, id := range ids {
        objID, err := primitive.ObjectIDFromHex(id)
        if err != nil {
            log.Printf("Invalid ObjectId format: %v", err)
            return nil, err
        }
        objectIDs = append(objectIDs, objID)
    }

    cursor, err := r.Db.Find(ctx, bson.M{"_id": bson.M{"$in": objectIDs}})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

	var products []models.Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *MongoProductRepository) AddProducts(ctx context.Context, products []models.Product) ([]models.Product, error) {
    log.Printf("Adding products: %v", products)

	docs := make([]interface{}, len(products))
	for i, p := range products {
		docs[i] = p
	}
	res, err := r.Db.InsertMany(ctx, docs)
	if err != nil {
		return nil, err
	}
	for i, id := range res.InsertedIDs {
		products[i].ID = id.(primitive.ObjectID).Hex()
	}
	return products, nil
}