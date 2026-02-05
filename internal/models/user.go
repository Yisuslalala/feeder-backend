package models

type User struct {
	ID           int64 `json:"id"`
	Email        string `json:"email"`
	Name         string `json:name`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
}
