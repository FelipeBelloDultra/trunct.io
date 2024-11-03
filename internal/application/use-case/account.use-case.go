package usecase

import (
	"errors"

	"github.com/FelipeBelloDultra/trunct.io/internal/application/entity"
	"github.com/FelipeBelloDultra/trunct.io/internal/application/repository"
	"golang.org/x/crypto/bcrypt"
)

type AccountUseCase struct {
	AccountRepository repository.AccountRepository
}

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrPasswordHashing    = errors.New("something went wrong with password hashing")
)

func (u *AccountUseCase) CreateAccount(name, email, password string) (*entity.Account, error) {
	_, err := u.AccountRepository.FindByEmail(email)
	if err != nil {
		return nil, ErrEmailAlreadyExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, ErrPasswordHashing
	}
	account := entity.NewAccount(name, email, string(passwordHash))

	u.AccountRepository.Create(account)

	return account, nil
}
