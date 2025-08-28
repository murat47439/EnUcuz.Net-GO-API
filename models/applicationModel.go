package models

import "time"

type Applications struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	User      *User     `json:"user,omitempty"`
	StoreName string    `json:"store_name" db:"store_name"`
	StoreDesc string    `json:"store_desc" db:"store_desc"`
	Email     string    `json:"email" db:"email"`
	Phone     int       `json:"phone" db:"phone"`
	TaxNumber string    `json:"tax_number,omitempty" db:"tax_number"`
	TaxOffice string    `json:"tax_office,omitempty" db:"tax_office"`
	Status    int       `json:"status,omitempty" db:"status"`
	AdminNote string    `json:"admin_note,omitempty" db:"admin_note"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
