package repo

import (
	"Store-Dio/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type BrandsRepo struct {
	db *sqlx.DB
}

func NewBrandsRepo(db *sqlx.DB) *BrandsRepo {
	return &BrandsRepo{
		db: db,
	}
}

func (br *BrandsRepo) InsertBrandData(brands *models.Brands) error {
	tx, err := br.db.Beginx()

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

	err = br.InsertBrandsData(tx, brands)
	if err != nil {
		return err
	}
	return nil
}
func (br *BrandsRepo) InsertBrandsData(tx *sqlx.Tx, brands *models.Brands) error {
	for _, brand := range brands.Brand {
		if err := br.InsertBrand(tx, brand); err != nil {
			return err
		}
	}
	return nil
}
func (br *BrandsRepo) InsertBrand(tx *sqlx.Tx, brand models.Brand) error {
	if brand.ID == 0 || brand.Name == "" {
		return fmt.Errorf("Invalid data")
	}

	query := `INSERT INTO brands(id,name,created_at) values($1 ,$2, NOW())
		ON CONFLICT(name) DO UPDATE SET
		name = EXCLUDED.name
	`
	_, err := tx.Exec(query, brand.ID, brand.Name)

	if err != nil {
		return fmt.Errorf("Database error : %w", err)
	}
	return nil

}
