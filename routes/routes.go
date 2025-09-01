package routes

import (
	"Store-Dio/controllers/admin"
	"Store-Dio/controllers/user"

	userMiddleware "Store-Dio/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(
	userController *user.UserController,
	productControllerUser *user.ProductController,
	productControllerAdmin *admin.ProductController,
	brandsController *admin.BrandsController,
	categoriesController *admin.CategoriesController,
	um *userMiddleware.UserMiddleware,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/users", func(user chi.Router) {
		user.Post("/register", userController.CreateUser)
		user.Post("/login", userController.Login)

		user.Group(func(auth chi.Router) {
			auth.Use(um.AuthMiddleware)
			auth.Get("/profile", userController.GetUserData)
			auth.Put("/update", userController.Update)
			auth.Post("/logout", userController.Logout)
		})

		user.Group(func(admin chi.Router) {
			admin.Use(um.OnlyAdmin)
			admin.Get("/admin", userController.GetUserData)
		})
	})

	r.Route("/products", func(product chi.Router) {
		product.Get("/", productControllerUser.GetProducts)
		product.Get("/{id}", productControllerUser.GetProduct)
		product.Group(func(prod chi.Router) {
			prod.Use(um.OnlyAdmin)
			prod.Post("/add", productControllerAdmin.AddProduct)
			prod.Put("/update", productControllerAdmin.UpdateProduct)
		})
	})

	r.Route("/admin", func(admin chi.Router) {
		admin.Group(func(ad chi.Router) {
			ad.Use(um.OnlyAdmin)
			ad.Post("/import-categories", categoriesController.InsertCategoriesData)
			ad.Post("/import-brands", brandsController.InsertBrandData)
			ad.Get("/get-categories-id", categoriesController.GetAllCategoriesID)
		})
	})
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Endpoint not allowed", http.StatusNotFound)
	})

	return r
}
