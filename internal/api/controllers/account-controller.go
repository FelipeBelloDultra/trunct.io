package controllers

import (
	"errors"
	"net/http"

	usecase "github.com/FelipeBelloDultra/trunct.io/internal/use-case"
	httpvalidator "github.com/FelipeBelloDultra/trunct.io/internal/validator/http"
)

type CreateAccountSchema struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c Controller) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var data httpvalidator.CreateUserReqValidator
	validationErrors, err := c.decodeAndValidateJSON(r, &data)

	if err != nil {
		if validationErrors != nil && errors.Is(err, ErrValidationFailed) {
			c.handleError(w, http.StatusUnprocessableEntity, "validation failed", validationErrors)
			return
		}

		if errors.Is(err, ErrFailedDecodeJSON) {
			c.handleError(w, http.StatusBadRequest, "failed decoding JSON", nil)
			return
		}

		c.internalServerError(w, "CreateAccount", err)
		return
	}

	id, err := c.AccountUseCase.CreateAccount(r.Context(), data.Name, data.Email, data.Password)

	if err != nil {
		if errors.Is(err, usecase.ErrEmailAlreadyExists) {
			c.handleError(w, http.StatusConflict, "email already exists", nil)
			return
		}

		c.internalServerError(w, "CreateAccount", err)
		return
	}

	c.encodeJSON(
		w,
		http.StatusCreated,
		Response{
			StatusCode: http.StatusCreated,
			Data: map[string]any{
				"id": id,
			},
		},
	)
}

func (c Controller) AuthenticateAccount(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("method not implemented"))
}

func (c Controller) ShowAuthenticatedAccount(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("method not implemented"))
}
