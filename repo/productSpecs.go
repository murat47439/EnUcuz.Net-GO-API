package repo

import (
	"Store-Dio/models"
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ProductSpecsRepo struct {
	db *sqlx.DB
}

func NewProductSpecsRepo(db *sqlx.DB) *ProductSpecsRepo {
	return &ProductSpecsRepo{
		db: db,
	}
}
func (psr *ProductSpecsRepo) GetProductDetail(data *models.Product) (*models.ProductDetail, error) {

	if data.ID == 0 {
		return nil, fmt.Errorf("Invalid data")
	}

	brepo := NewBrandsRepo(psr.db)

	brand, err := brepo.GetBrand(data.Brand)

	if err != nil {
		return nil, fmt.Errorf("Brand error : %w", err)
	}

	crepo := NewCategoriesRepo(psr.db)

	category, err := crepo.GetCategory(data.CategoryId)

	if err != nil {
		return nil, fmt.Errorf("Category error : %w", err)
	}

	var productDetail models.ProductDetail
	var wg sync.WaitGroup

	// WaitGroup sayısı: 14 alan
	wg.Add(14)

	// Değişkenler ve hata tutucular
	var battery models.Battery
	var platform models.Platform
	var network models.Network
	var display models.Display
	var launch models.Launch
	var body models.Body
	var memory models.Memory
	var sound models.Sound
	var comms models.Comms
	var features models.Features
	var colors []string
	var modelsArr []string
	var cameras models.Cameras

	var batteryErr, platformErr, networkErr, displayErr, launchErr, bodyErr, memoryErr, soundErr, commsErr, featuresErr, colorsErr, modelsErr, camerasErr error

	// Paralel goroutine’ler
	go func() { defer wg.Done(); battery, batteryErr = psr.getBattery(data.ID) }()
	go func() { defer wg.Done(); platform, platformErr = psr.getPlatform(data.ID) }()
	go func() { defer wg.Done(); network, networkErr = psr.getNetwork(data.ID) }()
	go func() { defer wg.Done(); display, displayErr = psr.getDisplay(data.ID) }()
	go func() { defer wg.Done(); launch, launchErr = psr.getLaunch(data.ID) }()
	go func() { defer wg.Done(); body, bodyErr = psr.getBody(data.ID) }()
	go func() { defer wg.Done(); memory, memoryErr = psr.getMemory(data.ID) }()
	go func() { defer wg.Done(); sound, soundErr = psr.getSound(data.ID) }()
	go func() { defer wg.Done(); comms, commsErr = psr.getComms(data.ID) }()
	go func() { defer wg.Done(); features, featuresErr = psr.getFeatures(data.ID) }()
	go func() { defer wg.Done(); colors, colorsErr = psr.getColors(data.ID) }()
	go func() { defer wg.Done(); modelsArr, modelsErr = psr.getModels(data.ID) }()
	go func() { defer wg.Done(); cameras, camerasErr = psr.getCameras(data.ID) }()

	// Brand ve Category paralel

	// Tüm goroutine’lerin bitmesini bekle
	wg.Wait()

	// Hataları kontrol et
	if batteryErr != nil {
		return nil, batteryErr
	}
	if platformErr != nil {
		return nil, platformErr
	}
	if networkErr != nil {
		return nil, networkErr
	}
	if displayErr != nil {
		return nil, displayErr
	}
	if launchErr != nil {
		return nil, launchErr
	}
	if bodyErr != nil {
		return nil, bodyErr
	}
	if memoryErr != nil {
		return nil, memoryErr
	}
	if soundErr != nil {
		return nil, soundErr
	}
	if commsErr != nil {
		return nil, commsErr
	}
	if featuresErr != nil {
		return nil, featuresErr
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
	productDetail.ID = data.ID
	productDetail.Name = data.Name
	productDetail.ImageUrl = data.ImageUrl
	productDetail.Brand = brand
	productDetail.Category = category
	productDetail.Battery = battery
	productDetail.Platform = platform
	productDetail.Network = network
	productDetail.Display = display
	productDetail.Launch = launch
	productDetail.Body = body
	productDetail.Memory = memory
	productDetail.Sound = sound
	productDetail.Comms = comms
	productDetail.Features = features
	productDetail.Colors = colors
	productDetail.Models = modelsArr
	productDetail.Cameras = cameras
	productDetail.CreatedAt = data.CreatedAt
	productDetail.UpdatedAt = data.UpdatedAt

	return &productDetail, nil
}
func (psr *ProductSpecsRepo) getBattery(prodid int) (models.Battery, error) {
	var battery models.Battery
	if prodid == 0 {
		return models.Battery{}, fmt.Errorf("Invalid data")
	}
	query := `SELECT type,charging FROM battery_specs WHERE product_id = $1`

	err := psr.db.QueryRowx(query, prodid).Scan(&battery.Type, pq.Array(&battery.Charging))
	if err != nil {
		return models.Battery{}, err
	}
	return battery, nil
}
func (psr *ProductSpecsRepo) getPlatform(prodid int) (models.Platform, error) {
	var platform models.Platform
	if prodid == 0 {
		return models.Platform{}, fmt.Errorf("Invalid data")
	}
	query := `SELECT os, chipset,cpu,gpu FROM platform_specs WHERE product_id = $1`

	err := psr.db.Get(&platform, query, prodid)

	if err != nil {
		return models.Platform{}, err
	}
	return platform, nil
}
func (psr *ProductSpecsRepo) getNetwork(prodid int) (models.Network, error) {
	var network models.Network
	if prodid == 0 {
		return models.Network{}, fmt.Errorf("Invalid data")
	}
	query := `SELECT technology, speed,g2,g3, g4, g5 FROM network_specs WHERE product_id = $1`

	err := psr.db.Get(&network, query, prodid)

	if err != nil {
		return models.Network{}, err
	}
	return network, nil
}
func (psr *ProductSpecsRepo) getDisplay(prodid int) (models.Display, error) {
	var display models.Display
	if prodid == 0 {
		return models.Display{}, fmt.Errorf("Invalid data")
	}
	query := `SELECT type, size,resolution,protection FROM display_specs WHERE product_id = $1`

	err := psr.db.Get(&display, query, prodid)

	if err != nil {
		return models.Display{}, err
	}
	return display, nil
}
func (psr *ProductSpecsRepo) getLaunch(prodid int) (models.Launch, error) {
	var launch models.Launch
	if prodid == 0 {
		return models.Launch{}, fmt.Errorf("Invalid data")
	}
	query := `SELECT announced, released,status FROM launch_specs WHERE product_id = $1`

	err := psr.db.Get(&launch, query, prodid)

	if err != nil {
		return models.Launch{}, err
	}
	return launch, nil
}
func (psr *ProductSpecsRepo) getBody(prodid int) (models.Body, error) {
	var body models.Body
	if prodid == 0 {
		return models.Body{}, fmt.Errorf("Invalid data")
	}
	query := `SELECT dimensions, weight,build, sim FROM body_specs WHERE product_id = $1`

	err := psr.db.Get(&body, query, prodid)

	if err != nil {
		return models.Body{}, err
	}
	return body, nil
}
func (psr *ProductSpecsRepo) getMemory(prodid int) (models.Memory, error) {
	var memory models.Memory
	if prodid == 0 {
		return models.Memory{}, fmt.Errorf("Invalid data")
	}
	query := `SELECT card_slot, internal FROM memory_specs WHERE product_id = $1`

	err := psr.db.Get(&memory, query, prodid)

	if err != nil {
		return models.Memory{}, err
	}
	return memory, nil
}
func (psr *ProductSpecsRepo) getSound(prodid int) (models.Sound, error) {
	var sound models.Sound
	if prodid == 0 {
		return models.Sound{}, fmt.Errorf("Invalid data")
	}
	query := `SELECT loudspeaker FROM sound_specs WHERE product_id = $1`

	err := psr.db.Get(&sound, query, prodid)

	if err != nil {
		return models.Sound{}, err
	}
	return sound, nil
}
func (psr *ProductSpecsRepo) getComms(prodid int) (models.Comms, error) {
	var comms models.Comms
	if prodid == 0 {
		return models.Comms{}, fmt.Errorf("Invalid data")
	}
	query := `SELECT wlan, bluetooth,positioning,nfc,radio,usb FROM comms_specs WHERE product_id = $1`

	err := psr.db.Get(&comms, query, prodid)

	if err != nil {
		return models.Comms{}, err
	}
	return comms, nil
}
func (psr *ProductSpecsRepo) getFeatures(prodid int) (models.Features, error) {
	var features models.Features
	if prodid == 0 {
		return models.Features{}, fmt.Errorf("Invalid data")
	}
	query := `SELECT sensors FROM feature_specs WHERE product_id = $1`

	err := psr.db.Get(&features, query, prodid)

	if err != nil {
		return models.Features{}, err
	}
	return features, nil
}
func (psr *ProductSpecsRepo) getCameras(prodid int) (models.Cameras, error) {
	var cameras models.Cameras
	if prodid == 0 {
		return cameras, fmt.Errorf("Invalid data")
	}

	query := `
        SELECT camera_role, camera_type, camera_specs, features, video
        FROM cameras
        WHERE product_id = $1
    `

	rows, err := psr.db.Queryx(query, prodid)
	if err != nil {
		return cameras, err
	}
	defer rows.Close()

	for rows.Next() {
		var role string
		var cam models.Camera
		err := rows.Scan(&role, &cam.Type, pq.Array(&cam.CameraSpecs), pq.Array(&cam.Features), pq.Array(&cam.Video))
		if err != nil {
			return cameras, err
		}

		switch role {
		case "MainCamera":
			cameras.MainCamera = cam
		case "SelfieCamera":
			cameras.SelfieCamera = cam
		}
	}

	return cameras, nil
}
func (psr *ProductSpecsRepo) getColors(prodid int) ([]string, error) {
	var colors []string
	if prodid == 0 {
		return nil, fmt.Errorf("Invalid data")
	}

	query := `
        SELECT color
        FROM product_colors
        WHERE product_id = $1
    `

	rows, err := psr.db.Queryx(query, prodid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var color string
		err := rows.Scan(&color)
		if err != nil {
			return nil, err
		}
		colors = append(colors, color)
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
        FROM product_models
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
