package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ViniciusSouzaDosReis/product-api/internal/dto"
	"github.com/ViniciusSouzaDosReis/product-api/internal/entity"
	"github.com/ViniciusSouzaDosReis/product-api/internal/infra/database/interfaces"
	entityPkg "github.com/ViniciusSouzaDosReis/product-api/pkg/entity"
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

// CreateProduct godoc
// @Summary      Create product
// @Description  Create product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        request	body dto.CreateProductInput true "product request"
// @Success      201
// @Failure      500	{object} Error
// @Router       /product [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var dtoProduct dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&dtoProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}
	p, err := entity.NewProduct(dtoProduct.Name, dtoProduct.Price)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetProductById godoc
// @Summary      Get Product by id
// @Description  Get Product by id
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        id	path string true "Product ID" Format(uuid)
// @Success      200	{object} entity.Product
// @Failure      400	{object} Error
// @Failure      404	{object} Error
// @Router       /product/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: "Id is nod valid",
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	product, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// GetProducts godoc
// @Summary      Get Products
// @Description  Get Products
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        page	query string false "Page number"
// @Param        limit	query string false "Limit number"
// @Param        sort	query string false "Sort page"
// @Success      200	{object} entity.Product
// @Failure      400	{object} Error
// @Failure      404	{object} Error
// @Failure      500	{object} Error
// @Router       /product [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}
	sort := r.URL.Query().Get("sort")

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// UpdateProduct godoc
// @Summary      Update product
// @Description  Update product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        request	body entity.Product true "product request"
// @Param        id	path string true "Product ID" Format(uuid)
// @Success      200	{array} entity.Product
// @Failure      400	{object} Error
// @Router       /product/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: "Id is invalid",
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	_, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	var product entity.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// DeleteProduct godoc
// @Summary      Delete product
// @Description  Delete product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        id	path string true "Product ID" Format(uuid)
// @Success      200	{array} entity.Product
// @Failure      400	{object} Error
// @Failure      404	{object} Error
// @Failure      500	{object} Error
// @Router       /product/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: "Id is not valid",
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	_, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusOK)
}
