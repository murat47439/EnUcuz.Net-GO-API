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
	db    *sqlx.DB
	psr   *ProductSpecsRepo
	brand *BrandsRepo
	cat   *CategoriesRepo
}

func NewProductRepo(db *sqlx.DB, psr *ProductSpecsRepo, brand *BrandsRepo, cat *CategoriesRepo) *ProductRepo {
	return &ProductRepo{
		db:    db,
		psr:   psr,
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
	id, err := pr.InsertBrands(data.Product.Brand, tx)
	if err != nil {
		return err
	}
	query := `INSERT INTO products(name,category_id,brand_id,launch_announced,launch_released, launch_status) VALUES($1,$2,$3, $4, $5, $6) RETURNING id`
	var prodID int
	err = tx.QueryRowx(query, data.Product.Name, 3719, id, data.Product.Announced, data.Product.Released, data.Product.Status).Scan(&prodID)
	if err != nil {
		return err
	}
	query = `
		INSERT INTO phone_details(
			product_id, current_os, upgradable_to, chipset, cpu, gpu,
			dimensions, weight, build, sim_info,
			network_technology, network_speed, g2, g3, g4, g5,
			gps, nfc, radio, wlan, bluetooth, usb, card_slot
		)
		VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23
		)
		`
	_, err = tx.Exec(
		query,
		prodID,
		data.PhoneDetail.CurrentOS,
		data.PhoneDetail.UpgradableOS,
		data.PhoneDetail.Chipset,
		data.PhoneDetail.CPU,
		data.PhoneDetail.GPU,
		data.PhoneDetail.Dimensions,
		data.PhoneDetail.Weight,
		data.PhoneDetail.Build,
		data.PhoneDetail.SimInfo,
		data.PhoneDetail.NetTechnology,
		data.PhoneDetail.NetSpeed,
		data.PhoneDetail.G2,
		data.PhoneDetail.G3,
		data.PhoneDetail.G4,
		data.PhoneDetail.G5,
		data.PhoneDetail.GPS,
		data.PhoneDetail.NFC,
		data.PhoneDetail.Radio,
		data.PhoneDetail.Wlan,
		data.PhoneDetail.Bluetooth,
		data.PhoneDetail.USB,
		data.PhoneDetail.CardSlot,
	)
	if err != nil {
		return fmt.Errorf("insert phone_details error: %v", err)
	}
	err = pr.InsertBattery(data.Battery, prodID, tx)
	if err != nil {
		return err
	}
	err = pr.InsertDisplay(data.Display, prodID, tx)
	if err != nil {
		return err
	}
	err = pr.InsertMemory(data.Memory, prodID, tx)
	if err != nil {
		return err
	}
	err = pr.InsertSound(data.Sound, prodID, tx)
	if err != nil {
		return err
	}
	err = pr.InsertSensors(data.Sensors, prodID, tx)
	if err != nil {
		return err
	}
	err = pr.InsertCamera(data.Cameras.MainCamera, prodID, "MainCamera", tx)
	if err != nil {
		return err
	}
	err = pr.InsertCamera(data.Cameras.SelfieCamera, prodID, "SelfieCamera", tx)
	if err != nil {
		return err
	}
	err = pr.InsertColor(data.Colors, prodID, tx)
	if err != nil {
		return err
	}
	err = pr.InsertModels(data.Models, prodID, tx)
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
	query = `INSERT INTO camera_lenses(camera_id, megapixels,aperture,focal_length, sensor_size, pixel_size,other_features,zoom) VALUES($1,$2,$3,$4,$5,$6, $7, $8, $9)`
	for _, dat := range data.Lenses {
		_, err = tx.Exec(query, cid, dat.Megapixels, dat.Aperture, dat.FocalLength, dat.SensorSize, dat.PixelSize, pq.Array(dat.OtherFeatures), dat.Zoom)
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
func (pr *ProductRepo) InsertSensors(data models.Sensors, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO sensors(product_id,features) VALUES($1,$2)`

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
func (pr *ProductRepo) InsertMemory(data []models.Memory, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO memory_options(product_id,storage,ram) VALUES($1, $2,$3)`
	for _, dat := range data {
		_, err := tx.Exec(query, id, dat.Storage, dat.RAM)
		if err != nil {
			return err
		}
	}
	return nil
}
func (pr *ProductRepo) InsertDisplay(data models.Display, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO display(product_id,type,size,resolution,protection,aspect_ratio,hdr,refresh_rate, brightness_typical,brightness_hbm,ppi) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`

	_, err := tx.Exec(query, id, data.Type, data.Size, data.Resolution, data.Protection, data.AspectRatio, pq.Array(data.HDR), data.RefreshRate, data.BrightnessTypical, data.BrightnessHbm, pq.Array(data.PPI))
	if err != nil {
		return err
	}
	return nil
}
func (pr *ProductRepo) InsertBattery(data models.Battery, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO battery(product_id,technology , capacity) VALUES($1,$2,$3)`

	_, err := tx.Exec(query, id, data.Technology, data.Capacity)
	if err != nil {
		return err
	}
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
	var product models.Product
	if prodid == 0 {
		return nil, fmt.Errorf("Invalid data")
	}
	query := `SELECT id,name,image_url,category_id,brand_id FROM products WHERE id = $1; `

	err := pr.db.Get(&product, query, prodid)

	if err != nil {
		return nil, fmt.Errorf("Database error : %w", err)
	}

	brand, err := pr.brand.GetBrand(product.BrandID)
	if err != nil {
		return nil, err
	}
	cat, err := pr.cat.GetCategory(product.CategoryId)
	if err != nil {
		return nil, err
	}
	product.Category = cat.Name
	product.Brand = brand.Name

	productDetail, err := pr.psr.GetProductDetail(product)

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
