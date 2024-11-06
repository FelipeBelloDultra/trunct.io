package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/FelipeBelloDultra/trunct.io/internal/api/controllers"
	"github.com/FelipeBelloDultra/trunct.io/internal/jwt"
)

func EnsureAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			sendUnauthorizedJSON(w, "authorization header is missing")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.VerifyToken(tokenString)
		if err != nil {
			sendUnauthorizedJSON(w, "invalid token")
			return
		}

		subject, err := token.Claims.GetSubject()
		if err != nil {
			sendUnauthorizedJSON(w, "invalid token")
			return
		}

		accountID := controllers.AccountIDKey("accountID")
		ctx := context.WithValue(r.Context(), accountID, subject)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func sendUnauthorizedJSON(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(
		controllers.Response{
			StatusCode: http.StatusUnauthorized,
			Error:      message,
		},
	)
}
