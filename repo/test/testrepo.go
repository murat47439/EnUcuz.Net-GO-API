package test

import (
	"Store-Dio/models/testm"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type DataRepo struct {
	db *sqlx.DB
}

func NewDataRepo(db *sqlx.DB) *DataRepo {
	return &DataRepo{
		db: db,
	}
}

func (dr *DataRepo) InsertData(data *testm.Items) error {
	if data == nil {
		return fmt.Errorf("Invalid data")
	}
	tx, err := dr.db.Beginx()
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
	for _, dat := range data.Products {
		if err := dr.InsertProducts(dat, tx); err != nil {
			return fmt.Errorf("failed to insert product %d: %w", dat.ID, err)
		}
	}

	return nil

}
func (dr *DataRepo) InsertProducts(data testm.Product, tx *sqlx.Tx) error {
	err := dr.InsertBrands(data.Brand, tx)
	if err != nil {
		return err
	}
	query := `INSERT INTO products(id,name,category_id,brand_id) VALUES($1,$2,$3, $4)`

	_, err = tx.Exec(query, data.ID, data.Name, 3719, data.Brand.ID)
	if err != nil {
		return err
	}
	err = dr.InsertBattery(data.Battery, data.ID, tx)
	if err != nil {
		return err
	}
	err = dr.InsertPlatform(data.Platform, data.ID, tx)
	if err != nil {
		return err
	}
	err = dr.InsertNetwork(data.Network, data.ID, tx)
	if err != nil {
		return err
	}
	err = dr.InsertDisplay(data.Display, data.ID, tx)
	if err != nil {
		return err
	}
	err = dr.InsertLaunch(data.Launch, data.ID, tx)
	if err != nil {
		return err
	}
	err = dr.InsertBody(data.Body, data.ID, tx)
	if err != nil {
		return err
	}
	err = dr.InsertMemory(data.Memory, data.ID, tx)
	if err != nil {
		return err
	}
	err = dr.InsertSound(data.Sound, data.ID, tx)
	if err != nil {
		return err
	}
	err = dr.InsertComms(data.Comms, data.ID, tx)
	if err != nil {
		return err
	}
	err = dr.InsertFeatures(data.Features, data.ID, tx)
	if err != nil {
		return err
	}
	err = dr.InsertCamera(data.Cameras.MainCamera, data.ID, "MainCamera", tx)
	if err != nil {
		return err
	}
	err = dr.InsertCamera(data.Cameras.SelfieCamera, data.ID, "SelfieCamera", tx)
	if err != nil {
		return err
	}
	err = dr.InsertColor(data.Colors, data.ID, tx)
	if err != nil {
		return err
	}
	err = dr.InsertModels(data.Models, data.ID, tx)
	if err != nil {
		return err
	}

	return nil
}
func (dr *DataRepo) InsertColor(data []string, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO product_colors(product_id, color) VALUES($1, $2)`

	for _, color := range data {
		_, err := tx.Exec(query, id, color)
		if err != nil {
			return err
		}
	}

	return nil
}
func (dr *DataRepo) InsertModels(data []string, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO product_models(product_id, model) VALUES($1, $2)`

	for _, model := range data {
		_, err := tx.Exec(query, id, model)
		if err != nil {
			return err
		}
	}

	return nil
}
func (dr *DataRepo) InsertCamera(data testm.Camera, id int, role string, tx *sqlx.Tx) error {
	query := `INSERT INTO cameras(product_id, camera_type, camera_specs, features, video, camera_role) VALUES($1,$2,$3,$4,$5, $6)`

	_, err := tx.Exec(query, id, data.Type, pq.Array(data.CameraSpecs), pq.Array(data.Features), pq.Array(data.Video), role)

	if err != nil {
		return err
	}
	return nil
}
func (dr *DataRepo) InsertFeatures(data testm.Features, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO feature_specs(product_id,sensors) VALUES($1,$2)`

	_, err := tx.Exec(query, id, data.Sensors)
	if err != nil {
		return err
	}
	return nil
}
func (dr *DataRepo) InsertComms(data testm.Comms, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO comms_specs(product_id,wlan,bluetooth,positioning,nfc,radio,usb)`

	_, err := tx.Exec(query, id, data.WLAN, data.Bluetooth, data.Positioning, data.NFC, data.Radio, data.USB)
	if err != nil {
		return err
	}
	return nil
}
func (dr *DataRepo) InsertSound(data testm.Sound, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO sound_specs(product_id,loudspeaker) VALUES($1,$2)`

	_, err := tx.Exec(query, id, data.Loudspeaker)
	if err != nil {
		return err
	}
	return nil
}
func (dr *DataRepo) InsertMemory(data testm.Memory, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO memory_specs(product_id,card_slot,internal) VALUES($1, $2,$3)`

	_, err := tx.Exec(query, id, data.CardSlot, data.Internal)
	if err != nil {
		return err
	}
	return nil
}
func (dr *DataRepo) InsertBody(data testm.Body, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO body_specs(product_id, dimensions, weight,build,sim) VALUES($1,$2,$3,$4,$5)`

	_, err := tx.Exec(query, id, data.Dimensions, data.Weight, data.Build, data.SIM)
	if err != nil {
		return err
	}
	return nil
}
func (dr *DataRepo) InsertLaunch(data testm.Launch, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO launch_specs(product_id,announced,released,status) VALUES($1, $2,$3,$4)`

	_, err := tx.Exec(query, id, data.Announced, data.Released, data.Status)
	if err != nil {
		return err
	}
	return nil
}
func (dr *DataRepo) InsertDisplay(data testm.Display, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO display_specs(product_id,type,size,resolution,protection) VALUES($1,$2,$3,$4,$5)`

	_, err := tx.Exec(query, id, data.Type, data.Size, data.Resolution, data.Protection)
	if err != nil {
		return err
	}
	return nil
}
func (dr *DataRepo) InsertNetwork(data testm.Network, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO network_specs(product_id, technology, speed, g2 , g3, g4 , g5) VALUES($1,$2,$3,$4,$5,$6,$7)`

	_, err := tx.Exec(query, id, data.Technology, data.Speed, data.G2, data.G3, data.G4, data.G5)
	if err != nil {
		return err
	}
	return nil
}
func (dr *DataRepo) InsertPlatform(data testm.Platform, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO platform_specs(product_id,os,chipset,cpu,gpu) VALUES($1,$2,$3,$4,$5) `

	_, err := tx.Exec(query, id, data.OS, data.Chipset, data.CPU, data.GPU)

	if err != nil {
		return err
	}
	return nil
}
func (dr *DataRepo) InsertBattery(data testm.Battery, id int, tx *sqlx.Tx) error {
	query := `INSERT INTO battery_specs(product_id, type, charging) VALUES($1,$2,$3)`

	_, err := tx.Exec(query, id, data.Type, pq.Array(data.Charging))
	if err != nil {
		return err
	}
	return nil
}
func (dr *DataRepo) InsertBrands(data testm.Brand, tx *sqlx.Tx) error {
	exists, err := dr.ExistsData(data.ID, "brands", tx)

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
func (dr *DataRepo) ExistsData(id int, table string, tx *sqlx.Tx) (bool, error) {
	if id == 0 || table == "" {
		return false, fmt.Errorf("Invalid data")
	}
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE id = $1)`, table)
	var exists bool
	err := tx.QueryRow(query, id).Scan(&exists)

	if err != nil {
		return false, err
	}
	return exists, nil
}
