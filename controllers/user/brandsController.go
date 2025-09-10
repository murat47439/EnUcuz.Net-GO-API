package user

import (
	"Store-Dio/services/brands"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UBrandController struct {
	BrandsService *brands.BrandsService
}

func NewUBrandController(service *brands.BrandsService) *UBrandController {
	return &UBrandController{
		BrandsService: service,
	}
}

func (uc *UBrandController) GetBrand(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}

	brand, err := uc.BrandsService.GetBrand(id)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"Brand": brand,
	})

}
func (uc *UBrandController) GetBrands(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		page = 1
	}
	search := query.Get("search")
	if search == "undefined" {
		search = ""
	}
	brands, err := uc.BrandsService.GetBrands(page, search)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"Brands": brands,
	})
}
