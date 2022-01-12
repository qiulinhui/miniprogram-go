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

func (user *User) FindUserByOpenid(openid string) error {
	return DB.Where("openid = ?", openid).First(&user).Error
}

func (user *User) Create() error {
	return DB.Create(&user).Error
}

func (user *User) UpdateSessionKey(sessionKey string) error {
	user.SessionKey = sessionKey
	return DB.Save(&user).Error
}
