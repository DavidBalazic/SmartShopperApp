package repository

import (
	"AuditService/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoAuditlogRepo struct {
	collection *mongo.Collection
}

func NewMongoRepository(collection *mongo.Collection) AuditRepository {
	return &mongoAuditlogRepo{collection}
}

func (r *mongoAuditlogRepo) SaveLog(log models.AuditLog) error {
	_, err := r.collection.InsertOne(context.Background(), log)
	return err
}