package dto

type CreateQuoteRequest struct {
	Author string `json:"author" binding:"required"`
}
