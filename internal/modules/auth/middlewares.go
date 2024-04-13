package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type contextKey string

const UserIdKey contextKey = "userID"

func Authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		words := strings.Split(authorization, " ")
		if len(words) != 2 || words[0] != "Bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := validateJwtToken(words[1])
		if err != nil || !token.Valid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userId, ok := claims["user_id"].(float64)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIdKey, uint(userId))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserIdFromContext(r *http.Request) uint {
	userId, ok := r.Context().Value(UserIdKey).(uint)
	if !ok {
		return 0
	}

	return userId
}
