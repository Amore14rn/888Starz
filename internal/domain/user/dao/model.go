package dao

import (
	"database/sql"
	"github.com/Amore14rn/888Starz/internal/domain/user/model"
	"github.com/Amore14rn/888Starz/pkg/utils/pointer"
	"time"
)

type UserStorage struct {
	ID        string       `json:"id"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	FullName  string       `json:"full_name"`
	Age       uint32       `json:"age"`
	IsMarried bool         `json:"is_married"`
	Password  string       `json:"password"`
	CreatedAt time.Time    `json:"created_at"`
	UpdateAt  sql.NullTime `json:"updated_at"`
	Orders    []Order      `json:"orders"`
}

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

type Order struct {
	ID        string
	UserID    string
	Products  []OrderProduct
	Timestamp time.Time
}

type OrderProduct struct {
	ProductID string
	Quantity  int
	Price     float64
}

func convertOrders(orders []Order) []model.Order {
	var modelOrders []model.Order
	for _, order := range orders {
		var orderProducts []model.OrderProduct
		for _, op := range order.Products {
			orderProduct := model.OrderProduct{
				ProductID: op.ProductID,
				Quantity:  op.Quantity,
				Price:     op.Price,
			}
			orderProducts = append(orderProducts, orderProduct)
		}

		modelOrder := model.Order{
			ID:        order.ID,
			UserID:    order.UserID,
			Products:  orderProducts,
			Timestamp: order.Timestamp,
		}

		modelOrders = append(modelOrders, modelOrder)
	}
	return modelOrders
}

func (us *UserStorage) ToDomain() model.User {
	var UpdatedAt *time.Time
	if us.UpdateAt.Valid {
		UpdatedAt = pointer.Pointer(us.UpdateAt.Time)
	}

	modelOrders := convertOrders(us.Orders)

	return model.User{
		ID:        us.ID,
		FirstName: us.FirstName,
		LastName:  us.LastName,
		FullName:  us.FullName,
		Age:       us.Age,
		IsMarried: us.IsMarried,
		Password:  us.Password,
		CreatedAt: us.CreatedAt,
		UpdatedAt: UpdatedAt,
		Orders:    modelOrders,
	}
}
