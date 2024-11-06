package httpvalidator

import (
	"context"

	"github.com/FelipeBelloDultra/trunct.io/internal/validator"
)

type ShortenURLReqValidator struct {
	OriginalURL string `json:"original_url"`
}

func (s ShortenURLReqValidator) Valid(context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(
		validator.IsURL(s.OriginalURL),
		"original_url",
		"invalid URL format",
	)

	return eval
}
