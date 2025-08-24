package app

import (
	"Store-Dio/controllers"
	"Store-Dio/repo"
	"Store-Dio/routes"
	"Store-Dio/services/users"
	"database/sql"
)

type App struct {
	DB             *sql.DB
	UserRoute      *routes.UserRoutes
	UserService    *users.UserService
	UserController *controllers.UserController
	UserRepo       *repo.UserRepo
}

func NewApp(db *sql.DB) *App {
	userRepo := repo.NewUserRepo(db)
	userService := users.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController)

	return &App{
		DB:             db,
		UserRoute:      userRoute,
		UserService:    userService,
		UserController: userController,
		UserRepo:       userRepo,
	}
}
