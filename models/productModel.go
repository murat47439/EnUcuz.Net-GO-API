package models

import "github.com/shopspring/decimal"

type Product struct {
	ID          int             `json:"id" db:"id"`
	Name        string          `json:"name" db:"name"`
	Description string          `json:"description,omitempty" db:"description"`
	Price       decimal.Decimal `json:"price" db:"price"`
	Stock       int             `json:"stock,omitempty" db:"stock"`
	ImageUrl    string          `json:"image_url" db:"image_url"`
	CategoryID  int             `json:"category_id" db:"category_id"`
	Category    *Category       `json:"category,omitempty"`
	StoreID     int             `json:"seller_id" db:"seller_id"`
	Store       *Store          `json:"seller,omitempty"`
}
type GetProducts struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset,omitempty"`
	Page   int    `json:"page"`
	SortBy string `json:"sort_by,omitempty"`
	Order  string `json:"order,omitempty"`
}
