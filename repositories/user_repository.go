package repositories

import "app/models"

type UserRepository interface {
	FindUserByOpenid(openid string) error
	Create() error
	UpdateSessionKey(sessionKey string) error
}

type userRepository struct {
	repository
}

func NewUserRepository() UserRepository {
	return &userRepository{
		*NewRepository(),
	}
}

func (r *userRepository) FindUserByOpenid(openid string) error {

	return r.DB.Where("openid = ?", openid).First(&models.User{}).Error
}

func (r *userRepository) Create() error {
	return r.DB.Create(&models.User{}).Error
}

func (r *userRepository) UpdateSessionKey(sessionKey string) error {
	user := &models.User{
		SessionKey: sessionKey,
	}
	return r.DB.Save(&user).Error
}
