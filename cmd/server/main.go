package main

import (
	"net/http"

	"github.com/ViniciusSouzaDosReis/product-api/configs"
	"github.com/ViniciusSouzaDosReis/product-api/internal/entity"
	"github.com/ViniciusSouzaDosReis/product-api/internal/infra/database/product_database"
	"github.com/ViniciusSouzaDosReis/product-api/internal/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	checkErr(err)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	checkErr(err)
	db.AutoMigrate(&entity.User{}, &entity.Product{})
	productHandler := handlers.NewProductHandler(product_database.NewProductDB(db))

	http.HandleFunc("POST /product", productHandler.CreateProduct)
	http.ListenAndServe(":8080", nil)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
