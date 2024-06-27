package interfaces

import (
	"github.com/ViniciusSouzaDosReis/product-api/internal/entity/product"
)

type ProductInterface interface {
	Create(product *product.Product) error
	FindAll(page, limit int, sort string) ([]product.Product, error)
	FindById(id string) (*product.Product, error)
	Delete(id string) error
	Update(product *product.Product) error
}
