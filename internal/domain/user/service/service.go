package service

import (
	"context"
	"github.com/Amore14rn/888Starz/internal/domain/user/model"
	"github.com/Amore14rn/888Starz/pkg/errors"
)

type repository interface {
	Create(ctx context.Context, req model.CreateUser) error
	//CreateOrder(ctx context.Context, req model.CreateOrder) error
	//AddToOrder(ctx context.Context, req model.AddToOrder) error
	//GetOrderByUserID(ctx context.Context, userID string) (model.Order, error)
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

//func (u *UserService) CreateOrder(ctx context.Context, req model.CreateOrder) (model.Order, error) {
//	err := u.repository.CreateOrder(ctx, req)
//	if err != nil {
//		return model.Order{}, err
//	}
//	return model.Order(model.NewOrder(
//		req.ID,
//		req.UserID,
//		req.Products,
//		req.Timestamp,
//	)), nil
//}
//
//func (u *UserService) AddToOrder(ctx context.Context, req model.AddToOrder) (model.Order, error) {
//	err := u.repository.AddToOrder(ctx, req)
//	if err != nil {
//		return model.Order{}, err
//	}
//	return model.Order(model.NewAddOrder(
//		req.ID,
//		req.UserID,
//		req.Products,
//		req.Timestamp,
//	)), nil
//}
