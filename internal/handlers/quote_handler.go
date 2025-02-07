package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mjmhtjain/meisterwerk/internal/models"
)

type QuoteHandler struct {
	quotes []models.Quote // For now, we'll use in-memory storage
}

func NewQuoteHandler() *QuoteHandler {
	return &QuoteHandler{
		quotes: make([]models.Quote, 0),
	}
}

func (h *QuoteHandler) CreateQuote() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.Quote
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		quote := models.Quote{
			ID:     uuid.New().String(),
			Author: req.Author,
			Status: "created",
		}

		h.quotes = append(h.quotes, quote)
		c.JSON(http.StatusCreated, quote)
	}
}
