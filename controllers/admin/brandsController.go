package admin

import (
	"Store-Dio/models"
	"Store-Dio/services/brands"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type BrandsController struct {
	BrandsService *brands.BrandsService
}

func NewBrandsController(brandsService *brands.BrandsService) *BrandsController {
	return &BrandsController{
		BrandsService: brandsService,
	}
}

func (bc *BrandsController) AddBrand(w http.ResponseWriter, r *http.Request) {
	var data *models.Brand

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	defer r.Body.Close()

	brand, err := bc.BrandsService.AddBrand(data)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Service error : %s"+err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Successfully",
		"Brand":   brand,
	})
}
func (bc *BrandsController) UpdateBrand(w http.ResponseWriter, r *http.Request) {
	var data *models.Brand

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	defer r.Body.Close()
	_, err = bc.BrandsService.UpdateBrand(data)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Successfully",
	})
}
func (bc *BrandsController) DeleteBrand(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	err = bc.BrandsService.DeleteBrand(id)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Successfully",
	})
}
