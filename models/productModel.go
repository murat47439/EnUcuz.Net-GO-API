package models

import (
	"time"
)

type Product struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description,omitempty" db:"description"`
	Stock       int       `json:"stock,omitempty" db:"stock"`
	ImageUrl    string    `json:"image_url" db:"image_url"`
	BrandID     int       `json:"brand_id" db:"brand_id"`
	CategoryID  int       `json:"category_id" db:"category_id"`
	Category    *Category `json:"category,omitempty"`
	StoreID     int       `json:"store_id" db:"store_id"`
	Store       *Store    `json:"store,omitempty"`
	CreatedAt   time.Time `json:"crated_at,omitempty" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
type GetProducts struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset,omitempty"`
	Page   int    `json:"page"`
	SortBy string `json:"sort_by,omitempty"`
	Order  string `json:"order,omitempty"`
}
