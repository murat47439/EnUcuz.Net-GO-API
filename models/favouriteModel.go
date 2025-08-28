package models

import "time"

type Favori struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	ProductID int       `json:"product_id" db:"product_id"`
	Product   *Product  `json:"product"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"crated_at"`
}
