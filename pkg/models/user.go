package models

type User struct {
	ID             int     `json:"id"`
	Username       string  `json:"username"`
	Password       string  `json:"password"`
	Role           Role    `json:"role"`
	MovieFavorites []Movie `gorm:"many2many:user_movies"`
	SerieFavorites []Serie `gorm:"many2many:user_series"`
}

type Role int

const (
	ContentManager Role = iota
	BasicUser
)

type Favorite struct {
	Movies []Movie `json:"movies"`
	Series []Serie `json:"series"`
}
