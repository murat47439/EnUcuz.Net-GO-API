package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Gender   int    `json:"gender"`
	Role     int    `json:"role"`
	Password string `json:"password"`
}
