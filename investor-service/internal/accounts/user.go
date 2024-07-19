package accounts

import "github.com/google/uuid"

type User struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	persisted   bool
}

func NewUser(name string, email string, phoneNumber string) User {
	return User{
		Id:          uuid.New(),
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
	}
}

func (u *User) Persisted() bool {
	return u.persisted
}
