package models

import (
	"database/sql"
)

type Product struct {
	ID          int                `json:"id" db:"id"`
	Name        string             `json:"name" db:"name"`
	Description string             `json:"description,omitempty" db:"description"`
	Stock       int                `json:"stock,omitempty" db:"stock"`
	Brand       string             `json:"brand_name,omitempty" db:"brand_name"`
	BrandID     int                `json:"brand_id" db:"brand_id"`
	ImageUrl    *string            `json:"image_url" db:"image_url"`
	SellerID    int                `json:"seller_id" db:"seller_id"`
	SellerName  string             `json:"seller_name,omitempty" db:"seller_name"`
	CategoryId  int                `json:"category_id" db:"category_id"`
	Category    string             `json:"category_name,omitempty" db:"category_name"`
	Released    *string            `json:"released,omitempty" db:"released"`
	Announced   *string            `json:"announced,omitempty" db:"announced"`
	Status      *string            `json:"status,omitempty" db:"status"`
	CreatedAt   sql.NullTime       `json:"created_at" db:"created_at"`
	UpdatedAt   sql.NullTime       `json:"updated_at" db:"updated_at"`
	DeletedAt   sql.NullTime       `json:"-" db:"deleted_at"`
	Attributes  []ProductAttribute `json:"attributes,omitempty" db:"attributes"`
}
