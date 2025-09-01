package middleware

import (
	"Store-Dio/config"
	"Store-Dio/models"
	"Store-Dio/repo"
	"net/http"
	"strings"
	"time"
)

type UserMiddleware struct {
	UserRepo *repo.UserRepo
}

func NewUserMiddleware(ur *repo.UserRepo) *UserMiddleware {
	return &UserMiddleware{
		UserRepo: ur,
	}
}

func (um *UserMiddleware) OnlyAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization Header ", http.StatusUnauthorized)
			config.Logger.Printf("Missing or invalid Authorization Header")
			return
		}

		var token models.Token

		getJWT := strings.TrimPrefix(authHeader, "Bearer ")

		token.Token = getJWT

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
			tokenstring := r.Header.Get("Authorization")
			if tokenstring == "" {
				config.Logger.Printf("Auth middleware error")
				http.Error(w, "Missing token", http.StatusUnauthorized)
				return
			}
			token, err := um.UserRepo.DecodeJWT(models.Token{Token: tokenstring})

			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			if token.ExpiresAt.Time.Before(time.Now()) {
				http.Error(w, "Token expired", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		}
	})
}
