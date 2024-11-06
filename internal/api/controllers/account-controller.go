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

type AuthenticateAccountSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c Controller) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var data httpvalidator.CreateAccountReqValidator
	validationErrors, err := c.decodeAndValidateJSON(r, &data)

	if err != nil {
		c.handleValidationError(w, err, "CreateAccount", validationErrors)
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

	c.sendJSONResponse(
		w,
		http.StatusCreated,
		map[string]any{
			"id": id,
		},
	)
}

func (c Controller) AuthenticateAccount(w http.ResponseWriter, r *http.Request) {
	var data httpvalidator.AuthenticateAccountReqValidator
	validationErrors, err := c.decodeAndValidateJSON(r, &data)

	if err != nil {
		c.handleValidationError(w, err, "AuthenticateAccount", validationErrors)
		return
	}

	token, err := c.AccountUseCase.AuthenticateAccount(r.Context(), data.Email, data.Password)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidCredentials) {
			c.handleError(w, http.StatusUnauthorized, "invalid credentials", nil)
			return
		}

		c.internalServerError(w, "AuthenticateAccount", err)
		return
	}

	c.sendJSONResponse(
		w,
		http.StatusOK,
		map[string]any{
			"token": token,
		},
	)
}

func (c Controller) ShowAuthenticatedAccount(w http.ResponseWriter, r *http.Request) {
	accountID := r.Context().Value(AccountIDKey("accountID")).(string)
	account, err := c.AccountUseCase.ShowAuthenticatedAccount(r.Context(), accountID)

	if err != nil {
		if errors.Is(err, usecase.ErrAccountNotFound) {
			c.handleError(w, http.StatusNotFound, "account not found", nil)
			return
		}

		if errors.Is(err, usecase.ErrAccountInvalidID) {
			c.handleError(w, http.StatusBadRequest, "invalid account ID", nil)
			return
		}

		c.internalServerError(w, "ShowAuthenticatedAccount", err)
		return
	}

	c.sendJSONResponse(
		w,
		http.StatusOK,
		map[string]any{
			"account": account,
		},
	)
}
