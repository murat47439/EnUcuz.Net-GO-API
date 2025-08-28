package models

type UpdateUser struct {
	Token Token `json:"token"`
	User  User  `json:"user"`
}
