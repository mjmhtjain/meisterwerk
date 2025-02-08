package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/meisterwerk/internal/services"
)

type ProductHandler struct {
	productService services.ProductServiceI
}

func NewProductHandler(service services.ProductServiceI) *ProductHandler {
	return &ProductHandler{
		productService: service,
	}
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, products)
}
