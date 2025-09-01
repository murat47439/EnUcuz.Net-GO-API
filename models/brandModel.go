package models

import "time"

type Brand struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
type Brands struct {
	Brand []Brand `json:"brands"`
}
