package model

import (
	"time"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	FullName  string
	Age       uint32
	IsMarried bool
	Password  string
	Orders    []Order
	CreatedAt time.Time
	UpdatedAt *time.Time
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

func (u *User) AddOrder(order Order) {
	u.Orders = append(u.Orders, order)
}

func NewUser(
	ID string,
	firstName string,
	lastName string,
	age uint32,
	isMarried bool,
	password string,
	orders Order,
	createdAt time.Time,
	updatedAt *time.Time,
) User {
	fullName := firstName + " " + lastName
	return User{
		ID:        ID,
		FirstName: firstName,
		LastName:  lastName,
		FullName:  fullName,
		Age:       age,
		IsMarried: isMarried,
		Password:  password,
		Orders:    []Order{orders},
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type CreateUser struct {
	ID        string
	FirstName string
	LastName  string
	FullName  string
	Age       uint32
	IsMarried bool
	Password  string
	Order     Order
	CreatedAt time.Time
}

func NewCreateUser(
	ID string,
	firstName string,
	lastName string,
	age uint32,
	isMarried bool,
	password string,
	order Order,
	createdAt time.Time,
) CreateUser {
	fullName := firstName + " " + lastName
	return CreateUser{
		ID:        ID,
		FirstName: firstName,
		LastName:  lastName,
		FullName:  fullName,
		Age:       age,
		IsMarried: isMarried,
		Order:     order,
		Password:  password,
		CreatedAt: createdAt,
	}
}

type CreateOrder struct {
	ID        string
	UserID    string
	Products  []OrderProduct
	Timestamp time.Time
}

func NewOrder(
	ID string,
	userID string,
	products []OrderProduct,
	timestamp time.Time,
) CreateOrder {
	return CreateOrder{
		ID:        ID,
		UserID:    userID,
		Products:  products,
		Timestamp: timestamp,
	}
}

type AddToOrder struct {
	ID        string
	UserID    string
	Products  []OrderProduct
	Timestamp time.Time
}

func NewAddOrder(
	ID string,
	userID string,
	products []OrderProduct,
	timestamp time.Time,
) AddToOrder {
	return AddToOrder{
		ID:        ID,
		UserID:    userID,
		Products:  products,
		Timestamp: timestamp,
	}
}
