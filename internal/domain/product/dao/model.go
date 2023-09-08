package dao

import (
	"github.com/Amore14rn/888Starz/internal/domain/product/model"
	"time"
)

type ProductStarage struct {
	ID          string           `json:"id"`
	Description string           `json:"description"`
	Tags        []string         `json:"tags"`
	Quantity    int              `json:"quantity"`
	History     []ProductHistory `json:"history"`
}

type ProductHistory struct {
	Price     float64
	Timestamp time.Time
}

func (p *ProductStarage) AddHistory(price float64, timestamp time.Time) {
	p.History = append(p.History, ProductHistory{
		Price:     price,
		Timestamp: timestamp,
	})
}

func convertProductHistory(history []ProductHistory) []model.ProductHistory {
	var modelHistory []model.ProductHistory
	for _, h := range history {
		modelHistory = append(modelHistory, model.ProductHistory{
			Price:     h.Price,
			Timestamp: h.Timestamp,
		})
	}
	return modelHistory
}

func (p *ProductStarage) ToDomain() model.Product {

	modelProductHistory := convertProductHistory(p.History)

	return model.Product{
		ID:          p.ID,
		Description: p.Description,
		Tags:        p.Tags,
		Quantity:    p.Quantity,
		History:     modelProductHistory,
	}
}
