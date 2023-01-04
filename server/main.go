package main

import (
	broadcast "go-api/server/broadcast"
	"go-api/server/handler"

	"log"
	"os"

	"go-api/server/payment"
	"go-api/server/product"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "user:password@tcp(127.0.0.1:3306)/go-api?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&payment.Payment{})

	b := broadcast.NewBroadcaster(20)

	// Product
	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	// Payment
	paymentRepository := payment.NewRepository(db)
	paymentService := payment.NewService(paymentRepository)
	paymentHandler := handler.NewPaymentHandler(paymentService, b)

	api := r.Group("/api")

	// Product routes
	api.POST("/product", productHandler.Create)
	api.GET("/products", productHandler.GetAll)
	api.GET("/product/:id", productHandler.GetById)
	api.PUT("/product/:id", productHandler.Update)
	api.DELETE("/product/:id", productHandler.Delete)

	// Payment routes
	api.POST("/payment", paymentHandler.Create)
	api.GET("/payments", paymentHandler.GetAll)
	api.GET("/payment/:id", paymentHandler.GetById)
	api.PUT("/payment/:id", paymentHandler.Update)
	api.DELETE("/payment/:id", paymentHandler.Delete)
	api.GET("/stream", paymentHandler.Stream)

	r.Run(":3000")
}
