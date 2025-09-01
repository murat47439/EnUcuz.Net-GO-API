package models

import "time"

type Category struct {
	ID          int        `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	ParentID    *int       `json:"parentId,omitempty" db:"parent_id"`
	SubCategory []Category `json:"subCategories,omitempty"`
	CreatedAt   time.Time  `json:"created_at,omitempty" db:"created_at"`
}
type Categories []Category
