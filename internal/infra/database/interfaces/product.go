package interfaces

import (
	"github.com/ViniciusSouzaDosReis/product-api/internal/entity"
)

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Delete(id string) error
	Update(product *entity.Product) error
}
