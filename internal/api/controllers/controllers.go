package controllers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	usecase "github.com/FelipeBelloDultra/trunct.io/internal/use-case"
	"github.com/FelipeBelloDultra/trunct.io/internal/validator"
)

type Response struct {
	StatusCode int                 `json:"status"`
	Data       interface{}         `json:"data,omitempty"`
	Error      string              `json:"error,omitempty"`
	Details    map[string][]string `json:"details,omitempty"`
}

type Controller struct {
	AccountUseCase usecase.AccountUseCase
}

type AccountIDKey string

var (
	ErrValidationFailed = errors.New("validation failed")
	ErrFailedDecodeJSON = errors.New("failed decoding JSON")
)

func (c *Controller) decodeAndValidateJSON(r *http.Request, data validator.Validator) (map[string][]string, error) {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return nil, ErrFailedDecodeJSON
	}

	problems := data.Valid(r.Context())
	if len(problems) > 0 {
		return problems, ErrValidationFailed
	}

	return nil, nil
}

func (c *Controller) encodeJSON(w http.ResponseWriter, statusCode int, res Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

func (c *Controller) handleError(w http.ResponseWriter, statusCode int, errMessage string, details map[string][]string) {
	response := Response{
		StatusCode: statusCode,
		Error:      errMessage,
	}
	if len(details) > 0 {
		response.Details = details
	}
	c.encodeJSON(w, statusCode, response)
}

func (c *Controller) internalServerError(w http.ResponseWriter, controllerName string, err error) {
	slog.Error("[CONTROLLER_ERROR]", "controller", controllerName, "error", err.Error())
	c.handleError(w, http.StatusInternalServerError, "internal server error", nil)
}

func (c *Controller) handleValidationError(w http.ResponseWriter, err error, controllerName string, validationErrors map[string][]string) {
	if validationErrors != nil && errors.Is(err, ErrValidationFailed) {
		c.handleError(w, http.StatusUnprocessableEntity, "validation failed", validationErrors)
		return
	}

	if errors.Is(err, ErrFailedDecodeJSON) {
		c.handleError(w, http.StatusBadRequest, "failed decoding JSON", nil)
		return
	}

	c.internalServerError(w, controllerName, err)
}

func (c *Controller) sendJSONResponse(w http.ResponseWriter, statusCode int, data map[string]any) {
	c.encodeJSON(
		w,
		statusCode,
		Response{
			StatusCode: statusCode,
			Data:       data,
		},
	)
}
