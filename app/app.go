package app

import (
	"Store-Dio/controllers"
	"Store-Dio/middleware"
	"Store-Dio/repo"
	"Store-Dio/routes"
	"Store-Dio/services/users"

	"github.com/jmoiron/sqlx"
)

type App struct {
	DB             *sqlx.DB
	UserRoute      *routes.UserRoutes
	UserService    *users.UserService
	UserController *controllers.UserController
	UserRepo       *repo.UserRepo
	UserMiddleware *middleware.UserMiddleware
}

func NewApp(db *sqlx.DB) *App {
	userRepo := repo.NewUserRepo(db)
	userMiddleware := middleware.NewUserMiddleware(userRepo)
	userService := users.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController, userMiddleware)

	return &App{
		DB:             db,
		UserRoute:      userRoute,
		UserService:    userService,
		UserController: userController,
		UserRepo:       userRepo,
		UserMiddleware: userMiddleware,
	}
}
