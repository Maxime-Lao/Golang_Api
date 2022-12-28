package handler

import (
	"fmt"
	"go-api/server/product"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{productService}
}

func (handler *productHandler) Store(c *gin.Context) {
	var input product.InputProduct
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newProduct, err := handler.productService.Store(input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := Response{
		Success: true,
		Message: "new product successfully added",
		Data:    newProduct,
	}

	c.JSON(http.StatusOK, response)
}

func (handler *productHandler) List(c *gin.Context) {

	products, err := handler.productService.ListAll()
	if err != nil {
		response := Response{
			Success: false,
			Message: "something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := Response{
		Success: true,
		Data:    products,
	}

	fmt.Println(products)

	c.JSON(http.StatusOK, response)
}
