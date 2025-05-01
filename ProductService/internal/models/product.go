package models

type Product struct {
	ID           string  `bson:"_id,omitempty"`
	Name         string  `bson:"name"`
	Description  string  `bson:"description"`
	Price        float64 `bson:"price"`
	Quantity     float64 `bson:"quantity"`
	Unit         string  `bson:"unit"`
	Store        string  `bson:"store"`
	PricePerUnit float64 `bson:"pricePerUnit"`
}