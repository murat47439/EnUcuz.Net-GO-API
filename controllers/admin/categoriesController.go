package admin

import (
	"Store-Dio/models"
	"Store-Dio/services/categories"
	"encoding/json"
	"net/http"
)

type CategoriesController struct {
	CategoriesService *categories.CategoriesService
}

func NewCategoriesController(categoriesService *categories.CategoriesService) *CategoriesController {
	return &CategoriesController{
		CategoriesService: categoriesService,
	}
}

func (cc *CategoriesController) InsertCategoriesData(w http.ResponseWriter, r *http.Request) {
	var cat models.Category
	err := json.NewDecoder(r.Body).Decode(&cat)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data %w"+err.Error())
		return
	}

	_, err = cc.CategoriesService.InsertCategoriesData(cat)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Action failed error : %w"+err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Successfully",
		"cat":     cat.Name,
	})
}
func (cc *CategoriesController) GetAllCategoriesID(w http.ResponseWriter, r *http.Request) {
	categories, err := cc.CategoriesService.GetAllCategoriesID()
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Successfully",
		"result":  categories,
	})
}
