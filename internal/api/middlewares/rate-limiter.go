package middlewares

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/FelipeBelloDultra/trunct.io/internal/api/controllers"
	"github.com/go-chi/httprate"
)

func ApplyRateLimiting(requestLimit int, windowLength time.Duration) func(next http.Handler) http.Handler {
	return httprate.Limit(
		requestLimit,
		windowLength,
		httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint),
		httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(controllers.Response{
				StatusCode: http.StatusTooManyRequests,
				Error:      "number of attempts exceeded, please try again later",
			})
		}),
	)
}
