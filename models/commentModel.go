package models

type Comment struct {
	ID        int      `json:"id" db:"id"`
	UserID    int      `json:"user_id" db:"user_id"`
	ProductID int      `json:"product_id" db:"product_id"`
	Product   *Product `json:"product,omitempty"`
	Content   string   `json:"content" db:"content"`
	Rating    int      `json:"rating" db:"rating"`
	Status    int      `json:"status,omitempty" db:"status"`
}
