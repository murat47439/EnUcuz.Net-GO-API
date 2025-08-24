package middleware

import (
	"Store-Dio/repo"
	"net/http"
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
	return nil
}
