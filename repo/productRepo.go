package repo

import (
	"Store-Dio/config"
	"Store-Dio/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db    *sqlx.DB
	brand *BrandsRepo
	cat   *CategoriesRepo
}

func NewProductRepo(db *sqlx.DB, brand *BrandsRepo, cat *CategoriesRepo) *ProductRepo {
	return &ProductRepo{
		db:    db,
		brand: brand,
		cat:   cat}
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
func (pr *ProductRepo) AddProduct(ctx context.Context, data models.Product, tx *sqlx.Tx) error {
	query := `INSERT INTO products(name,description,stock,image_url,category_id,created_at,brand_id,seller_id) VALUES($1,$2,$3,$4,$5,NOW(),$6,$7)`
	_, err := tx.ExecContext(ctx, query, data.Name, data.Description, data.Stock, data.ImageUrl, data.CategoryId, data.BrandID, data.SellerID)
	if err != nil {
		return fmt.Errorf("Database error %w", err)
	}
	return nil
}
func (pr *ProductRepo) ExistsData(name string, tx *sqlx.Tx) (bool, error) {
	if name == "" {
		return false, fmt.Errorf("Invalid data")
	}
	query := `SELECT EXISTS(SELECT 1 FROM brands WHERE name = $1 AND deleted_at IS NULL)`
	var exists bool
	err := tx.QueryRow(query, name).Scan(&exists)

	if err != nil {
		return false, err
	}
	return exists, nil
}
func (pr *ProductRepo) UpdateProduct(ctx context.Context, product *models.Product) error {
	query := `UPDATE products SET name = $1, description = $2,stock = $3 WHERE id = $4 AND deleted_at IS NULL`

	res, err := pr.db.ExecContext(ctx, query, product.Name, product.Description, product.Stock, product.ID)
	if err != nil {
		return fmt.Errorf("Database error : %w", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("Product Not Found")
	}
	return nil

}
func (pr *ProductRepo) GetProduct(ctx context.Context, prodid int) (*models.Product, error) {
	var product models.Product
	if prodid == 0 {
		return nil, fmt.Errorf("Invalid data")
	}
	_, err := pr.CheckProduct(prodid)
	if err != nil {
		return nil, err
	}
	err = pr.db.GetContext(ctx, &product, "SELECT * FROM products WHERE id = $1 AND deleted_at IS NULL", prodid)

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
	limit := 52
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
	query := `UPDATE products SET deleted_at = NOW() WHERE id = $1 AND seller_id = $2`

	_, err = tx.Exec(query, data.ID, data.SellerID)

	if err != nil {
		return fmt.Errorf("Database error : ", err.Error())
	}
	return nil
}
