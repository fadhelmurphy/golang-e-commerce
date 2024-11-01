package usecases

import (
	"golang-ecommerce/domain"
	"golang-ecommerce/repositories"
)

type ProductUsecase interface {
	GetAllProducts() ([]domain.Product, error)
	GetProductByID(id uint) (*domain.Product, error)
	CreateProduct(product *domain.Product) error
}

type productUsecase struct {
	productRepo repositories.ProductRepository
}

func NewProductUsecase(repo repositories.ProductRepository) ProductUsecase {
	return &productUsecase{productRepo: repo}
}

func (u *productUsecase) GetAllProducts() ([]domain.Product, error) {
	return u.productRepo.FindAll()
}

func (u *productUsecase) GetProductByID(id uint) (*domain.Product, error) {
	return u.productRepo.FindByID(id)
}

func (u *productUsecase) CreateProduct(product *domain.Product) error {
	return u.productRepo.Create(product)
}
