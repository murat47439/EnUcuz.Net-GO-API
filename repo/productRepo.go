package repo

import (
	"Store-Dio/config"
	"Store-Dio/models"
	"database/sql"
	"fmt"
	"sync"

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
			if commitErr := tx.Commit(); commitErr != nil {
				err = fmt.Errorf("tx commit error: %w", commitErr)
			}
		}
	}()
	err = pr.InsertBrands(data.Brand, tx)
	if err != nil {
		return err
	}
	query := `INSERT INTO products(id,name,category_id,brand_id,launch_announced,launch_released, launch_status) VALUES($1,$2,$3, $4, $5, $6, $7)`

	_, err = tx.Exec(query, data.ID, data.Name, 3719, data.Brand.ID, data.Launch.Announced, data.Launch.Released, data.Launch.Status)
	if err != nil {
		return err
	}
	query = `
INSERT INTO phone_details(
    product_id, current_os, upgradable_to, chipset, cpu, gpu,
    body_dimensions, body_weight, body_build, sim_info,
    network_technology, network_speed, network_2g, network_3g, network_4g, network_5g,
    gps, nfc, radio, wlan, bluetooth, usb, card_slot
)
VALUES (
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23
)
`
	_, err = tx.Exec(
		query,
		data.ID,
		data.Platform.CurrentOS,
		data.Platform.UpgradableOS,
		data.Platform.Chipset,
		data.Platform.CPU,
		data.Platform.GPU,
		data.Body.Dimensions,
		data.Body.Weight,
		data.Body.Build,
		data.Body.SIM,
		data.Network.Technology,
		data.Network.Speed,
		data.Network.G2,
		data.Network.G3,
		data.Network.G4,
		data.Network.G5,
		data.Comms.Positioning,
		data.Comms.NFC,
		data.Comms.Radio,
		data.Comms.WLAN,
		data.Comms.Bluetooth,
		data.Comms.USB,
		data.Memory.CardSlot,
	)
	if err != nil {
		return fmt.Errorf("insert phone_details error: %v", err)
	}
	err = pr.InsertBattery(data.Battery, data.ID, tx)
	if err != nil {
		return err
	}
	err = pr.InsertDisplay(data.Display, data.ID, tx)
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
	query := `INSERT INTO colors(product_id, color) VALUES($1, $2)`

	for _, color := range data {
		_, err := tx.Exec(query, id, color)
		if err != nil {
			return err
		}
	}

	return nil
}
func (pr *ProductRepo) InsertModels(data []string, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO models(product_id, model) VALUES($1, $2)`

	_, err := tx.Exec(query, id, pq.Array(data))
	if err != nil {
		return err
	}

	return nil
}
func (pr *ProductRepo) InsertCamera(data models.Camera, id int, role string, tx *sqlx.Tx) error {
	query := `INSERT INTO cameras(product_id, type) VALUES($1,$2) RETURNING id`
	var cid int

	err := tx.QueryRowx(query, id, role).Scan(&cid)
	if err != nil {
		return err
	}
	query = `INSERT INTO camera_lenses(camera_id, megapixels,aperture,focal_length, sensor_size, type, pixel_size,other_features,zoom) VALUES($1,$2,$3,$4,$5,$6, $7, $8, $9)`
	for _, dat := range data.Lenses {
		_, err = tx.Exec(query, cid, dat.Megapixels, dat.Aperture, dat.FocalLength, dat.SensorSize, dat.Type, dat.PixelSize, pq.Array(dat.OtherFeatures), dat.Zoom)
		if err != nil {
			return err
		}
	}
	query = `INSERT INTO camera_features(camera_id, feature) VALUES($1,$2)`
	for _, fat := range data.Features {
		_, err = tx.Exec(query, cid, fat.Spec)
		if err != nil {
			return err
		}
	}
	query = `INSERT INTO camera_video(camera_id, video_spec) VALUES($1,$2)`
	for _, vat := range data.Video {
		_, err = tx.Exec(query, cid, vat.Spec)
		if err != nil {
			return err
		}
	}
	return nil
}
func (pr *ProductRepo) InsertFeatures(data models.Features, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO sensors(product_id,sensors) VALUES($1,$2)`

	_, err := tx.Exec(query, id, pq.Array(data.Sensors))
	if err != nil {
		return err
	}
	return nil
}
func (pr *ProductRepo) InsertSound(data models.Sound, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO sound_specs(product_id,loudspeaker,features) VALUES($1,$2,$3)`

	_, err := tx.Exec(query, id, data.Loudspeaker, data.Features)
	if err != nil {
		return err
	}
	return nil
}
func (pr *ProductRepo) InsertMemory(data models.Memory, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO memory_options(product_id,storage,ram) VALUES($1, $2,$3)`
	for _, dat := range data.InternalOptions {
		_, err := tx.Exec(query, id, dat.Storage, dat.RAM)
		if err != nil {
			return err
		}
	}
	return nil
}
func (pr *ProductRepo) InsertDisplay(data models.Display, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO display(product_id,type,size,resolution,protection,aspect_ratio,hdr,refresh_rate, brightness_typical,brightness_hbm,other_features) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`

	_, err := tx.Exec(query, id, data.PanelType, data.SizeInches, data.ResolutionPixels, data.Protection, data.AspectRatio, pq.Array(data.HDR), data.RefreshRate, data.Brightness.Typical, data.Brightness.Hbm, pq.Array(data.OtherFeatures))
	if err != nil {
		return err
	}
	return nil
}
func (pr *ProductRepo) InsertBattery(data models.Battery, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO battery(product_id,technology , capacity) VALUES($1,$2,$3) RETURNING id`

	var bid int
	err := tx.QueryRowx(query, id, data.Technology, data.Capacity).Scan(&bid)
	if err != nil {
		return err
	}
	for _, dat := range data.ChargingDetails {
		query = `INSERT INTO battery_detail(battery_id, type, description, power) VALUES($1,$2,$3,$4)`

		_, err = tx.Exec(query, bid, dat.Type, dat.Description, dat.Power)
		if err != nil {
			return fmt.Errorf("Charging detail error : %v", err)
		}
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
func (pr *ProductRepo) InsertData(data []models.ProductDetail) error {
	if data == nil {
		return fmt.Errorf("Invalid data")
	}
	for _, dat := range data {
		config.Logger.Printf("GELDİS")
		if err := pr.AddProduct(dat); err != nil {
			return fmt.Errorf("failed to insert product %d: %w", dat.ID, err)
		}
	}

	return nil

}
