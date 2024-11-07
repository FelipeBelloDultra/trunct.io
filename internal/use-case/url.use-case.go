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

		_, err := uuc.queries.FindURLByCode(ctx, code)
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

func (uuc *URLUseCase) RedirectToURLByCode(ctx context.Context, code string) (string, error) {
	url, err := uuc.queries.FindURLByCode(ctx, code)
	if err == pgx.ErrNoRows {
		return "", ErrURLNotFound
	}

	if err != nil {
		return "", err
	}

	url.Clicks++
	_, err = uuc.queries.UpdateURL(
		ctx,
		pgstore.UpdateURLParams{
			ID:          url.ID,
			Clicks:      url.Clicks,
			OriginalUrl: url.OriginalUrl,
		},
	)
	if err != nil {
		return "", err
	}

	return url.OriginalUrl, nil
}

func (uuc *URLUseCase) FetchURLs(ctx context.Context, accountID string, limit, offset int32) ([]pgstore.FetchURLsByAccountIDRow, int64, error) {
	id, err := uuid.Parse(accountID)
	if err != nil {
		return nil, 0, ErrAccountInvalidID
	}

	rows, err := uuc.queries.FetchURLsByAccountID(
		ctx,
		pgstore.FetchURLsByAccountIDParams{
			OwnerID: id,
			Limit:   limit,
			Offset:  offset,
		},
	)
	if err != nil {
		return nil, 0, err
	}

	totalCount, err := uuc.queries.CountURLsByAccountID(ctx, id)
	if err != nil {
		return nil, 0, err
	}

	return rows, totalCount, nil
}
