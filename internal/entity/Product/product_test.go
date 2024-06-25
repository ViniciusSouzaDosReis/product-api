package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 1999.00)
	assert.Equal(t, nil, err)
	assert.NotNil(t, product)

	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 1999.00, product.Price)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.CreatedAt)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 1999.00)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Product 1", 0)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceisRequired, err)
}

func TestProductWhenInvalidPrice(t *testing.T) {
	product, err := NewProduct("Product 1", -10)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice, err)
}
