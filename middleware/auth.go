package middleware

import (
	"Store-Dio/config"
	"Store-Dio/models"
	"Store-Dio/repo"
	"context"
	"net/http"
	"time"
)

type UserMiddleware struct {
	UserRepo *repo.UserRepo
}
type contextKey string

const UserIDKey contextKey = "userID"

func NewUserMiddleware(ur *repo.UserRepo) *UserMiddleware {
	return &UserMiddleware{
		UserRepo: ur,
	}
}

func (um *UserMiddleware) OnlyAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("access_token")

		if err != nil {
			http.Error(w, "No data", http.StatusNotFound)
			return
		}
		tokenString := cookie.Value

		var token models.Token

		token.Token = tokenString

		jwtToken, err := um.UserRepo.DecodeJWT(token)

		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		isAdmin, err := um.UserRepo.OnlyAdmin(jwtToken.UserID)

		if err != nil {
			config.Logger.Printf("DB error checking admin role: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if !isAdmin {
			http.Error(w, "Forbidden: Admins only", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func (um *UserMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		{
			cookie, err := r.Cookie("access_token")

			if err != nil {
				http.Error(w, "No data", http.StatusNotFound)
				return
			}
			tokenString := cookie.Value

			token, err := um.UserRepo.DecodeJWT(models.Token{Token: tokenString})

			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			if token.ExpiresAt.Time.Before(time.Now()) {
				http.Error(w, "Token expired", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), UserIDKey, token.UserID)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
