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
func (ps *ProductService) AddProduct(data models.ProductDetail) (bool, error) {
	if data.Product.Name == "" {
		return false, fmt.Errorf("Invalid data")
	}
	err := ps.ProductRepo.AddProduct(data)

	if err != nil {
		return false, fmt.Errorf("Error : %w", err.Error())
	}
	return true, nil
}
func (ps *ProductService) UpdateProduct(product models.Product) (*models.Product, error) {
	return nil, fmt.Errorf("The service is unavailable")

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
func (ps *ProductService) GetProduct(id int) (*models.ProductDetail, error) {
	if id == 0 {
		return nil, fmt.Errorf("Invalid data")
	}
	product, err := ps.ProductRepo.GetProductDetail(id)

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
func (ps *ProductService) CompareProducts(id1, id2 int) ([]models.ProductDetail, error) {
	if id1 == 0 || id2 == 0 {
		return []models.ProductDetail{}, fmt.Errorf("Invalid data")
	}
	result, err := ps.ProductRepo.CompareProduct(id1, id2)
	if err != nil {
		return []models.ProductDetail{}, nil
	}
	return result, nil
}
