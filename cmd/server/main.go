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
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configs, err := configs.LoadConfig(".")
	checkErr(err)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	checkErr(err)
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productHandler := handlers.NewProductHandler(product_database.NewProductDB(db))
	userHandler := handlers.NewUserHandler(user_database.NewUserDB(db), configs.TokenAuth, configs.JWTExperesIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/product", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProductById)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/user", userHandler.CreateUser)
	r.Post("/user/generate_token", userHandler.GenerateToken)

	http.ListenAndServe(":8080", r)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
