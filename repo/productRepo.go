package repo

import (
	"Store-Dio/config"
	"Store-Dio/models"
	"database/sql"
	"fmt"
	"sync"

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

func (pr *ProductRepo) AddProduct(data models.ProductDetail) error {
	return nil
}

func (pr *ProductRepo) InsertBrands(data string, tx *sqlx.Tx) (int, error) {
	exists, err := pr.ExistsData(data, tx)
	var id int
	if err != nil {
		return 0, err
	}
	if exists {
		query := `SELECT id FROM brands WHERE name = $1 AND deleted_at IS NULL`
		err = tx.QueryRowx(query, data).Scan(&id)
		if err != nil {
			return 0, err
		}
		return id, nil
	}
	query := `INSERT INTO brands (name,created_at) VALUES ($1, NOW()) RETURNING id`

	err = tx.QueryRowx(query, data).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil

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
func (pr *ProductRepo) UpdateProduct(product *models.Product) (bool, error) {
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
func (pr *ProductRepo) GetProductDetail(prodid int) (*models.ProductDetail, error) {
	return nil, nil
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
	query := `UPDATE products SET deleted_at = NOW() WHERE id = $1`

	_, err = tx.Exec(query, data.ID)

	if err != nil {
		return fmt.Errorf("Database error : ", err.Error())
	}
	return nil
}
func (pr *ProductRepo) CompareProduct(prodid1, prodid2 int) ([]models.ProductDetail, error) {
	if prodid1 == 0 || prodid2 == 0 {
		return nil, fmt.Errorf("Invalid data")
	}
	var prods []models.ProductDetail
	var wg sync.WaitGroup
	var prod1, prod2 *models.ProductDetail
	var err1, err2 error

	wg.Add(2)
	go func() { defer wg.Done(); prod1, err1 = pr.GetProductDetail(prodid1) }()
	go func() { defer wg.Done(); prod2, err2 = pr.GetProductDetail(prodid2) }()
	wg.Wait()

	if err1 != nil || err2 != nil {
		return nil, fmt.Errorf("Hata1: %v, Hata2: %v", err1, err2)
	}

	prods = append(prods, *prod1, *prod2)
	return prods, nil
}
