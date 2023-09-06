package model

import "time"

type Product struct {
	ID          string
	Description string
	Tags        []string
	Quantity    int
	History     []ProductHistory
}

type ProductHistory struct {
	Price     float64
	Timestamp time.Time
}

func (p *Product) AddHistory(price float64, timestamp time.Time) {
	p.History = append(p.History, ProductHistory{
		Price:     price,
		Timestamp: timestamp,
	})
}

func NewProduct(
	ID string,
	description string,
	tags []string,
	quantity int,
	history []ProductHistory,
) Product {
	return Product{
		ID:          ID,
		Description: description,
		Tags:        tags,
		Quantity:    quantity,
		History:     history,
	}
}

type CreateProduct struct {
	ID          string
	Description string
	Tags        []string
	Quantity    int
}

func NewCreateProduct(
	ID string,
	description string,
	tags []string,
	quantity int,
) CreateProduct {
	return CreateProduct{
		ID:          ID,
		Description: description,
		Tags:        tags,
		Quantity:    quantity,
	}
}
