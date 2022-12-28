package main

import (
    "go-api/server/handler"

	"log"
	"os"

// 	"go-api/server/payment"
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

    productRepository := product.NewRepository(db)
    productService := product.NewService(productRepository)
    productHandler := handler.NewProductHandler(productService)

	api := r.Group("/api")

    api.POST("/product", productHandler.Store)
    api.GET("/product", productHandler.List)

    r.Run(":3000")
}
