package admin

import (
	"Store-Dio/models"
	"Store-Dio/services/products"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductController struct {
	ProductService *products.ProductService
}

func NewProductController(service *products.ProductService) *ProductController {
	return &ProductController{ProductService: service}
}
func (pc *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	var product models.Product

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil || id == product.ID {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}

	updproduct, err := pc.ProductService.UpdateProduct(product)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"Product": updproduct,
		"Message": "Successfully",
	})
}
func (pc *ProductController) AddProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}

	_, err = pc.ProductService.AddProduct(product)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Successfully",
	})
}
func (pc *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	err = pc.ProductService.DeleteProduct(id)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]string{
		"Message": "Successfully",
	})
}
