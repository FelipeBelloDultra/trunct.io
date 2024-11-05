package controllers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type CreateAccountSchema struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c Controller) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var schema CreateAccountSchema
	if err := json.NewDecoder(r.Body).Decode(&schema); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request body"))
		return
	}

	id, err := c.AccountUseCase.CreateAccount(r.Context(), schema.Name, schema.Email, schema.Password)
	if err != nil {
		slog.Error("something went wrong", "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id.String(),
	})
}

func (c Controller) AuthenticateAccount(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("method not implemented"))
}

func (c Controller) ShowAuthenticatedAccount(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("method not implemented"))
}
