package user

import (
	"context"
	"github.com/Amore14rn/888Starz/internal/domain/user/model"
	"github.com/Amore14rn/888Starz/pkg/errors"
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
	// Check user's age
	if req.Age < 18 {
		return model.User{}, errors.New("User must be at least 18 years old")
	}

	// Check password length
	if len(req.Password) < 8 {
		return model.User{}, errors.New("Password must be at least 8 characters long")
	}

	// Check for the presence of digits in the password
	var hasDigit bool
	for _, char := range req.Password {
		if char >= '0' && char <= '9' {
			hasDigit = true
			break
		}
	}
	if !hasDigit {
		return model.User{}, errors.New("Password must contain at least one digit")
	}

	// Check for the presence of uppercase letters in the password
	var hasUpper bool
	for _, char := range req.Password {
		if char >= 'A' && char <= 'Z' {
			hasUpper = true
			break
		}
	}
	if !hasUpper {
		return model.User{}, errors.New("Password must contain at least one uppercase letter")
	}

	err := u.repository.Create(ctx, req)
	if err != nil {
		return model.User{}, err
	}
	return model.User{
		ID:        req.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		FullName:  req.FullName,
		Age:       req.Age,
		IsMarried: req.IsMarried,
		Password:  req.Password,
		CreatedAt: req.CreatedAt,
	}, nil
}

func (u *UserService) CreateOrder(ctx context.Context, req model.CreateOrder) (model.Order, error) {
	err := u.repository.CreateOrder(ctx, req)
	if err != nil {
		return model.Order{}, err
	}
	return model.Order{
		ID:        req.ID,
		UserID:    req.UserID,
		Products:  req.Products,
		Timestamp: req.Timestamp,
	}, nil
}

func (u *UserService) AddToOrder(ctx context.Context, req model.AddToOrder) (model.Order, error) {
	err := u.repository.AddToOrder(ctx, req)
	if err != nil {
		return model.Order{}, err
	}
	return model.Order{
		ID:        req.ID,
		UserID:    req.UserID,
		Products:  req.Products,
		Timestamp: req.Timestamp,
	}, nil
}

func (u *UserService) GetOrderByUserID(ctx context.Context, userID string) (model.Order, error) {
	order, err := u.repository.GetOrderByUserID(ctx, userID)
	if err != nil {
		return model.Order{}, err
	}
	return model.Order{
		ID:        order.ID,
		UserID:    order.UserID,
		Products:  order.Products,
		Timestamp: order.Timestamp,
	}, nil
}
