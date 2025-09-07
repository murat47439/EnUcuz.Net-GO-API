package user

import (
	"Store-Dio/middleware"
	"Store-Dio/models"
	"Store-Dio/services/favories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type FavoriesController struct {
	FavoriesServices *favories.FavoriesService
}

func NewFavoriesController(service *favories.FavoriesService) *FavoriesController {
	return &FavoriesController{
		FavoriesServices: service,
	}
}
func GetUserIDFromContext(r *http.Request) (int, int, bool) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	userRole, ok := r.Context().Value(middleware.UserRole).(int)

	return userID, userRole, ok
}
func (fc *FavoriesController) AddFavori(w http.ResponseWriter, r *http.Request) {
	userID, _, ok := GetUserIDFromContext(r)
	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	var product *models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	err = fc.FavoriesServices.AddFavori(product, userID)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Successfully",
	})

}
func (fc *FavoriesController) RemoveFavori(w http.ResponseWriter, r *http.Request) {
	userID, _, ok := GetUserIDFromContext(r)

	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	fav, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		RespondWithError(w, http.StatusNotFound, "Invalid data")
		return
	}

	err = fc.FavoriesServices.RemoveFavori(fav, userID)

	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Successfully",
	})

}
func (fc *FavoriesController) GetFavourites(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	userID, _, ok := GetUserIDFromContext(r)
	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		page = 1
	}

	data, err := fc.FavoriesServices.GetFavourites(page, userID)

	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message":    "Successfully",
		"Favourites": data,
	})

}
