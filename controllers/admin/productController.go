package admin

import (
	"Store-Dio/models"
	"Store-Dio/services/products"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

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
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	var product models.Product

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil || id != product.ID {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}

	updproduct, err := pc.ProductService.UpdateProductForAdmin(ctx, product)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"Product": updproduct,
		"message": "Successfully",
	})
}
func (pc *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	err = pc.ProductService.DeleteProductForAdmin(ctx, id)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Successfully",
	})
}
func (pc *ProductController) GetLogs(w http.ResponseWriter, r *http.Request) {
	cwd, _ := os.Getwd()
	logPath := filepath.Join(cwd, "logs", "app.log")

	data, err := os.ReadFile(logPath)
	if err != nil {
		http.Error(w, "Cannot read log", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
