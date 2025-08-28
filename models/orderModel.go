package models

import "github.com/shopspring/decimal"

type Order struct {
	ID        int             `json:"id" db:"id"`
	UserID    int             `json:"user_id" db:"user_id"`
	Total     decimal.Decimal `json:"total" db:"total"`
	Status    int             `json:"status,omitempty" db:"status"`
	AddressID int             `json:"address_id,omitempty" db:"address_id"`
}
type OrderDetail struct {
	ID        int             `json:"id" db:"id"`
	OrderID   int             `json:"order_id" db:"order_id"`
	Order     *Order          `json:"order,omitempty"`
	ProductID int             `json:"product_id" db:"product_id"`
	Product   *Product        `json:"product,omitempty"`
	Quantity  int             `json:"quantity" db:"quantity"`
	UnitPrice decimal.Decimal `json:"unit_price" db:"unit_price"`
	Discount  decimal.Decimal `json:"discount" db:"discount"`
	Total     decimal.Decimal `json:"total" db:"total"`
}
