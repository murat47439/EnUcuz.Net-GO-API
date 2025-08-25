package models

type User struct {
	ID       int    `db:"id"  json:"id"`
	Email    string `db:"email" json:"email"`
	Phone    int    `db:"phone" json:"phone"`
	Name     string `db:"name" json:"name"`
	Surname  string `db:"surname" json:"surname"`
	Gender   int    `db:"gender" json:"gender"`
	Role     int    `db:"role" json:"role"`
	Password string `db:"password" json:"password"`
}
