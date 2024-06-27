package product_database

import (
	"github.com/ViniciusSouzaDosReis/product-api/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProductDB(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var producst []entity.Product
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

func (p *Product) FindById(id string) (*entity.Product, error) {
	var product entity.Product
	if err := p.DB.First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *Product) Delete(id string) error {
	var product entity.Product
	if err := p.DB.First(&product, "id = ?", id).Error; err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}

func (p *Product) Update(productToUpdate *entity.Product) error {
	var productFound entity.Product
	if err := p.DB.First(&productFound, "id = ?", productToUpdate.ID).Error; err != nil {
		return err
	}

	return p.DB.Save(productToUpdate).Error
}
