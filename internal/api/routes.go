package api

import (
	"net/http"

	"github.com/FelipeBelloDultra/trunct.io/internal/api/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *API) BindRoutes() {
	api.Router.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
	)

	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("pong"))
			})

			r.Route("/accounts", func(r chi.Router) {
				r.Post("/", api.Controller.CreateAccount)
				r.Post("/session", api.Controller.AuthenticateAccount)
				r.Group(func(r chi.Router) {
					r.Use(middlewares.EnsureAuthenticated)
					r.Get("/me", api.Controller.ShowAuthenticatedAccount)
				})
			})

			r.Group(func(r chi.Router) {
				r.Use(middlewares.EnsureAuthenticated)
				r.Route("/urls", func(r chi.Router) {
					r.Post("/shorten", api.Controller.ShortenURL)
					r.Get("/", api.Controller.FetchURLs)
					r.Get("/{id}", api.Controller.ShowURLById)
					r.Get("/{id}/stats", api.Controller.ShowURLStatsById)
				})
			})
		})
	})

	api.Router.Get("/g/{url_code}", api.Controller.RedirectToURLByCode)
}
