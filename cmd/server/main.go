package main

import (
	"net/http"

	"github.com/ViniciusSouzaDosReis/product-api/configs"
	"github.com/ViniciusSouzaDosReis/product-api/internal/entity"
	"github.com/ViniciusSouzaDosReis/product-api/internal/infra/database/product_database"
	"github.com/ViniciusSouzaDosReis/product-api/internal/infra/database/user_database"
	"github.com/ViniciusSouzaDosReis/product-api/internal/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	userHandler := handlers.NewUserHandler(user_database.NewUserDB(db))

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/product", productHandler.CreateProduct)
	r.Get("/product", productHandler.GetProducts)
	r.Get("/product/{id}", productHandler.GetProductById)
	r.Put("/product/{id}", productHandler.UpdateProduct)
	r.Delete("/product/{id}", productHandler.DeleteProduct)

	r.Post("/user", userHandler.CreateUser)

	http.ListenAndServe(":8080", r)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
