package database

import (
	"github.com/zemartins81/posgoApi/internal/entity"
	"gorm.io/gorm"
)

type USer struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *USer {
	return &USer{
		DB: db,
	}
}

func (u *USer) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *USer) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
