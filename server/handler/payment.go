package handler

import (
	"go-api/server/payment"
	"io"
	"net/http"
	"strconv"

	broadcast "go-api/server/broadcast"

	"github.com/gin-gonic/gin"
)

type paymentHandler struct {
	paymentService payment.Service
	broadcast      broadcast.Broadcaster
}

func NewPaymentHandler(paymentService payment.Service, broadcast broadcast.Broadcaster) *paymentHandler {
	return &paymentHandler{paymentService, broadcast}
}

type Message struct {
	Text      string
	Name      string
	PricePaid float64
}

func (th *paymentHandler) Create(c *gin.Context) {
	// Get json body
	var input payment.InputPayment
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Cannot extract JSON body",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newPayment, err := th.paymentService.Create(input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	th.broadcast.Submit(Message{
		Text:      "New payment created",
		Name:      newPayment.Product.Name,
		PricePaid: newPayment.Product.Price,
	})

	response := &Response{
		Success: true,
		Message: "New Payment created",
		Data:    newPayment,
	}
	c.JSON(http.StatusCreated, response)
}

func (th *paymentHandler) GetAll(c *gin.Context) {
	payments, err := th.paymentService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Data:    payments,
	})
}

func (th *paymentHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	payment, err := th.paymentService.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Data:    payment,
	})
}

func (th *paymentHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	// Get json body
	var input payment.InputPayment
	err = c.ShouldBindJSON(&input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Cannot extract JSON body",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	uPayment, err := th.paymentService.Update(id, input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	th.broadcast.Submit(Message{
		Text:      "Payment updated",
		Name:      uPayment.Product.Name,
		PricePaid: uPayment.Product.Price,
	})
	response := &Response{
		Success: true,
		Message: "Payment updated",
		Data:    uPayment,
	}
	c.JSON(http.StatusCreated, response)
}

func (th *paymentHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	err = th.paymentService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "Payment successfully deleted",
	})
}

func (ph *paymentHandler) Stream(c *gin.Context) {
	listener := make(chan interface{})
	ph.broadcast.Register(listener)
	defer ph.broadcast.Unregister(listener)

	clientGone := c.Request.Context().Done()
	c.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			return false
		case message := <-listener:
			serviceMsg, ok := message.(Message)
			if !ok {
				c.SSEvent("message", message)
				return false
			}
			c.SSEvent("message", serviceMsg.Text+", "+"Produit : "+serviceMsg.Name+", "+"Prix : "+strconv.FormatFloat(serviceMsg.PricePaid, 'f', 2, 64))
			return true
		}
	})
}
