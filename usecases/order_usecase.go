package usecases

import (
	"golang-ecommerce/domain"
	"golang-ecommerce/repositories"
)

type OrderUsecase interface {
	CreateOrder(order *domain.Order) error
	GetOrdersByUserID(userID uint) ([]domain.Order, error)
}

type orderUsecase struct {
	orderRepo repositories.OrderRepository
}

func NewOrderUsecase(repo repositories.OrderRepository) OrderUsecase {
	return &orderUsecase{orderRepo: repo}
}

func (u *orderUsecase) CreateOrder(order *domain.Order) error {
	return u.orderRepo.Create(order)
}

func (u *orderUsecase) GetOrdersByUserID(userID uint) ([]domain.Order, error) {
	return u.orderRepo.FindByUserID(userID)
}
