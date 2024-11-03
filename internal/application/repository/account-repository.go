package repository

import "github.com/FelipeBelloDultra/trunct.io/internal/application/entity"

type AccountRepository interface {
	Create(account *entity.Account) error
	FindByEmail(email string) (*entity.Account, error)
	FindByID(id string) (*entity.Account, error)
}
