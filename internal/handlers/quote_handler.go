package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/meisterwerk/internal/dto"
	"github.com/mjmhtjain/meisterwerk/internal/services"
)

type QuoteHandler struct {
	quoteService services.QuoteServiceI
}

func NewQuoteHandler(service services.QuoteServiceI) *QuoteHandler {
	return &QuoteHandler{
		quoteService: service,
	}
}

func (h *QuoteHandler) CreateQuote(c *gin.Context) {
	var req dto.CreateQuoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quoteResponse, err := h.quoteService.CreateQuote(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, quoteResponse)
}

func (h *QuoteHandler) GetQuote(c *gin.Context) {
	id := c.Param("id")

	quoteResponse, err := h.quoteService.GetQuote(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, quoteResponse)
}

func (h *QuoteHandler) UpdateQuoteStatus(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateQuoteStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.quoteService.UpdateQuoteStatus(id, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Quote status updated successfully"})
}
