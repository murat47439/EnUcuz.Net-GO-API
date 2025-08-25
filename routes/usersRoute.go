package routes

import (
	"Store-Dio/controllers"
	"net/http"
)

type UserRoutes struct {
	UserController *controllers.UserController
}

func NewUserRoute(uc *controllers.UserController) *UserRoutes {
	return &UserRoutes{
		UserController: uc,
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
	mux.HandleFunc("/profile/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ur.UserController.GetUserData(w, r)

		default:
			RespondWithError(w, http.StatusBadRequest, "Method not allowed")
		}
	})
	return mux
}
