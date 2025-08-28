package routes

import (
	"Store-Dio/controllers"
	"Store-Dio/middleware"
	"net/http"
)

type ProductRoutes struct {
	ProductController *controllers.ProductController
	UserMiddleware    *middleware.UserMiddleware
}

func NewProductRoute(pc *controllers.ProductController, um *middleware.UserMiddleware) *ProductRoutes {
	return &ProductRoutes{
		ProductController: pc,
		UserMiddleware:    um,
	}
}
func (pr *ProductRoutes) SetupProductRoute() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		{

		}
	})
}
