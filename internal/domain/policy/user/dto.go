package user

import "time"

type CreateUserInput struct {
	ID        string
	FirstName string
	LastName  string
	FullName  string
	Age       uint32
	IsMarried bool
	Password  string
	CreatedAt time.Time
}
