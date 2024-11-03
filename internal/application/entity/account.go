package entity

import "github.com/google/uuid"

type Account struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
}

func NewAccount(name, email, passwordHash string) *Account {
	return &Account{
		ID:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
	}
}
