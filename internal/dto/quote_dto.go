package dto

type CreateQuoteRequest struct {
	Author       string `json:"author" binding:"required"`
	CustomerName string `json:"customer_name" binding:"required"`
}

type QuoteResponse struct {
	ID           string `json:"id"`
	Author       string `json:"author"`
	CustomerName string `json:"customer_name"`
	Status       string `json:"status"`
}
