package products

import (
	"Store-Dio/models"
	"Store-Dio/repo"
	"fmt"
)

type ProductService struct {
	ProductRepo *repo.ProductRepo
}

func NewProductService(repo *repo.ProductRepo) *ProductService {
	return &ProductService{ProductRepo: repo}
}

func (ps *ProductService) AddProduct(product models.Product) (bool, error) {
	if product.Name == "" || product.Description == "" || product.BrandID == 0 || product.Stock == 0 || product.ImageUrl == "" || product.CategoryID == 0 || product.StoreID == 0 {
		return false, fmt.Errorf("Invalid data")
	}
	exists, err := ps.ProductRepo.CheckProduct(&models.FilterProd{ID: product.ID, ImageUrl: product.ImageUrl, Name: product.Name})
	if err != nil {
		return false, err
	}
	if exists {
		return false, fmt.Errorf("Product already exists")
	}
	_, err = ps.ProductRepo.AddProduct(&product)

	if err != nil {
		return false, err
	}
	return true, nil
}
func (ps *ProductService) UpdateProduct(product models.Product) (*models.Product, error) {
	if product.ID == 0 || product.Name == "" || product.Description == "" || product.BrandID == 0 || product.ImageUrl == "" || product.CategoryID == 0 || product.StoreID == 0 {
		return nil, fmt.Errorf("Invalid data")
	}
	_, err := ps.ProductRepo.UpdateProduct(&product)

	if err != nil {
		return nil, err
	}

	return &product, nil

}
func (ps *ProductService) DeleteProduct(id int) error {
	if id == 0 {
		return fmt.Errorf("Invalid data")
	}
	data, err := ps.ProductRepo.GetProduct(id)

	if err != nil {
		return err
	}

	err = ps.ProductRepo.DeleteProduct(data)

	if err != nil {
		return err
	}
	return nil
}
func (ps *ProductService) GetProduct(id int) (*models.Product, error) {
	if id == 0 {
		return nil, fmt.Errorf("Invalid data")
	}
	product, err := ps.ProductRepo.GetProduct(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}
func (ps *ProductService) GetProducts(page int, search string) ([]*models.Product, error) {
	if page < 1 {
		page = 1
	}

	products, err := ps.ProductRepo.GetProducts(page, search)

	if err != nil {
		return nil, fmt.Errorf("Error : %s" + err.Error())
	}

	return products, nil
}
