package models

type Category struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description,omitempty" db:"description"`
}
