package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/meisterwerk/internal/services"
)

type OrderHandler struct {
	orderService services.OrderServiceI
}

func NewOrderHandler(service services.OrderServiceI) *OrderHandler {
	return &OrderHandler{
		orderService: service,
	}
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.orderService.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, orders)
}
