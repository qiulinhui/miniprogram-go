package models

import (
	"time"
)

type User struct {
	ID         uint
	Nickname   string
	Openid     string
	Avatar     string
	Phone      string
	SessionKey string `json:"-"`
	CreatedAt  time.Time
}
