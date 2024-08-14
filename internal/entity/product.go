package entity

import (
	"errors"
	"github.com/zemartins81/posgoApi/pkg/entity"
	"time"
)

var (
	ErrIDIsRequired    = errors.New("ID is required")
	ErrInvalidID       = errors.New("Invalid ID")
	ErrNameIsRequired  = errors.New("Name is required")
	ErrPriceIsRequired = errors.New("Price is required")
	ErrInvalidPrice    = errors.New("Price must be greater than 0")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	if err := product.Validate(); err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price == 0 {
		return ErrPriceIsRequired
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}
