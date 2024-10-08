package database

import "github.com/zemartins81/posgoApi/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) (products []entity.Product, err error)
	FindById(id string) (product *entity.Product, err error)
	Update(product *entity.Product) error
	Delete(id string) error
}
