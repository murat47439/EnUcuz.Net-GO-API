package repo

import (
	"Store-Dio/config"
	"Store-Dio/models"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) *ProductRepo {
	return &ProductRepo{db: db}
}
func (pr *ProductRepo) CheckProduct(prodid int) (bool, error) {

	var exists bool
	if prodid == 0 {
		return false, fmt.Errorf("Invalid product id")
	}
	query := "SELECT EXISTS (SELECT 1 FROM products WHERE id = $1)"

	err := pr.db.Get(&exists, query, prodid)

	if err != nil {
		return false, err
	}
	config.Logger.Printf("Geldi %v", exists)

	return exists, nil

}
func (pr *ProductRepo) AddProduct(product *models.Product) (bool, error) {
	if product.Name == "" || product.Description == "" || product.BrandID == 0 || product.Stock == 0 || product.ImageUrl == "" || product.CategoryID == 0 || product.StoreID == 0 {
		return false, fmt.Errorf("Invalid data")
	}
	query := "INSERT INTO products (store_id, name, description, brand_id, stock, image_url, category_id) VALUES (:store_id, :name, :description, :brand_id, :stock, :image_url, :category_id)"

	_, err := pr.db.NamedExec(query, map[string]interface{}{
		"store_id":    product.StoreID,
		"name":        product.Name,
		"description": product.Description,
		"brand_id":    product.BrandID,
		"stock":       product.Stock,
		"image_url":   product.ImageUrl,
		"category_id": product.CategoryID,
	})
	if err != nil {
		return false, fmt.Errorf("Failed to add product to database.")
	}
	return true, nil
}
func (pr *ProductRepo) UpdateProduct(product *models.Product) (bool, error) {
	if product.ID == 0 || product.Name == "" || product.Description == "" || product.BrandID == 0 || product.Stock == 0 || product.ImageUrl == "" || product.CategoryID == 0 || product.StoreID == 0 {
		return false, fmt.Errorf("Invalid data")
	}

	exists, err := pr.CheckProduct(product.ID)

	if err != nil {
		return false, err
	}
	if !exists {
		return false, fmt.Errorf("Product not found")
	}

	query := `
			UPDATE products 
			SET name = :name,
				description = :description,
				brand_id = :brand_id,
				stock = :stock,
				image_url = :image_url,
				category_id = :category_id,
				store_id = :store_id
			WHERE id = :id
			`

	_, err = pr.db.NamedExec(query, map[string]interface{}{
		"id":          product.ID,
		"name":        product.Name,
		"description": product.Description,
		"brand_id":    product.BrandID,
		"stock":       product.Stock,
		"image_url":   product.ImageUrl,
		"category_id": product.CategoryID,
		"store_id":    product.StoreID,
	})

	if err != nil {
		config.Logger.Printf("Failed update product")
		return false, fmt.Errorf("Failed update product ")
	}
	return true, nil
}
func (pr *ProductRepo) GetProduct(prodid int) (*models.Product, error) {
	var product models.Product
	if prodid == 0 {
		return nil, fmt.Errorf("Invalid data")
	}
	_, err := pr.CheckProduct(prodid)
	if err != nil {
		return nil, err
	}
	err = pr.db.Get(&product, "SELECT * FROM products WHERE id = $1", prodid)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Product not found")
		}
		return nil, err
	}

	return &product, nil
}
