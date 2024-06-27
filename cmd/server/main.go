package main

import (
	"encoding/json"
	"net/http"

	"github.com/ViniciusSouzaDosReis/product-api/configs"
	"github.com/ViniciusSouzaDosReis/product-api/internal/dto"
	"github.com/ViniciusSouzaDosReis/product-api/internal/entity/product"
	"github.com/ViniciusSouzaDosReis/product-api/internal/entity/user"
	"github.com/ViniciusSouzaDosReis/product-api/internal/infra/database/interfaces"
	"github.com/ViniciusSouzaDosReis/product-api/internal/infra/database/product_database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	checkErr(err)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	checkErr(err)
	db.AutoMigrate(&user.User{}, &product.Product{})
	productHandler := NewProductHandler(product_database.NewProduct(db))

	http.HandleFunc("POST /product", productHandler.CreateProduct)
	http.ListenAndServe(":8080", nil)
}

type ProductHandler struct {
	ProductDB interfaces.ProductInterface
}

func NewProductHandler(db interfaces.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var dtoProduct dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&dtoProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	p, err := product.NewProduct(dtoProduct.Name, dtoProduct.Price)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
