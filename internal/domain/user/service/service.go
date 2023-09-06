package service

import (
	"context"
	"github.com/Amore14rn/888Starz/internal/domain/user/model"
	"github.com/Amore14rn/888Starz/pkg/errors"
	"time"
)

type repository interface {
	Create(ctx context.Context, req model.CreateUser) error
	CreateOrder(ctx context.Context, req model.CreateOrder) error
	AddToOrder(ctx context.Context, req model.AddToOrder) error
	GetOrderByUserID(ctx context.Context, userID string) (model.Order, error)
}

type UserService struct {
	repository repository
}

func NewUserService(repository repository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (u *UserService) CreateUser(ctx context.Context, req model.CreateUser) (model.User, error) {
	// Проверка возраста пользователя
	if req.Age < 18 {
		return model.User{}, errors.New("Пользователь должен быть не младше 18 лет")
	}

	// Проверка длины пароля
	if len(req.Password) < 8 {
		return model.User{}, errors.New("Пароль должен содержать не менее 8 символов")
	}

	// Проверка наличия цифр в пароле
	var isDigit bool
	for _, char := range req.Password {
		if char >= '0' && char <= '9' {
			isDigit = true
			break
		}
	}
	if !isDigit {
		return model.User{}, errors.New("Пароль должен содержать хотя бы одну цифру")
	}

	// Проверка наличия заглавных букв в пароле
	var isUpper bool
	for _, char := range req.Password {
		if char >= 'A' && char <= 'Z' {
			isUpper = true
			break
		}
	}
	if !isUpper {
		return model.User{}, errors.New("Пароль должен содержать хотя бы одну заглавную букву")
	}

	err := u.repository.Create(ctx, req)
	if err != nil {
		return model.User{}, err
	}
	return model.NewUser(
		req.ID,
		req.FirstName,
		req.LastName,
		req.Age,
		req.IsMarried,
		req.Password,
		req.Order,
		req.CreatedAt,
		nil), nil
}

func (u *UserService) CreateOrder(ctx context.Context, req model.CreateOrder) (model.Order, error) {
	err := u.repository.CreateOrder(ctx, req)
	if err != nil {
		return model.Order{}, err
	}
	return model.Order(model.NewOrder(
		req.ID,
		req.UserID,
		req.Products,
		req.Timestamp,
	)), nil
}

func (u *User) AddToOrder(orderToAdd Order) model.AddToOrder {
	// Create a new instance of model.AddToOrder based on the information in orderToAdd.
	addToOrder := model.AddToOrder{
		ID:        orderToAdd.ID,
		UserID:    orderToAdd.UserID,
		Products:  orderToAdd.Products,
		Timestamp: time.Now(), // You can set the current timestamp here or use the one from the orderToAdd.
	}

	// Add the newly created addToOrder to the User's Orders.
	u.Orders = append(u.Orders, orderToAdd)

	return addToOrder
}

func (u *UserService) fetchUserOrder(ctx context.Context, userID string) (model.Order, error) {
	// Assuming you have a repository method for fetching orders by UserID.
	order, err := u.repository.GetOrderByUserID(ctx, userID)
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}
