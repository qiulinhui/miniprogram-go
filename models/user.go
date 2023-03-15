package models

import (
	"app/app"
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
	return app.DB.Where("openid = ?", openid).First(&user).Error
}

func (user *User) Create() error {
	return app.DB.Create(&user).Error
}

func (user *User) UpdateSessionKey(sessionKey string) error {
	user.SessionKey = sessionKey
	return app.DB.Save(&user).Error
}
