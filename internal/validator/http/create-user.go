package httpvalidator

import (
	"context"

	"github.com/FelipeBelloDultra/trunct.io/internal/validator"
)

type CreateUserReqValidator struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req CreateUserReqValidator) Valid(context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(
		validator.NotBlank(req.Name),
		"name",
		"name is required",
	)
	eval.CheckField(
		validator.MinChars(req.Name, 4),
		"name",
		"name must have at least 4 characters",
	)
	eval.CheckField(
		validator.MaxChars(req.Name, 100),
		"name",
		"name must have a maximum of 100 characters",
	)
	eval.CheckField(
		validator.Matches(req.Email, validator.EmailRegex),
		"email",
		"invalid email format",
	)
	eval.CheckField(
		validator.MaxChars(req.Email, 100),
		"email",
		"email must have a maximum of 100 characters",
	)
	eval.CheckField(
		validator.MinChars(req.Password, 8),
		"password",
		"password must have at least 8 characters",
	)
	eval.CheckField(
		validator.MaxChars(req.Password, 100),
		"password",
		"password must have a maximum of 50 characters",
	)

	return eval
}
