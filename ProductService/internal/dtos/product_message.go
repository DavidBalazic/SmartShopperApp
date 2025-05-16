package dtos

type ProductMessage struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Quantity     float64     `json:"quantity"`
	Unit         string  `json:"unit"`
	Store        string  `json:"store"`
	PricePerUnit float64 `json:"pricePerUnit"`
}