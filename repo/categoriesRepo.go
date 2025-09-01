package repo

import (
	"Store-Dio/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type CategoriesRepo struct {
	db *sqlx.DB
}

func NewCategoriesRepo(db *sqlx.DB) *CategoriesRepo {
	return &CategoriesRepo{
		db: db,
	}
}

func (cr *CategoriesRepo) InsertCategory(tx *sqlx.Tx, cat models.Category) error {
	query := `
		INSERT INTO categories (id, name, parent_id, created_at) VALUES($1, $2, $3, NOW())
		ON CONFLICT (id) DO UPDATE SET
		name = EXCLUDED.name,
		parent_id = EXCLUDED.parent_id
	`
	_, err := tx.Exec(query, cat.ID, cat.Name, cat.ParentID)
	if err != nil {
		return fmt.Errorf("InsertCategoriesData Error")
	}
	return nil
}
func (cr *CategoriesRepo) InsertCategoryDataRecursive(tx *sqlx.Tx, cat models.Category) error {
	if err := cr.InsertCategory(tx, cat); err != nil {
		return err
	}

	for _, sub := range cat.SubCategory {
		if err := cr.InsertCategoryDataRecursive(tx, sub); err != nil {
			return err
		}
	}
	return nil
}

func (cr *CategoriesRepo) InsertCategoriesRecursive(cat models.Category) (err error) {
	tx, err := cr.db.Beginx()
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

	err = cr.InsertCategoryDataRecursive(tx, cat)
	if err != nil {
		return err
	}

	return nil
}
func (cr *CategoriesRepo) GetAllCategoriesID() (models.Categories, error) {
	var categories models.Categories
	query := `SELECT id FROM categories`

	err := cr.db.Select(&categories, query)

	if err != nil {
		return nil, fmt.Errorf("Categories not found. Error : ", err.Error())
	}
	return categories, nil
}
