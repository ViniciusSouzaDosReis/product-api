package product_database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ViniciusSouzaDosReis/product-api/internal/entity"
	"github.com/ViniciusSouzaDosReis/product-api/internal/infra/database/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
)

var sqliteDialecto = sqlite.Open("file::memory:")

func TestCreateProduct(t *testing.T) {
	db, err := utils.CreateDBConnection(sqliteDialecto, &entity.Product{})
	assert.NoError(t, err)
	newProduct, _ := entity.NewProduct("Product 1", 10.0)
	productDB := NewProductDB(db)

	err = productDB.Create(newProduct)
	assert.Equal(t, nil, err)

	var productFound entity.Product
	err = db.First(&productFound, "id = ?", newProduct.ID).Error
	assert.Equal(t, nil, err)

	assert.Equal(t, productFound.Name, newProduct.Name)
	assert.Equal(t, productFound.Price, newProduct.Price)
}

func TestFindAll(t *testing.T) {
	db, err := utils.CreateDBConnection(sqliteDialecto, &entity.Product{})
	assert.NoError(t, err)

	for i := 0; i < 24; i++ {
		newProduct, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(newProduct)
	}

	productDB := NewProductDB(db)
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

func TestFindById(t *testing.T) {
	db, err := utils.CreateDBConnection(sqliteDialecto, &entity.Product{})
	assert.NoError(t, err)

	newProduct, _ := entity.NewProduct("Product 1", 10.0)
	db.Create(newProduct)

	productDB := NewProductDB(db)
	productFound, err := productDB.FindById(newProduct.ID.String())
	assert.NoError(t, err)

	assert.Equal(t, newProduct.Name, productFound.Name)
	assert.Equal(t, newProduct.Price, productFound.Price)
}

func TestDelete(t *testing.T) {
	db, err := utils.CreateDBConnection(sqliteDialecto, &entity.Product{})
	assert.NoError(t, err)

	newProduct, _ := entity.NewProduct("Product 1", 10.0)
	db.Create(newProduct)

	productDB := NewProductDB(db)
	err = productDB.Delete(newProduct.ID.String())
	assert.NoError(t, err)

	var foundProduct entity.Product
	err = db.First(&foundProduct, "id = ?", newProduct.ID.String()).Error
	assert.Error(t, err)
}

func TestUpdate(t *testing.T) {
	db, err := utils.CreateDBConnection(sqliteDialecto, &entity.Product{})
	assert.NoError(t, err)

	newProduct, _ := entity.NewProduct("Product 1", 10.0)
	db.Create(newProduct)

	productDB := NewProductDB(db)
	productToUpdate := entity.Product{
		ID:        newProduct.ID,
		Name:      "Product 2",
		Price:     110.0,
		CreatedAt: newProduct.CreatedAt,
	}
	err = productDB.Update(&productToUpdate)
	assert.NoError(t, err)

	var productFound entity.Product
	db.First(&productFound, "id = ?", productToUpdate.ID.String())
	assert.Equal(t, productToUpdate.Name, productFound.Name)
	assert.Equal(t, productToUpdate.Price, productFound.Price)
}
