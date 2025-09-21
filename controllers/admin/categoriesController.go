package admin

import (
	"Store-Dio/models"
	"Store-Dio/services/categories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CategoriesController struct {
	CategoriesService *categories.CategoriesService
}

func NewCategoriesController(categoriesService *categories.CategoriesService) *CategoriesController {
	return &CategoriesController{
		CategoriesService: categoriesService,
	}
}
func (cc *CategoriesController) AddCategory(w http.ResponseWriter, r *http.Request) {
	var category *models.Category
	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	defer r.Body.Close()

	data, err := cc.CategoriesService.AddCategory(category)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message":  "Successfully",
		"Category": data,
	})
}
func (cc *CategoriesController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var category *models.Category

	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	defer r.Body.Close()
	err = cc.CategoriesService.UpdateCategory(category)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Successfully",
	})
}
func (cc *CategoriesController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}

	err = cc.CategoriesService.DeleteCategory(id)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Successfully",
	})
}
