package app

import (
	"Store-Dio/controllers"
	"Store-Dio/middleware"
	"Store-Dio/repo"
	"Store-Dio/routes"
	"Store-Dio/services/products"
	"Store-Dio/services/users"

	"github.com/jmoiron/sqlx"
)

type App struct {
	DB *sqlx.DB

	//User

	UserRoute      *routes.UserRoutes
	UserService    *users.UserService
	UserController *controllers.UserController
	UserRepo       *repo.UserRepo
	UserMiddleware *middleware.UserMiddleware

	//Product

	ProductRoute      *routes.ProductRoutes
	ProductService    *products.ProductService
	ProductController *controllers.ProductController
	ProductRepo       *repo.ProductRepo
}

func NewApp(db *sqlx.DB) *App {
	userRepo := repo.NewUserRepo(db)
	userMiddleware := middleware.NewUserMiddleware(userRepo)
	userService := users.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController, userMiddleware)

	return &App{
		DB: db,
		//User
		UserRoute:      userRoute,
		UserService:    userService,
		UserController: userController,
		UserRepo:       userRepo,
		UserMiddleware: userMiddleware,
	}
}
