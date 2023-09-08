package user

import (
	"context"
	"github.com/Amore14rn/888Starz/internal/domain/user/model"
	"github.com/Amore14rn/888Starz/internal/domain/user/service"
	"github.com/Amore14rn/888Starz/pkg/common/core/clock"
	"github.com/Amore14rn/888Starz/pkg/errors"
	"time"
)

type IdentityGenerator interface {
	GenerateUUIDv4String() string
}

type Clock interface {
	Now() time.Time
}

type Policy struct {
	userService *service.UserService

	identity IdentityGenerator
	clock    Clock
}

func NewUserPolicy(userService *service.UserService, identity IdentityGenerator, clock clock.Clock) *Policy {
	return &Policy{
		userService: userService,
		identity:    identity,
		clock:       clock,
	}
}

func (u *Policy) CreateUser(ctx context.Context, input CreateUserInput) (CreateUserOutput, error) {
	// Check user's age
	if input.Age < 18 {
		return CreateUserOutput{}, errors.New("User must be at least 18 years old")
	}

	// Check password length
	if len(input.Password) < 8 {
		return CreateUserOutput{}, errors.New("Password must be at least 8 characters long")
	}

	// Check for the presence of digits in the password
	var hasDigit bool
	for _, char := range input.Password {
		if char >= '0' && char <= '9' {
			hasDigit = true
			break
		}
	}
	if !hasDigit {
		return CreateUserOutput{}, errors.New("Password must contain at least one digit")
	}

	// Check for the presence of uppercase letters in the password
	var hasUpper bool
	for _, char := range input.Password {
		if char >= 'A' && char <= 'Z' {
			hasUpper = true
			break
		}
	}
	if !hasUpper {
		return CreateUserOutput{}, errors.New("Password must contain at least one uppercase letter")
	}

	createUser := model.NewCreateUser(
		u.identity.GenerateUUIDv4String(),
		input.FirstName,
		input.LastName,
		input.Age,
		input.IsMarried,
		input.Password,
		input.Order,
		u.clock.Now(),
	)

	user, err := u.userService.CreateUser(ctx, createUser)
	if err != nil {
		return CreateUserOutput{}, errors.Wrap(err, "userService.CreateUser")
	}

	return CreateUserOutput{
		User: user,
	}, nil
}

//func (u *Policy) CreateOrder(ctx context.Context, req model.CreateOrder) (model.Order, error) {
//	err := u.repository.CreateOrder(ctx, req)
//	if err != nil {
//		return model.Order{}, err
//	}
//	return model.Order{
//		ID:        req.ID,
//		UserID:    req.UserID,
//		Products:  req.Products,
//		Timestamp: req.Timestamp,
//	}, nil
//}
//
//func (u *Policy) AddToOrder(ctx context.Context, req model.AddToOrder) (model.Order, error) {
//	err := u.repository.AddToOrder(ctx, req)
//	if err != nil {
//		return model.Order{}, err
//	}
//	return model.Order{
//		ID:        req.ID,
//		UserID:    req.UserID,
//		Products:  req.Products,
//		Timestamp: req.Timestamp,
//	}, nil
//}
//
//func (u *Policy) GetOrderByUserID(ctx context.Context, userID string) (model.Order, error) {
//	order, err := u.userService.GetOrderByUserID(ctx, userID)
//	if err != nil {
//		return model.Order{}, err
//	}
//	return model.Order{
//		ID:        order.ID,
//		UserID:    order.UserID,
//		Products:  order.Products,
//		Timestamp: order.Timestamp,
//	}, nil
//}
