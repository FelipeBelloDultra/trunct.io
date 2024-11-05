package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/FelipeBelloDultra/trunct.io/internal/api"
	"github.com/FelipeBelloDultra/trunct.io/internal/api/controllers"
	usecase "github.com/FelipeBelloDultra/trunct.io/internal/use-case"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("LOAD_ENV_FILE") == "true" {
		if err := godotenv.Load(); err != nil {
			slog.Error("error loading.env file", "err", err.Error())
			panic(err)
		}
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("TRUNCT_DATABASE_USER"),
		os.Getenv("TRUNCT_DATABASE_PASSWORD"),
		os.Getenv("TRUNCT_DATABASE_HOST"),
		os.Getenv("TRUNCT_DATABASE_PORT"),
		os.Getenv("TRUNCT_DATABASE_NAME"),
	))
	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	api := api.API{
		Router: chi.NewMux(),
		Controller: controllers.Controller{
			AccountUseCase: usecase.NewAccountUseCase(pool),
		},
	}
	api.BindRoutes()

	fmt.Printf("listening on port :%d\n", 3333)
	if err := http.ListenAndServe(":3333", api.Router); err != nil {
		slog.Error("failed to start server", "err", err.Error())
		panic(err)
	}
}
