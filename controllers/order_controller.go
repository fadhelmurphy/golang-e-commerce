package controllers

import (
	"encoding/json"
	"golang-ecommerce/domain"
	"golang-ecommerce/usecases"
	"net/http"
	"strconv"
)

type OrderController struct {
	orderUsecase usecases.OrderUsecase
}

func NewOrderController(usecase usecases.OrderUsecase) *OrderController {
	return &OrderController{orderUsecase: usecase}
}

func (c *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order domain.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := c.orderUsecase.CreateOrder(&order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func (c *OrderController) GetOrdersByUserID(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	orders, err := c.orderUsecase.GetOrdersByUserID(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
