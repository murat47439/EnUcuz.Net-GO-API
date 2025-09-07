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
	query := "SELECT EXISTS (SELECT 1 FROM products WHERE id = $1 AND deleted_at IS NULL)"

	err := pr.db.Get(&exists, query, prodid)

	if err != nil {
		return false, err
	}
	config.Logger.Printf("Geldi %v", exists)

	return exists, nil

}
func (pr *ProductRepo) CheckProductByName(name, imageUrl string) (bool, error) {

	var exists bool
	if name == "" || imageUrl == "" {
		return false, fmt.Errorf("Name or ImageUrl cannot be empty")
	}
	query := "SELECT EXISTS (SELECT 1 FROM products WHERE name = $1 AND image_url = $2 AND deleted_at IS NULL)"

	err := pr.db.Get(&exists, query, name, imageUrl)

	if err != nil {
		return false, err
	}

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
	err = pr.db.Get(&product, "SELECT * FROM products WHERE id = $1 AND deleted_at IS NULL", prodid)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Product not found")
		}
		return nil, err
	}

	return &product, nil
}
func (pr *ProductRepo) GetProducts(page int, search string) ([]*models.Product, error) {
	var products []*models.Product
	offset := (page - 1) * 50
	limit := 50
	query := `SELECT * FROM products WHERE name ILIKE $1 AND deleted_at IS NULL  LIMIT $2 OFFSET $3`

	rows, err := pr.db.Queryx(query, "%"+search+"%", limit, offset) // DB: *sqlx.DB

	if err != nil {
		return nil, fmt.Errorf("Database error : %s" + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product

		if err := rows.StructScan(&p); err != nil {
			return nil, fmt.Errorf("Scan error : %s", err.Error())
		}
		products = append(products, &p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Rows error : %s", err.Error())
	}

	return products, nil
}
func (pr *ProductRepo) DeleteProduct(data *models.Product) error {
	tx, err := pr.db.Beginx()

	if err != nil {
		return fmt.Errorf("TX Error :%s", err.Error())
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	query := `UPDATE products SET deleted_at = NOW() WHERE id = $1`

	_, err = tx.Exec(query, data.ID)

	if err != nil {
		return fmt.Errorf("Database error : ", err.Error())
	}
	return nil
}
