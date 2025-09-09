package controllers

import (
	"Store-Dio/controllers/admin"
	"Store-Dio/controllers/user"
	"Store-Dio/services"
)

type Controller struct {
	AdminbrandsController     *admin.BrandsController
	AdminCategoriesController *admin.CategoriesController
	AdminProductController    *admin.ProductController
	AdminReviewController     *admin.ReviewController

	UserProductController    *user.ProductController
	UserController           *user.UserController
	UserBrandsController     *user.UBrandController
	UserCategoriesController *user.UCategoriesController
	UserFavoriesControllr    *user.FavoriesController
	UserReviewController     *user.ReviewController
}

func NewController(service *services.Service) *Controller {
	adminBrandsController := admin.NewBrandsController(service.BrandsService)
	adminCategoriesController := admin.NewCategoriesController(service.CategoriesService)
	adminProductController := admin.NewProductController(service.ProductsService)
	adminReviewController := admin.NewReviewController(service.ReviewsService)

	userProductController := user.NewProductController(service.ProductsService)
	userBrandsController := user.NewUBrandController(service.BrandsService)
	userCategoriesController := user.NewUCategoriesController(service.CategoriesService)
	userController := user.NewUserController(service.UsersService)
	userFavoritesController := user.NewFavoriesController(service.FavoriesService)
	userReviewController := user.NewReviewController(service.ReviewsService)

	return &Controller{
		AdminbrandsController:     adminBrandsController,
		AdminCategoriesController: adminCategoriesController,
		AdminProductController:    adminProductController,
		AdminReviewController:     adminReviewController,

		UserProductController:    userProductController,
		UserController:           userController,
		UserBrandsController:     userBrandsController,
		UserCategoriesController: userCategoriesController,
		UserFavoriesControllr:    userFavoritesController,
		UserReviewController:     userReviewController,
	}
}
