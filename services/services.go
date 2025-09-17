package services

import (
	"Store-Dio/repo"
	"Store-Dio/services/attributes"
	"Store-Dio/services/brands"
	"Store-Dio/services/categories"
	"Store-Dio/services/favories"
	"Store-Dio/services/products"
	"Store-Dio/services/reviews"
	"Store-Dio/services/users"
)

type Service struct {
	BrandsService     *brands.BrandsService
	CategoriesService *categories.CategoriesService
	ProductsService   *products.ProductService
	UsersService      *users.UserService
	FavoriesService   *favories.FavoriesService
	ReviewsService    *reviews.ReviewService
	AttributeService  *attributes.AttributeService
}

func NewService(repo *repo.Repo) *Service {

	brandsService := brands.NewBrandsService(repo.BrandsRepo)
	categoriesService := categories.NewCategoriesService(repo.CategoriesRepo)
	productsService := products.NewProductService(repo.ProductRepo)
	usersService := users.NewUserService(repo.UserRepo)
	favoriesService := favories.NewFavoriesService(repo.FavoriesRepo)
	reviewsService := reviews.NewReviewService(repo.ReviewsRepo)
	attributeService := attributes.NewAttributeService(repo.AttributeRepo, repo.ProductRepo)

	return &Service{

		BrandsService:     brandsService,
		CategoriesService: categoriesService,
		ProductsService:   productsService,
		UsersService:      usersService,
		FavoriesService:   favoriesService,
		ReviewsService:    reviewsService,
		AttributeService:  attributeService,
	}
}
