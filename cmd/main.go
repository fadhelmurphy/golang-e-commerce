package main

import (
	"fmt"
	configs "golang-ecommerce/configs"
	"golang-ecommerce/controllers"
	"golang-ecommerce/helpers"
	"golang-ecommerce/repositories"
	"golang-ecommerce/usecases"
	"log"
	"net/http"
)

func main() {
	configs.InitDB()
	db := configs.GetDB()

	// Product setup
	productRepo := repositories.NewProductRepository(db)
	productUsecase := usecases.NewProductUsecase(productRepo)
	productController := controllers.NewProductController(productUsecase)

	// User setup
	userRepo := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userController := controllers.NewUserController(userUsecase)

	// Order setup
	orderRepo := repositories.NewOrderRepository(db)
	orderUsecase := usecases.NewOrderUsecase(orderRepo)
	orderController := controllers.NewOrderController(orderUsecase)

	http.HandleFunc("/products", helpers.ResponseWrapper(productController.GetAllProducts))
	http.HandleFunc("/product", productController.GetProductByID)
	http.HandleFunc("/product/create", productController.CreateProduct)

	http.HandleFunc("/user", userController.GetUserByID)
	http.HandleFunc("/user/create", userController.CreateUser)

	http.HandleFunc("/order/create", orderController.CreateOrder)
	http.HandleFunc("/orders", orderController.GetOrdersByUserID)

	fmt.Println("Server is running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}