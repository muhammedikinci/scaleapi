package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

type Role int

const (
	ContentManager Role = iota
	BasicUser
)
