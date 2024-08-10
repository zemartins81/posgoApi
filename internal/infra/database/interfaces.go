package database

import "github.com/zemartins81/posgoApi/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByUser(email string) (*entity.User, error)
}
