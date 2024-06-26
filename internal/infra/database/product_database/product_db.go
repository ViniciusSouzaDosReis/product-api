package product_database

import (
	"github.com/ViniciusSouzaDosReis/product-api/internal/entity/product"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *product.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]product.Product, error) {
	var producst []product.Product
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = p.DB.
			Limit(limit).
			Offset((page - 1) * limit).
			Order("created_at " + sort).
			Find(&producst).
			Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&producst).Error
	}
	return producst, err
}

func (p *Product) FindById(id string) (*product.Product, error) {
	var product product.Product
	if err := p.DB.First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *Product) Delete(id string) error {
	var product product.Product
	if err := p.DB.First(&product, "id = ?", id).Error; err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}

func (p *Product) Update(productToUpdate *product.Product) error {
	var productFound product.Product
	if err := p.DB.First(&productFound, "id = ?", productToUpdate.ID).Error; err != nil {
		return err
	}

	return p.DB.Save(productToUpdate).Error
}
