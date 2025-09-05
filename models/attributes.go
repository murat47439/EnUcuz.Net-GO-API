package models

import "database/sql"

type Attribute struct {
	ID        int           `json:"id" db:"id"`
	Name      string        `json:"name" db:"name"`
	DeletedAt *sql.NullTime `json:"-" db:"deleted_at"`
}
