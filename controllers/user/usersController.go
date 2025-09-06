package user

import (
	"Store-Dio/config"
	"Store-Dio/models"
	"Store-Dio/services/users"
	"encoding/json"
	"net/http"
	"time"
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
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/api",
		Expires:  time.Now().Add(15 * time.Minute),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/api/refresh",
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	RespondWithJSON(w, http.StatusAccepted, map[string]string{
		"message": "Succesfully",
	})
}
func (uc *UserController) Logout(w http.ResponseWriter, r *http.Request) {

	userID, _, ok := GetUserIDFromContext(r)

	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorizated")
		return
	}

	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, "Missing refresh token", http.StatusUnauthorized)
		return
	}
	token := cookie.Value

	_, err = uc.UserService.Logout(token, userID)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Succesfully",
	})

}
func (uc *UserController) GetUserData(w http.ResponseWriter, r *http.Request) {
	userID, _, ok := GetUserIDFromContext(r)

	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorizated")
		return
	}

	user, err := uc.UserService.GetUserDataByID(userID)

	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"Message": "Successfully",
		"User":    user,
	})
}
func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
	userID, role, ok := GetUserIDFromContext(r)

	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorizated")
		return
	}
	var data *models.User

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	if role != 1 {
		if userID != data.ID {
			RespondWithError(w, http.StatusBadRequest, "Error")
			return
		}
	}

	user, err := uc.UserService.Update(data)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"Message": "Successfully",
		"User": map[string]interface{}{
			"id":      user.ID,
			"name":    user.Name,
			"surname": user.Surname,
			"email":   user.Email,
			"phone":   user.Phone,
		},
	})
}
