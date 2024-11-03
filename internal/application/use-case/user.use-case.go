package usecase

import (
	"errors"

	"github.com/FelipeBelloDultra/trunct.io/internal/application/entity"
	"github.com/FelipeBelloDultra/trunct.io/internal/application/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UserRepository repository.UserRepository
}

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrPasswordHashing   = errors.New("something went wrong with password hashing")
)

func (u *UserUseCase) CreateUser(name, email, password string) (*entity.User, error) {
	_, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, ErrUserAlreadyExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, ErrPasswordHashing
	}
	user := entity.NewUser(name, email, string(passwordHash))

	u.UserRepository.Create(user)

	return user, nil
}
