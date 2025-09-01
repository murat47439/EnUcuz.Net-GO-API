package app

import (
	"Store-Dio/controllers/admin"
	"Store-Dio/controllers/user"

	"Store-Dio/middleware"
	"Store-Dio/repo"
	"Store-Dio/routes"
	"Store-Dio/services/brands"
	"Store-Dio/services/categories"
	"Store-Dio/services/products"
	"Store-Dio/services/users"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type App struct {
	DB *sqlx.DB

	//User

	UserService    *users.UserService
	UserController *user.UserController
	UserRepo       *repo.UserRepo
	UserMiddleware *middleware.UserMiddleware

	//Product

	ProductService         *products.ProductService
	ProductControllerUser  *user.ProductController
	ProductControllerAdmin *admin.ProductController
	ProductRepo            *repo.ProductRepo

	// Category

	CategoriesService    *categories.CategoriesService
	CategoriesController *admin.CategoriesController
	CategoriesRepo       *repo.CategoriesRepo

	// Brands

	BrandsService    *brands.BrandsService
	BrandsController *admin.BrandsController
	BrandsRepo       *repo.BrandsRepo

	// Route

	Route *chi.Mux
}

func NewApp(db *sqlx.DB) *App {

	// User
	userRepo := repo.NewUserRepo(db)
	userMiddleware := middleware.NewUserMiddleware(userRepo)
	userService := users.NewUserService(userRepo)
	userController := user.NewUserController(userService)

	// Product

	productRepo := repo.NewProductRepo(db)
	productService := products.NewProductService(productRepo)
	productControllerUser := user.NewProductController(productService)
	productControllerAdmin := admin.NewProductController(productService)

	// Category

	categoriesRepo := repo.NewCategoriesRepo(db)
	categoriesService := categories.NewCategoriesService(categoriesRepo)
	categoriesController := admin.NewCategoriesController(categoriesService)

	// Brand

	brandsRepo := repo.NewBrandsRepo(db)
	brandsService := brands.NewBrandsService(brandsRepo)
	brandsController := admin.NewBrandsController(brandsService)

	// Route

	route := routes.SetupRoutes(userController, productControllerUser, productControllerAdmin, brandsController, categoriesController, userMiddleware)

	return &App{
		DB: db,
		//User
		UserService:    userService,
		UserController: userController,
		UserRepo:       userRepo,
		UserMiddleware: userMiddleware,

		// Product

		ProductService:         productService,
		ProductControllerUser:  productControllerUser,
		ProductControllerAdmin: productControllerAdmin,
		ProductRepo:            productRepo,

		// Category

		CategoriesService:    categoriesService,
		CategoriesController: categoriesController,
		CategoriesRepo:       categoriesRepo,

		//Brands

		BrandsService:    brandsService,
		BrandsController: brandsController,
		BrandsRepo:       brandsRepo,

		// Route

		Route: route,
	}
}
