package api

import (
	"github.com/FelipeBelloDultra/trunct.io/internal/api/controllers"
	"github.com/go-chi/chi/v5"
)

type API struct {
	Router     *chi.Mux
	Controller controllers.Controller
}
