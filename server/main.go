package main

import (
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

	db.AutoMigrate(&payment.Payment{})
	db.AutoMigrate(&product.Product{})

	/*
		api.POST("/task", taskHandler.Store)
		api.GET("/task", taskHandler.FetchAll)
		api.GET("/task/:id", taskHandler.FetchById)
		api.PUT("/task/:id", taskHandler.Update)
		api.DELETE("/task/:id", taskHandler.Delete)
	*/
	r.Run(":3000")
}
