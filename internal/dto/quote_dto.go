package dto

type CreateQuoteRequest struct {
	Author       string   `json:"author" binding:"required"`
	CustomerName string   `json:"customer_name" binding:"required"`
	ProductList  []string `json:"product_list" binding:"required"`
}

type QuoteResponse struct {
	ID           string            `json:"id"`
	Author       string            `json:"author"`
	CustomerName string            `json:"customer_name"`
	ProductList  []ProductResponse `json:"product_list"`
	TotalPrice   float64           `json:"total_price"`
	TotalTax     float64           `json:"total_tax"`
	Status       string            `json:"status"`
}

type UpdateQuoteStatusRequest struct {
	Status string `json:"status" binding:"required"`
}
