package controllers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	usecase "github.com/FelipeBelloDultra/trunct.io/internal/use-case"
	httpvalidator "github.com/FelipeBelloDultra/trunct.io/internal/validator/http"
)

func (c Controller) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var data httpvalidator.ShortenURLReqValidator
	validationErrors, err := c.decodeAndValidateJSON(r, &data)

	if err != nil {
		c.handleValidationError(w, err, "ShortenURL", validationErrors)
		return
	}

	accountID := r.Context().Value(AccountIDKey("accountID")).(string)
	code, err := c.URLUseCase.ShortenURL(r.Context(), data.OriginalURL, accountID)
	if err != nil {
		if errors.Is(err, usecase.ErrFailedToGetURLCode) {
			c.handleError(w, http.StatusBadRequest, "failed shortening URL", nil)
			return
		}

		if errors.Is(err, usecase.ErrAccountInvalidID) {
			c.handleError(w, http.StatusBadRequest, "invalid account ID", nil)
			return
		}

		c.internalServerError(w, "ShortenURL", err)
		return
	}

	c.sendJSONResponse(
		w,
		http.StatusCreated,
		map[string]any{
			"code": code,
		},
	)
}

func (c Controller) FetchURLs(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("method not implemented"))
}

func (c Controller) ShowURLById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("method not implemented"))
}

func (c Controller) ShowURLStatsById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("method not implemented"))
}

func (c Controller) RedirectToURLByCode(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "url_code")
	if code == "" {
		c.handleError(w, http.StatusBadRequest, "missing URL code", nil)
		return
	}

	url, err := c.URLUseCase.RedirectToURLByCode(r.Context(), code)
	if err != nil {
		if errors.Is(err, usecase.ErrURLNotFound) {
			c.handleError(w, http.StatusNotFound, "URL not found", nil)
			return
		}

		c.internalServerError(w, "RedirectToURLByCode", err)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}
