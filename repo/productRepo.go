package repo

import (
	"Store-Dio/config"
	"Store-Dio/models"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ProductRepo struct {
	db  *sqlx.DB
	psr *ProductSpecsRepo
}

func NewProductRepo(db *sqlx.DB, psr *ProductSpecsRepo) *ProductRepo {
	return &ProductRepo{
		db:  db,
		psr: psr}
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

	tx, err := pr.db.Beginx()
	if err != nil {
		return err
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
	err = pr.InsertBrands(data.Brand, tx)
	if err != nil {
		return err
	}
	query := `INSERT INTO products(id,name,category_id,brand_id) VALUES($1,$2,$3, $4)`

	_, err = tx.Exec(query, data.ID, data.Name, 3719, data.Brand.ID)
	if err != nil {
		return err
	}
	err = pr.InsertBattery(data.Battery, data.ID, tx)
	if err != nil {
		return err
	}
	err = pr.InsertPlatform(data.Platform, data.ID, tx)
	if err != nil {
		return err
	}
	err = pr.InsertNetwork(data.Network, data.ID, tx)
	if err != nil {
		return err
	}
	err = pr.InsertDisplay(data.Display, data.ID, tx)
	if err != nil {
		return err
	}
	err = pr.InsertLaunch(data.Launch, data.ID, tx)
	if err != nil {
		return err
	}
	err = pr.InsertBody(data.Body, data.ID, tx)
	if err != nil {
		return err
	}
	err = pr.InsertMemory(data.Memory, data.ID, tx)
	if err != nil {
		return err
	}
	err = pr.InsertSound(data.Sound, data.ID, tx)
	if err != nil {
		return err
	}
	err = pr.InsertComms(data.Comms, data.ID, tx)
	if err != nil {
		return err
	}
	err = pr.InsertFeatures(data.Features, data.ID, tx)
	if err != nil {
		return err
	}
	err = pr.InsertCamera(data.Cameras.MainCamera, data.ID, "MainCamera", tx)
	if err != nil {
		return err
	}
	err = pr.InsertCamera(data.Cameras.SelfieCamera, data.ID, "SelfieCamera", tx)
	if err != nil {
		return err
	}
	err = pr.InsertColor(data.Colors, data.ID, tx)
	if err != nil {
		return err
	}
	err = pr.InsertModels(data.Models, data.ID, tx)
	if err != nil {
		return err
	}

	return nil
}
func (pr *ProductRepo) InsertColor(data []string, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO product_colors(product_id, color) VALUES($1, $2)`

	for _, color := range data {
		_, err := tx.Exec(query, id, color)
		if err != nil {
			return err
		}
	}

	return nil
}
func (pr *ProductRepo) InsertModels(data []string, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO product_models(product_id, model) VALUES($1, $2)`

	for _, model := range data {
		_, err := tx.Exec(query, id, model)
		if err != nil {
			return err
		}
	}

	return nil
}
func (pr *ProductRepo) InsertCamera(data models.Camera, id int, role string, tx *sqlx.Tx) error {
	query := `INSERT INTO cameras(product_id, camera_type, camera_specs, features, video, camera_role) VALUES($1,$2,$3,$4,$5, $6)`

	_, err := tx.Exec(query, id, data.Type, pq.Array(data.CameraSpecs), pq.Array(data.Features), pq.Array(data.Video), role)

	if err != nil {
		return err
	}
	return nil
}
func (pr *ProductRepo) InsertFeatures(data models.Features, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO feature_specs(product_id,sensors) VALUES($1,$2)`

	_, err := tx.Exec(query, id, data.Sensors)
	if err != nil {
		return err
	}
	return nil
}
func (pr *ProductRepo) InsertComms(data models.Comms, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO comms_specs(product_id,wlan,bluetooth,positioning,nfc,radio,usb) VALUES($1,$2,$3,$4,$5,$6,$7)`

	_, err := tx.Exec(query, id, data.WLAN, data.Bluetooth, data.Positioning, data.NFC, data.Radio, data.USB)
	if err != nil {
		return err
	}
	return nil
}
func (pr *ProductRepo) InsertSound(data models.Sound, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO sound_specs(product_id,loudspeaker) VALUES($1,$2)`

	_, err := tx.Exec(query, id, data.Loudspeaker)
	if err != nil {
		return err
	}
	return nil
}
func (pr *ProductRepo) InsertMemory(data models.Memory, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO memory_specs(product_id,card_slot,internal) VALUES($1, $2,$3)`

	_, err := tx.Exec(query, id, data.CardSlot, data.Internal)
	if err != nil {
		return err
	}
	return nil
}
func (pr *ProductRepo) InsertBody(data models.Body, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO body_specs(product_id, dimensions, weight,build,sim) VALUES($1,$2,$3,$4,$5)`

	_, err := tx.Exec(query, id, data.Dimensions, data.Weight, data.Build, data.SIM)
	if err != nil {
		return err
	}
	return nil
}
func (pr *ProductRepo) InsertLaunch(data models.Launch, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO launch_specs(product_id,announced,released,status) VALUES($1, $2,$3,$4)`

	_, err := tx.Exec(query, id, data.Announced, data.Released, data.Status)
	if err != nil {
		return err
	}
	return nil
}
func (pr *ProductRepo) InsertDisplay(data models.Display, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO display_specs(product_id,type,size,resolution,protection) VALUES($1,$2,$3,$4,$5)`

	_, err := tx.Exec(query, id, data.Type, data.Size, data.Resolution, data.Protection)
	if err != nil {
		return err
	}
	return nil
}
func (pr *ProductRepo) InsertNetwork(data models.Network, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO network_specs(product_id, technology, speed, g2 , g3, g4 , g5) VALUES($1,$2,$3,$4,$5,$6,$7)`

	_, err := tx.Exec(query, id, data.Technology, data.Speed, data.G2, data.G3, data.G4, data.G5)
	if err != nil {
		return err
	}
	return nil
}
func (pr *ProductRepo) InsertPlatform(data models.Platform, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO platform_specs(product_id,os,chipset,cpu,gpu) VALUES($1,$2,$3,$4,$5) `

	_, err := tx.Exec(query, id, data.OS, data.Chipset, data.CPU, data.GPU)

	if err != nil {
		return err
	}
	return nil
}
func (pr *ProductRepo) InsertBattery(data models.Battery, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO battery_specs(product_id, type, charging) VALUES($1,$2,$3)`

	_, err := tx.Exec(query, id, data.Type, pq.Array(data.Charging))
	if err != nil {
		return err
	}
	return nil
}
func (pr *ProductRepo) InsertBrands(data *models.Brand, tx *sqlx.Tx) error {
	exists, err := pr.ExistsData(data.ID, tx)

	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	query := `INSERT INTO brands (id,name,created_at) VALUES ($1, $2, NOW())`

	_, err = tx.Exec(query, data.ID, data.Name)

	if err != nil {
		return err
	}
	return nil

}
func (pr *ProductRepo) ExistsData(id int, tx *sqlx.Tx) (bool, error) {
	if id == 0 {
		return false, fmt.Errorf("Invalid data")
	}
	query := `SELECT EXISTS(SELECT 1 FROM brands WHERE id = $1)`
	var exists bool
	err := tx.QueryRow(query, id).Scan(&exists)

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
	var product models.Product
	if prodid == 0 {
		return nil, fmt.Errorf("Invalid data")
	}
	query := `SELECT id,name,image_url,category_id,brand_id FROM products WHERE id = $1; `

	err := pr.db.Get(&product, query, prodid)

	if err != nil {
		return nil, fmt.Errorf("Database error : %w", err)
	}
	productDetail, err := pr.psr.GetProductDetail(&product)

	if err != nil {
		return nil, err
	}
	return productDetail, nil
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
