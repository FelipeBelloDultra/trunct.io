package repository

import "github.com/FelipeBelloDultra/trunct.io/internal/application/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id string) (*entity.User, error)
}
