package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 100)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.Name)
	assert.NotEmpty(t, product.Price)
	assert.NotEmpty(t, product.CreatedAt)
	assert.Equal(t, product.Name, "Product 1")
	assert.Equal(t, product.Price, 100)
}

func TestProductWhenNameISRequired(t *testing.T) {
	product, err := NewProduct("", 100)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, err, ErrNameIsRequired)
}

func TestProductWhenPriceISRequired(t *testing.T) {
	product, err := NewProduct("Product 1", 0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, err, ErrPriceIsRequired)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Product 1", -1)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, err, ErrInvalidPrice)
}
