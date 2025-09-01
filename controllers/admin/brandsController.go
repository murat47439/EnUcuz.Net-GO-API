package admin

import (
	"Store-Dio/models"
	"Store-Dio/services/brands"
	"encoding/json"
	"net/http"
)

type BrandsController struct {
	BrandsService *brands.BrandsService
}

func NewBrandsController(brandsService *brands.BrandsService) *BrandsController {
	return &BrandsController{
		BrandsService: brandsService,
	}
}

func (bc *BrandsController) InsertBrandData(w http.ResponseWriter, r *http.Request) {
	var brands *models.Brands

	err := json.NewDecoder(r.Body).Decode(&brands)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	err = bc.BrandsService.InsertBrandData(brands)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Error : %w"+err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Successfully",
	})
}
