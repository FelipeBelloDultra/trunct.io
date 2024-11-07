package api

import (
	"net/http"
	"time"

	"github.com/FelipeBelloDultra/trunct.io/internal/api/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *API) BindRoutes() {
	api.Router.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		middlewares.ApplyRateLimiting(10, time.Minute),
	)

	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("pong"))
			})

			r.Route("/accounts", func(r chi.Router) {
				r.Use(middlewares.ApplyRateLimiting(5, time.Minute))
				r.Post("/", api.Controller.CreateAccount)
				r.Post("/session", api.Controller.AuthenticateAccount)

				r.Group(func(r chi.Router) {
					r.Use(middlewares.EnsureAuthenticated)
					r.Get("/me", api.Controller.ShowAuthenticatedAccount)
				})
			})

			r.Group(func(r chi.Router) {
				r.Use(
					middlewares.EnsureAuthenticated,
					middlewares.ApplyRateLimiting(20, time.Minute),
				)
				r.Route("/urls", func(r chi.Router) {
					r.Post("/shorten", api.Controller.ShortenURL)
					r.Get("/", api.Controller.FetchURLs)
				})
			})
		})
	})

	api.Router.Group(func(r chi.Router) {
		r.Use(middlewares.ApplyRateLimiting(50, time.Minute))
		api.Router.Get("/g/{url_code}", api.Controller.RedirectToURLByCode)
	})
}
