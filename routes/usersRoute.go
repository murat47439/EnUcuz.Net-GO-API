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

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			ur.UserController.CreateUser(w, r)
		default:
			RespondWithError(w, 100, "Method Not Allowed")
			return
		}
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		default:
			RespondWithError(w, 100, "Method Not Allowed")
			return
		}

	})
	return mux
}
