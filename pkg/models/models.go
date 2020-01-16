package models

type NewUser struct {
	About    string `json:"about"    db:"about"`
	Email    string `json:"email"    db:"email"`
	Fullname string `json:"fullname" db:"fullname"`
}

type User struct {
	About    string `json:"about"    db:"about"`
	Email    string `json:"email"    db:"email"`
	Fullname string `json:"fullname" db:"fullname"`
	Nickname string `json:"nickname" db:"nickname"`
}
