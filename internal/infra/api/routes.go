package api

import (
	"net/http"

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

			r.Route("/users", func(r chi.Router) {
				r.Post("/", api.Controller.CreateUser)
				r.Post("/session", api.Controller.AuthenticateUser)
				r.Group(func(r chi.Router) {
					// TODO: Add authentication middleware
					r.Get("/me", api.Controller.ShowAuthenticatedUser)
				})
			})

			r.Group(func(r chi.Router) {
				// TODO: Add authentication middleware
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
