package main

import (
	"net/http"

	"github.com/ViniciusSouzaDosReis/product-api/configs"
	_ "github.com/ViniciusSouzaDosReis/product-api/docs"
	"github.com/ViniciusSouzaDosReis/product-api/internal/entity"
	"github.com/ViniciusSouzaDosReis/product-api/internal/infra/database/product_database"
	"github.com/ViniciusSouzaDosReis/product-api/internal/infra/database/user_database"
	"github.com/ViniciusSouzaDosReis/product-api/internal/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title Go Expert API Example
// @version 1.0
// @description Product API with authentication
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

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

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))

	http.ListenAndServe(":8080", r)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
