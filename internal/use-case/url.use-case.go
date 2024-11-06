package usecase

import (
	"context"
	"math/rand/v2"

	"github.com/FelipeBelloDultra/trunct.io/internal/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type URLUseCase struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewURLUseCase(pool *pgxpool.Pool) URLUseCase {
	return URLUseCase{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321"

func (uuc *URLUseCase) genCode() string {
	const n = 8
	byts := make([]byte, 8)
	for i := range n {
		byts[i] = characters[rand.IntN(len(characters))]
	}
	return string(byts)
}

func (uuc *URLUseCase) ShortenURL(ctx context.Context, originalURL, ownerID string) (string, error) {
	id, err := uuid.Parse(ownerID)
	if err != nil {
		return "", ErrAccountInvalidID
	}

	const maxRetries = 5
	var code string
	for i := range maxRetries {
		code = uuc.genCode()
		currentTry := i + 1

		_, err := uuc.queries.FindByCode(ctx, code)
		if err == pgx.ErrNoRows {
			break
		}

		if currentTry == maxRetries {
			return "", ErrFailedToGetURLCode
		}
	}

	if code == "" {
		return "", ErrFailedToGetURLCode
	}

	_, err = uuc.queries.CreateURL(
		ctx,
		pgstore.CreateURLParams{
			OriginalUrl: originalURL,
			Code:        code,
			OwnerID:     id,
		},
	)
	if err != nil {
		return "", err
	}

	return code, nil
}
