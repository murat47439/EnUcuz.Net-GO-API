package repo

import (
	"Store-Dio/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AttributeRepo struct {
	db *sqlx.DB
}

func NewAttributeRepo(db *sqlx.DB) *AttributeRepo {
	return &AttributeRepo{
		db: db,
	}
}
func (ar *AttributeRepo) AddAttribute(data *models.Attribute) error {
	tx, err := ar.db.Beginx()

	if err != nil {
		return fmt.Errorf("TX Error : %s", err.Error())
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

	query := `INSERT INTO attributes (name) VALUES($1)`

	_, err = tx.Exec(query, data.Name)

	if err != nil {
		return fmt.Errorf("Database error : %s", err.Error())
	}
	return nil
}
func (ar *AttributeRepo) UpdateAttribute(data *models.Attribute) error {
	tx, err := ar.db.Beginx()

	if err != nil {
		return fmt.Errorf("TX Error : %s", err.Error())
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

	exists, err := ar.CheckAttribute(data.ID, tx)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Invalid data")
	}
	query := `UPDATE attributes SET name = $1 WHERE id = $2`

	_, err = tx.Exec(query, data.Name, data.ID)
	if err != nil {
		return fmt.Errorf("Database error : %s", err.Error())
	}
	return nil
}
func (ar *AttributeRepo) DeleteAttribute(data *models.Attribute) error {
	tx, err := ar.db.Beginx()

	if err != nil {
		return fmt.Errorf("TX Error : %s", err.Error())
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
	exists, err := ar.CheckAttribute(data.ID, tx)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Invalid data")
	}
	query := `UPDATE attributes SET deleted_at = NOW() WHERE id = $1`

	_, err = tx.Exec(query, data.ID)
	if err != nil {
		return fmt.Errorf("Database error : %s", err.Error())
	}
	return nil
}
func (ar *AttributeRepo) GetAttribute(id int) (*models.Attribute, error) {
	var data *models.Attribute
	query := `SELECT id, name FROM attributes WHERE id = $1`

	err := ar.db.Get(&data, query, id)

	if err != nil {
		return nil, fmt.Errorf("Database error : %s", err.Error())
	}
	return data, nil
}
func (ar *AttributeRepo) GetAttributes(page int, search string) ([]*models.Attribute, error) {
	var attributes []*models.Attribute
	limit := 50
	offset := (page - 1) * 50
	query := `SELECT id, name FROM attributes WHERE name LIKE $1 LIMIT $2 OFFSET $3`

	rows, err := ar.db.Queryx(query, "%"+search+"%", limit, offset)
	if err != nil {
		return nil, fmt.Errorf("Database error: %s", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var attribute models.Attribute

		if err = rows.StructScan(&attribute); err != nil {
			return nil, fmt.Errorf("Rows error : ", err.Error())
		}
		attributes = append(attributes, &attribute)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Rows error : ", err.Error())
	}

	return attributes, nil
}
func (ar *AttributeRepo) CheckAttribute(id int, tx *sqlx.Tx) (bool, error) {
	if id == 0 {
		return false, fmt.Errorf("Invalid data")
	}

	query := `SELECT EXISTS(SELECT 1 FROM attributes WHERE id = $1)`
	var exists bool
	err := tx.Get(&exists, query, id)

	if err != nil {
		return false, fmt.Errorf("Database error : %s", err.Error())
	}
	return exists, nil
}
