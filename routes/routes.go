package routes

import (
	"Store-Dio/controllers"

	userMiddleware "Store-Dio/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(
	controller *controllers.Controller,
	um *userMiddleware.UserMiddleware,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Post("/register", controller.UserController.CreateUser)
		r.Post("/login", controller.UserController.Login)

		r.Group(func(auth chi.Router) {
			auth.Use(um.AuthMiddleware)
			auth.Get("/profile", controller.UserController.GetUserData)
			auth.Put("/update", controller.UserController.Update)
		})
		r.Route("/refresh", func(ref chi.Router) {
			ref.Use(um.AuthMiddleware)
			ref.Post("/logout", controller.UserController.Logout)
		})
		r.Route("/products", func(product chi.Router) {
			product.Get("/", controller.UserProductController.GetProducts)
			product.Get("/{id}", controller.UserProductController.GetProduct)
			product.Group(func(prod chi.Router) {
				prod.Use(um.OnlyAdmin)
				prod.Post("/add", controller.AdminProductController.AddProduct)
				prod.Put("/{id}", controller.AdminProductController.UpdateProduct)
				prod.Delete("/{id}", controller.AdminProductController.DeleteProduct)
			})
		})
		r.Route("/brands", func(brand chi.Router) {
			brand.Get("/", controller.UserBrandsController.GetBrands)
			brand.Get("/{id}", controller.UserBrandsController.GetBrand)
			brand.Group(func(brand chi.Router) {
				brand.Use(um.OnlyAdmin)
				brand.Post("/", controller.AdminbrandsController.AddBrand)
				brand.Put("/{id}", controller.AdminbrandsController.UpdateBrand)
				brand.Delete("/{id}", controller.AdminbrandsController.DeleteBrand)
			})
		})
		r.Route("/categories", func(cat chi.Router) {
			cat.Get("/", controller.UserCategoriesController.GetCategories)
			cat.Get("/{id}", controller.UserCategoriesController.GetCategory)
			cat.Group(func(cat chi.Router) {
				cat.Use(um.OnlyAdmin)
				cat.Post("/", controller.AdminCategoriesController.AddCategory)
				cat.Put("/{id}", controller.AdminCategoriesController.UpdateCategory)
				cat.Delete("/{id}", controller.AdminCategoriesController.DeleteCategory)
			})
		})

		r.Route("/favourites", func(fav chi.Router) {
			fav.Group(func(fav chi.Router) {
				fav.Use(um.AuthMiddleware)
				fav.Get("/", controller.UserFavoriesControllr.GetFavourites)
				fav.Post("/", controller.UserFavoriesControllr.AddFavori)
				fav.Delete("/{id}", controller.UserFavoriesControllr.RemoveFavori)
			})

		})

	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Endpoint not allowed", http.StatusNotFound)
	})

	return r
}
