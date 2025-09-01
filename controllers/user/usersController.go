package user

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

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Data")
		return
	}

	accessToken, refreshToken, err := uc.UserService.Login(user)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Login error"+err.Error())
		return
	}

	RespondWithJSON(w, http.StatusAccepted, map[string]string{
		"token":         accessToken,
		"refresh-token": refreshToken,
		"message":       "Succesfully",
	})
}
func (uc *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	var token models.RefreshToken

	err := json.NewDecoder(r.Body).Decode(&token)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid token")
		return
	}
	_, err = uc.UserService.Logout(token)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusAccepted, map[string]string{
		"message": "Successfully",
	})
}
func (uc *UserController) GetUserData(w http.ResponseWriter, r *http.Request) {
	var token models.Token

	err := json.NewDecoder(r.Body).Decode(&token)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	user, err := uc.UserService.GetUserDataByID(token)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Service Error : "+err.Error())
		return
	}
	RespondWithJSON(w, http.StatusAccepted, map[string]interface{}{
		"message": "successfully",
		"Name":    user.Name,
		"Surname": user.Surname,
		"Email":   user.Email,
		"Phone":   user.Phone,
		"Gender":  user.Gender,
		"Role":    user.Role,
	})

}
func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
	var claims models.UpdateUser
	err := json.NewDecoder(r.Body).Decode(&claims)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}

	if claims.User.ID != 0 || claims.Token.Token == "" {
		RespondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	token, err := uc.UserService.UserRepo.DecodeJWT(claims.Token)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	claims.User.ID = token.UserID

	_, err = uc.UserService.Update(claims.User)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Succesfully",
		"name":    claims.User.Name,
		"surname": claims.User.Surname,
		"email":   claims.User.Email,
		"phone":   claims.User.Phone,
		"gender":  claims.User.Gender,
	})

}
