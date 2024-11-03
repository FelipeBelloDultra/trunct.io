package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/FelipeBelloDultra/trunct.io/internal/infra/api"
	"github.com/FelipeBelloDultra/trunct.io/internal/infra/api/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("LOAD_ENV_FILE") == "true" {
		if err := godotenv.Load(); err != nil {
			slog.Error("error loading.env file", "err", err.Error())
			panic(err)
		}
	}

	api := api.API{
		Router:     chi.NewMux(),
		Controller: controllers.NewController(),
	}
	api.BindRoutes()

	fmt.Printf("listening on port :%d\n", 3333)
	if err := http.ListenAndServe(":3333", api.Router); err != nil {
		slog.Error("failed to start server", "err", err.Error())
		panic(err)
	}
}
