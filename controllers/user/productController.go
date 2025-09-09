package user

import (
	"Store-Dio/services/products"
	"net/http"
	"os"
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
		"Product": product,
	})

}
func (pc *ProductController) GetLogs(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("../../logs/app.log")
	if err != nil {
		http.Error(w, "Cannot read log", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (pc *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		page = 1
	}
	search := query.Get("search")

	products, err := pc.ProductService.GetProducts(page, search)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Error : %s"+err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message":  "Successfully",
		"Products": products,
	})
}
