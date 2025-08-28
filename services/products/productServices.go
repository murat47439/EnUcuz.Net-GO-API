package products

import "Store-Dio/repo"

type ProductService struct {
	ProductRepo *repo.ProductRepo
}

func NewProductService(repo *repo.ProductRepo) *ProductService {
	return &ProductService{ProductRepo: repo}
}
