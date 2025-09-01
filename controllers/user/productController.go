package user

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

func (pc *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	product, err := pc.ProductService.GetProduct(id)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"id":          product.ID,
		"name":        product.Name,
		"description": product.Description,
		"brand_id":    product.BrandID,
		"stock":       product.Stock,
		"category_id": product.CategoryID,
		"store_id":    product.StoreID,
		"image_url":   product.ImageUrl,
	})

}

func (pc *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}

}
