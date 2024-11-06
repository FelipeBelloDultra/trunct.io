package usecase

import (
	"context"
	"errors"

	"github.com/FelipeBelloDultra/trunct.io/internal/jwt"
	"github.com/FelipeBelloDultra/trunct.io/internal/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type AccountUseCase struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

type AuthenticatedAccount struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func NewAccountUseCase(pool *pgxpool.Pool) AccountUseCase {
	return AccountUseCase{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrPasswordHashing    = errors.New("something went wrong with password hashing")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrAccountInvalidID   = errors.New("invalid account id")
	ErrAccountNotFound    = errors.New("account not found")
)

func (aus *AccountUseCase) CreateAccount(ctx context.Context, name, email, password string) (uuid.UUID, error) {
	_, err := aus.queries.FindAccountByEmail(ctx, email)

	if err != nil && err != pgx.ErrNoRows {
		return uuid.UUID{}, err
	}

	if err == nil {
		return uuid.UUID{}, ErrEmailAlreadyExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return uuid.UUID{}, ErrPasswordHashing
	}

	id, err := aus.queries.CreateAccount(ctx,
		pgstore.CreateAccountParams{
			Name:         name,
			Email:        email,
			PasswordHash: []byte(passwordHash),
		},
	)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (aus *AccountUseCase) AuthenticateAccount(ctx context.Context, email, password string) (string, error) {
	account, err := aus.queries.FindAccountByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", ErrInvalidCredentials
		}

		return "", err
	}

	if err := bcrypt.CompareHashAndPassword(account.PasswordHash, []byte(password)); err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := jwt.CreateTokenFromID(account.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (aus *AccountUseCase) ShowAuthenticatedAccount(ctx context.Context, id string) (AuthenticatedAccount, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return AuthenticatedAccount{}, ErrAccountInvalidID
	}

	account, err := aus.queries.FindAccountByID(ctx, uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return AuthenticatedAccount{}, ErrAccountNotFound
		}

		return AuthenticatedAccount{}, err
	}

	return AuthenticatedAccount{
		ID:    account.ID,
		Name:  account.Name,
		Email: account.Email,
	}, nil
}
