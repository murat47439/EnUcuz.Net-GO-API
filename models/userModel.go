package models

type User struct {
	ID       int    `db:"id"  json:"id"`
	Email    string `db:"email" json:"email"`
	Phone    int    `db:"phone" json:"phone"`
	Name     string `db:"name" json:"name"`
	Surname  string `db:"surname" json:"surname"`
	Gender   int    `db:"gender" json:"gender,omitempty"`
	Role     int    `db:"role" json:"role,omitempty"`
	Password string `db:"password" json:"password"`
}
type UpdateUser struct {
	Token Token `json:"token"`
	User  User  `json:"user"`
}
