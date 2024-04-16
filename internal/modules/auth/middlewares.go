package auth

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type contextKey string

const UserIdKey contextKey = "userID"

func Authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookieToken, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := validateJwtToken(cookieToken.Value)
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

func AuthorizedRedirect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookieToken, err := r.Cookie("token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		token, err := validateJwtToken(cookieToken.Value)
		if err != nil || !token.Valid {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userId, ok := claims["user_id"].(float64)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
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
