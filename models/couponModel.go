package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Coupon struct {
	ID            int             `json:"id" db:"id"`
	Code          string          `json:"code" db:"code"`
	CategoryID    int             `json:"category_id,omitempty" db:"category_id"`
	Category      *Category       `json:"category,omitempty"`
	ProductID     int             `json:"product_id,omitempty" db:"product_id"`
	Product       *Product        `json:"product,omitempty"`
	Description   string          `json:"description" db:"description"`
	DiscountType  int             `json:"discount_type" db:"discount_type"`
	Discount      decimal.Decimal `json:"discount" db:"discount"`
	MinimumAmount decimal.Decimal `json:"minimum_amount" db:"minimum_amount"`
	StartDate     time.Time       `json:"start_date" db:"start_date"`
	EndDate       time.Time       `json:"end_date" db:"end_date"`
	Active        int             `json:"active" db:"active"`
	Count         int             `json:"count" db:"count"`
	Max           int             `json:"max" db:"max"`
	CreatedAt     time.Time       `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at,omitempty" db:"updated_at"`
}
