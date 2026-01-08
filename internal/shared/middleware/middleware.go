package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
	"study-stack/internal/entities/tokens/stateless"
	"study-stack/internal/shared/consts"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := bearerToken(r)
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userData, err := stateless.VerifyAccessToken(token)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Println(err)
			return
		}

		ctx := context.WithValue(r.Context(), consts.UserDataContextKey, userData)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func bearerToken(r *http.Request) string {
	h := r.Header.Get("Authorization")
	const prefix = "Bearer "
	if !strings.HasPrefix(h, prefix) {
		return ""
	}
	return strings.TrimPrefix(h, prefix)
}
