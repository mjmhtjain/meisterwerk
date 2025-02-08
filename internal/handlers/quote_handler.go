package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mjmhtjain/meisterwerk/internal/dto"
	"github.com/mjmhtjain/meisterwerk/internal/models"
	"gorm.io/gorm"
)

type QuoteHandler struct {
	db *gorm.DB
}

func NewQuoteHandler(db *gorm.DB) *QuoteHandler {
	return &QuoteHandler{
		db: db,
	}
}

func (h *QuoteHandler) CreateQuote() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CreateQuoteRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		quote := models.Quote{
			ID:     uuid.New().String(),
			Author: req.Author,
			Status: "created",
		}

		if err := h.db.Create(&quote).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, quote)
	}
}
