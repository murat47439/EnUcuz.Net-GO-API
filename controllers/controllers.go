package controllers

import (
	"Store-Dio/controllers/admin"
	"Store-Dio/controllers/user"
	"Store-Dio/services"
)

type Controller struct {
	AdminAttributeController  *admin.AttributeController
	AdminbrandsController     *admin.BrandsController
	AdminCategoriesController *admin.CategoriesController
	AdminProductController    *admin.ProductController

	UserProductController    *user.ProductController
	UserController           *user.UserController
	UserBrandsController     *user.UBrandController
	UserCategoriesController *user.UCategoriesController
	UserFavoriesControllr    *user.FavoriesController
}

func NewController(service *services.Service) *Controller {
	adminAttributeController := admin.NewAttributeController(service.AttributesService)
	adminBrandsController := admin.NewBrandsController(service.BrandsService)
	adminCategoriesController := admin.NewCategoriesController(service.CategoriesService)
	adminProductController := admin.NewProductController(service.ProductsService)

	userProductController := user.NewProductController(service.ProductsService)
	userBrandsController := user.NewUBrandController(service.BrandsService)
	userCategoriesController := user.NewUCategoriesController(service.CategoriesService)
	userController := user.NewUserController(service.UsersService)
	userFavoritesController := user.NewFavoriesController(service.FavoriesService)

	return &Controller{
		AdminAttributeController:  adminAttributeController,
		AdminbrandsController:     adminBrandsController,
		AdminCategoriesController: adminCategoriesController,
		AdminProductController:    adminProductController,

		UserProductController:    userProductController,
		UserController:           userController,
		UserBrandsController:     userBrandsController,
		UserCategoriesController: userCategoriesController,
		UserFavoriesControllr:    userFavoritesController,
	}
}
