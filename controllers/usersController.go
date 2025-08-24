package controllers

import (
	"Store-Dio/config"
	"Store-Dio/models"
	"Store-Dio/services/users"
	"encoding/json"
	"net/http"
)

type UserController struct {
	UserService *users.UserService
}

func NewUserController(us *users.UserService) *UserController {
	return &UserController{
		UserService: us,
	}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		config.Logger.Printf("CreateUser decode error: %v", err)
		RespondWithJSON(w, http.StatusBadRequest, map[string]string{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	_, err = uc.UserService.CreateUser(user)

	if err != nil {
		config.Logger.Printf("Failed to create user:  %v", err)
		RespondWithJSON(w, http.StatusBadRequest, map[string]string{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Succesfully",
	})
}

func (uc *UserController) Login(w http.ResponseWriter, r http.Request) {

}
