package models

type Address struct {
	Title         string `json:"title" db:"title"`
	ID            int    `json:"id" db:"id"`
	UserID        int    `json:"user_id" db:"user_id"`
	NameSurname   string `json:"fullname" db:"fullname"`
	Phone         int    `json:"phone" db:"phone"`
	AddressLine   string `json:"address_line" db:"address_line"`
	City          string `json:"city" db:"city"`
	District      string `json:"district" db:"district"`
	Neighbourhood string `json:"neighbourhood" db:"neighbourhood"`
	ZipCode       string `json:"zip_code" db:"zip_code"`
	Street        string `json:"street" db:"street"`
	IsDefault     int    `json:"is_default,omitempty" db:"is_default"`
}
