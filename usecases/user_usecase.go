package usecases

import (
	"golang-ecommerce/domain"
	"golang-ecommerce/repositories"
)

type UserUsecase interface {
	GetUserByID(id uint) (*domain.User, error)
	CreateUser(user *domain.User) error
}

type userUsecase struct {
	userRepo repositories.UserRepository
}

func NewUserUsecase(repo repositories.UserRepository) UserUsecase {
	return &userUsecase{userRepo: repo}
}

func (u *userUsecase) GetUserByID(id uint) (*domain.User, error) {
	return u.userRepo.FindByID(id)
}

func (u *userUsecase) CreateUser(user *domain.User) error {
	return u.userRepo.Create(user)
}
