package controllers

import (
	"Store-Dio/models"
	"Store-Dio/services/products"
	"encoding/json"
	"net/http"
)

type ProductController struct {
	ProductService *products.ProductService
}

func NewProductController(service *products.ProductService) *ProductController {
	return &ProductController{ProductService: service}
}

func (pr *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}

}
