package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ViniciusSouzaDosReis/product-api/internal/dto"
	"github.com/ViniciusSouzaDosReis/product-api/internal/entity"
	"github.com/ViniciusSouzaDosReis/product-api/internal/infra/database/interfaces"
	"github.com/go-chi/chi/v5"
)

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
	p, err := entity.NewProduct(dtoProduct.Name, dtoProduct.Price)
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

func (h *ProductHandler) FindProductById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
