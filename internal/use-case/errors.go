package usecase

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrPasswordHashing    = errors.New("something went wrong with password hashing")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrAccountInvalidID   = errors.New("invalid account id")
	ErrAccountNotFound    = errors.New("account not found")
	ErrFailedToGetURLCode = errors.New("failed to get URL code")
)
