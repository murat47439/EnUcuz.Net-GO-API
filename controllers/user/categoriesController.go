package user

import (
	"Store-Dio/services/categories"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UCategoriesController struct {
	CategoriesService *categories.CategoriesService
}

func NewUCategoriesController(service *categories.CategoriesService) *UCategoriesController {
	return &UCategoriesController{
		CategoriesService: service,
	}
}
func (uc *UCategoriesController) GetCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	category, err := uc.CategoriesService.GetCategory(id)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"Category": category,
	})
}
func (uc *UCategoriesController) GetCategories(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		page = 1
	}
	search := query.Get("search")
	if search == "undefined" {
		search = ""
	}
	categories, err := uc.CategoriesService.GetCategories(page, search)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"Categories": categories,
	})
}
