package routes

import (
	"Store-Dio/controllers"
	"Store-Dio/middleware"
	"net/http"
)

type UserRoutes struct {
	UserController *controllers.UserController
	UserMiddleware *middleware.UserMiddleware
}

func NewUserRoute(uc *controllers.UserController, um *middleware.UserMiddleware) *UserRoutes {
	return &UserRoutes{
		UserController: uc,
		UserMiddleware: um,
	}
}

func (ur *UserRoutes) SetupUsersRoute() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			ur.UserController.CreateUser(w, r)
		default:
			RespondWithError(w, 100, "Method Not Allowed")
			return
		}
	})
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			ur.UserController.Login(w, r)
		default:
			RespondWithError(w, http.StatusBadRequest, "Method not allowed")
			return
		}
	})
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ur.UserController.GetUserData(w, r)

		default:
			RespondWithError(w, http.StatusBadRequest, "Method not allowed")
		}
	})
	mux.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodPut:
			ur.UserController.Update(w, r)
		default:
			RespondWithError(w, http.StatusBadRequest, "Method not allowed")
		}

	})
	mux.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			ur.UserController.Logout(w, r)
		default:
			RespondWithError(w, http.StatusBadRequest, "Method not allowed")
		}
	})
	mux.Handle("/admin", ur.UserMiddleware.OnlyAdmin(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ur.UserController.GetUserData(w, r)

		default:
			RespondWithError(w, http.StatusBadRequest, "Method not allowed")
		}
	})))
	return mux
}
