package controllers

import (
	"errors"
	"net/http"
	"strconv"

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
	page := int32(1)
	limit := int32(10)

	if p := r.URL.Query().Get("page"); p != "" {
		if val, err := strconv.Atoi(p); err == nil {
			page = int32(val)

			if page <= 0 {
				c.handleError(w, http.StatusBadRequest, "page must be a positive integer", nil)
				return
			}
		}
	}

	if l := r.URL.Query().Get("limit"); l != "" {
		if val, err := strconv.Atoi(l); err == nil {
			limit = int32(val)

			if limit <= 0 {
				c.handleError(w, http.StatusBadRequest, "limit must be a positive integer", nil)
				return
			}
		}
	}

	accountID := r.Context().Value(AccountIDKey("accountID")).(string)
	offset := (page - 1) * limit
	urls, totalURL, err := c.URLUseCase.FetchURLs(r.Context(), accountID, limit, offset)
	if err != nil {
		c.internalServerError(w, "FetchURLs", err)
		return
	}

	totalPages := (totalURL + int64(limit) - 1) / int64(limit)
	var nextPage, prevPage int32
	if int64(page) < totalPages {
		nextPage = []int32{page + 1}[0]
	}
	if page > 1 {
		prevPage = []int32{page - 1}[0]
	}

	c.sendJSONResponse(
		w,
		http.StatusOK,
		map[string]any{
			"items": urls,
			"pagination": PaginationResponse{
				TotalCount:   int(totalURL),
				Limit:        int(limit),
				CurrentPage:  int(page),
				TotalPages:   int(totalPages),
				NextPage:     int(nextPage),
				PreviousPage: int(prevPage),
			},
		},
	)
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
