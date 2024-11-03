package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/FelipeBelloDultra/trunct.io/internal/infra/api"
	"github.com/FelipeBelloDultra/trunct.io/internal/infra/api/controllers"
	"github.com/go-chi/chi/v5"
)

func main() {
	api := api.API{
		Router:     chi.NewMux(),
		Controller: controllers.NewController(),
	}
	api.BindRoutes()

	fmt.Printf("listening on port :%d\n", 3333)
	if err := http.ListenAndServe("localhost:3333", api.Router); err != nil {
		slog.Error("failed to start server", "err", err.Error())
		panic(err)
	}
}
