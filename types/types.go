package types

import "time"

type RegisterUserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct  {
	ID          uint   `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"-"` // omit password in JSON response
	CreatedAt   time.Time   `json:"createdAt"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user User) error
}