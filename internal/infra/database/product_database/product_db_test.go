package product_database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ViniciusSouzaDosReis/product-api/internal/entity/product"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Equal(t, nil, err)
	db.AutoMigrate(&product.Product{})

	newProduct, _ := product.NewProduct("Product 1", 10.0)
	productDB := NewProduct(db)

	err = productDB.Create(newProduct)
	assert.Equal(t, nil, err)

	var productFound product.Product
	err = db.First(&productFound, "id = ?", newProduct.ID).Error
	assert.Equal(t, nil, err)

	assert.Equal(t, productFound.Name, newProduct.Name)
	assert.Equal(t, productFound.Price, newProduct.Price)
}

func TestFindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Equal(t, nil, err)
	db.AutoMigrate(&product.Product{})

	for i := 0; i < 24; i++ {
		newProduct, err := product.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(newProduct)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 0", products[0].Name)
	assert.Equal(t, "Product 9", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 4)
	assert.Equal(t, "Product 20", products[0].Name)
	assert.Equal(t, "Product 23", products[3].Name)
}
