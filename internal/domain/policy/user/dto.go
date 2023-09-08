package user

import (
	"github.com/Amore14rn/888Starz/internal/domain/user/model"
	"time"
)

type CreateUserInput struct {
	ID        string
	FirstName string
	LastName  string
	FullName  string
	Age       uint32
	IsMarried bool
	Password  string
	CreatedAt time.Time
	Order     model.Order
}

func NewCreateUserInput(firstName string, lastName string, fullname string, age uint32, isMarried bool, password string, order model.Order) CreateUserInput {
	return CreateUserInput{
		FirstName: firstName,
		LastName:  lastName,
		FullName:  fullname,
		Age:       age,
		IsMarried: isMarried,
		Password:  password,
		Order:     order,
	}
}

type CreateUserOutput struct {
	User model.User
}
