package models

type Error struct {
	Message string `json:"message"`
}

type User struct {
	About    string `json:"about"    db:"about"`
	Email    string `json:"email"    db:"email"`
	Fullname string `json:"fullname" db:"fullname"`
	Nickname string `json:"nickname" db:"nickname"`
}

type UserUpdate struct {
	About    string `json:"about"    db:"about"`
	Email    string `json:"email"    db:"email"`
	Fullname string `json:"fullname" db:"fullname"`
}
