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
	CreatedAt time.Time
	UpdatedAt *time.Time
	Orders    []Order
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
	createdAt time.Time,
	updatedAt *time.Time,
	orders Order,
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
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Orders:    []Order{orders},
	}
}

type CreateOrder struct {
	ID        string
	UserID    string
	Products  []OrderProduct
	Timestamp time.Time
}

func NewCreateOrder(
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

type CreateUser struct {
	ID        string
	FirstName string
	LastName  string
	FullName  string
	Age       uint32
	IsMarried bool
	Password  string
	CreatedAt time.Time
}

func NewCreateUser(
	ID string,
	firstName string,
	lastName string,
	age uint32,
	isMarried bool,
	password string,
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
		Password:  password,
		CreatedAt: createdAt,
	}
}
