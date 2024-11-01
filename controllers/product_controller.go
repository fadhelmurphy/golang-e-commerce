package controllers

import (
	"encoding/json"
	"golang-ecommerce/domain"
	"golang-ecommerce/usecases"
	"net/http"
	"strconv"
)

type ProductController struct {
	productUsecase usecases.ProductUsecase
}

func NewProductController(usecase usecases.ProductUsecase) *ProductController {
	return &ProductController{productUsecase: usecase}
}

func (c *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	products, err := c.productUsecase.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (c *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := c.productUsecase.GetProductByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product domain.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := c.productUsecase.CreateProduct(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
