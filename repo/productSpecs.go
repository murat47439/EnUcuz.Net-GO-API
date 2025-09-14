package repo

import (
	"Store-Dio/models"
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
)

type ProductSpecsRepo struct {
	db *sqlx.DB
}

func NewProductSpecsRepo(db *sqlx.DB) *ProductSpecsRepo {
	return &ProductSpecsRepo{
		db: db,
	}
}
func (psr *ProductSpecsRepo) GetProductDetail(data models.Product) (*models.ProductDetail, error) {
	if data.ID == 0 {
		return nil, fmt.Errorf("Invalid data")
	}

	var productDetail models.ProductDetail
	var wg sync.WaitGroup
	wg.Add(9)

	// Değişkenler ve hata tutucular
	var phoneDetail models.PhoneDetail
	var battery models.Battery
	var display models.Display
	var memory []models.Memory
	var sound models.Sound
	var sensors models.Sensors
	var colors []string
	var modelsArr []string
	var cameras models.Cameras

	var phoneDetailErr, batteryErr, displayErr, memoryErr, soundErr, sensorErr, colorsErr, modelsErr, camerasErr error

	// Paralel goroutine’ler
	go func() { defer wg.Done(); phoneDetail, phoneDetailErr = psr.getPhoneDetail(data.ID) }()
	go func() { defer wg.Done(); battery, batteryErr = psr.getBattery(data.ID) }()
	go func() { defer wg.Done(); display, displayErr = psr.getDisplay(data.ID) }()
	go func() { defer wg.Done(); memory, memoryErr = psr.getMemory(data.ID) }()
	go func() { defer wg.Done(); sound, soundErr = psr.getSound(data.ID) }()
	go func() { defer wg.Done(); sensors, sensorErr = psr.getSensors(data.ID) }()
	go func() { defer wg.Done(); colors, colorsErr = psr.getColors(data.ID) }()
	go func() { defer wg.Done(); modelsArr, modelsErr = psr.getModels(data.ID) }()
	go func() { defer wg.Done(); cameras, camerasErr = psr.getCameras(data.ID) }()

	// Brand ve Category paralel

	// Tüm goroutine’lerin bitmesini bekle
	wg.Wait()

	// Hataları kontrol et
	if phoneDetailErr != nil {
		return nil, phoneDetailErr
	}
	if batteryErr != nil {
		return nil, batteryErr
	}
	if displayErr != nil {
		return nil, displayErr
	}
	if memoryErr != nil {
		return nil, memoryErr
	}
	if soundErr != nil {
		return nil, soundErr
	}
	if sensorErr != nil {
		return nil, sensorErr
	}
	if colorsErr != nil {
		return nil, colorsErr
	}
	if modelsErr != nil {
		return nil, modelsErr
	}
	if camerasErr != nil {
		return nil, camerasErr
	}

	// Tüm verileri ata
	productDetail.Product = data
	productDetail.PhoneDetail = phoneDetail
	productDetail.Battery = battery
	productDetail.Display = display
	productDetail.Memory = memory
	productDetail.Sound = sound
	productDetail.Sensors = sensors
	productDetail.Colors = colors
	productDetail.Models = modelsArr
	productDetail.Cameras = cameras

	return &productDetail, nil
}
func (psr *ProductSpecsRepo) getPhoneDetail(prodid int) (models.PhoneDetail, error) {
	if prodid == 0 {
		return models.PhoneDetail{}, fmt.Errorf("Invalid data")
	}
	query := `SELECT current_os,upgradable_to,chipset,cpu,gpu,dimensions,weight,build,sim_info,network_technology,network_speed,g2, g3, g4, g5,gps,nfc,radio,wlan,bluetooth,usb,card_slot FROM phone_details WHERE product_id = $1 `
	var detail models.PhoneDetail

	err := psr.db.Get(&detail, query, prodid)
	if err != nil {
		return models.PhoneDetail{}, err
	}
	return detail, nil
}
func (psr *ProductSpecsRepo) getBattery(prodid int) (models.Battery, error) {
	var battery models.Battery
	if prodid == 0 {
		return models.Battery{}, fmt.Errorf("Invalid data")
	}
	query := `SELECT technology,capacity FROM battery WHERE product_id = $1`

	err := psr.db.QueryRowx(query, prodid).Scan(&battery.Technology, battery.Capacity)
	if err != nil {
		return models.Battery{}, err
	}
	return battery, nil
}
func (psr *ProductSpecsRepo) getDisplay(prodid int) (models.Display, error) {
	var display models.Display
	if prodid == 0 {
		return models.Display{}, fmt.Errorf("Invalid data")
	}
	query := `SELECT type, size,resolution,protection,aspect_ratio,hdr,refresh_rate,brightness_typical,brightness_hbm,ppi FROM display WHERE product_id = $1`

	err := psr.db.Get(&display, query, prodid)

	if err != nil {
		return models.Display{}, err
	}
	return display, nil
}
func (psr *ProductSpecsRepo) getMemory(prodid int) ([]models.Memory, error) {
	var memory []models.Memory
	if prodid == 0 {
		return []models.Memory{}, fmt.Errorf("Invalid data")
	}
	query := `SELECT  storage,ram FROM memory_options WHERE product_id = $1`

	err := psr.db.Select(&memory, query, prodid)

	if err != nil {
		return []models.Memory{}, err
	}
	return memory, nil
}
func (psr *ProductSpecsRepo) getSound(prodid int) (models.Sound, error) {
	var sound models.Sound
	if prodid == 0 {
		return models.Sound{}, fmt.Errorf("Invalid data")
	}
	query := `SELECT loudspeaker,features FROM sound_specs WHERE product_id = $1`

	err := psr.db.Get(&sound, query, prodid)

	if err != nil {
		return models.Sound{}, err
	}
	return sound, nil
}
func (psr *ProductSpecsRepo) getSensors(prodid int) (models.Sensors, error) {
	var sensors models.Sensors
	if prodid == 0 {
		return models.Sensors{}, fmt.Errorf("Invalid data")
	}
	query := `SELECT features FROM sensors WHERE product_id = $1`

	err := psr.db.Get(&sensors, query, prodid)

	if err != nil {
		return models.Sensors{}, err
	}
	return sensors, nil
}
func (psr *ProductSpecsRepo) getCameras(prodid int) (models.Cameras, error) {
	// cameras tablosundan slice çek
	cameraRows := []struct {
		ID   int    `db:"id"`
		Type string `db:"type"`
	}{}

	err := psr.db.Select(&cameraRows, "SELECT id, type FROM cameras WHERE product_id=$1", prodid)
	if err != nil {
		return models.Cameras{}, err
	}

	// range ile dön
	var cams models.Cameras
	for _, c := range cameraRows {
		var cam models.Camera

		// Lenses
		err = psr.db.Select(&cam.Lenses, `
		SELECT megapixels, aperture, focal_length, sensor_size, type, pixel_size, other_features, zoom 
		FROM camera_lenses WHERE camera_id=$1`, c.ID)
		if err != nil {
			return cams, err
		}

		// Features
		err = psr.db.Select(&cam.Features, `
		SELECT features as spec FROM camera_features WHERE camera_id=$1`, c.ID)
		if err != nil {
			return cams, err
		}

		// Video
		err = psr.db.Select(&cam.Video, `
		SELECT video_spec as video FROM camera_video WHERE camera_id=$1`, c.ID)
		if err != nil {
			return cams, err
		}

		// Kamerayı struct içine ata
		if c.Type == "MainCamera" {
			cams.MainCamera = cam
		} else if c.Type == "SelfieCamera" {
			cams.SelfieCamera = cam
		}
	}

	return cams, nil

}
func (psr *ProductSpecsRepo) getColors(prodid int) ([]string, error) {
	if prodid == 0 {
		return nil, fmt.Errorf("Invalid data")
	}

	var colors []string
	query := `SELECT color FROM colors WHERE product_id = $1`

	err := psr.db.Select(&colors, query, prodid)
	if err != nil {
		return nil, err
	}

	return colors, nil
}

func (psr *ProductSpecsRepo) getModels(prodid int) ([]string, error) {
	var models []string
	if prodid == 0 {
		return nil, fmt.Errorf("Invalid data")
	}

	query := `
        SELECT model
        FROM models
        WHERE product_id = $1
    `

	rows, err := psr.db.Queryx(query, prodid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var model string
		err := rows.Scan(&model)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}

	return models, nil
}
