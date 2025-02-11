package dto

type OrderResponse struct {
	ID        string `json:"id"`
	Quote_ID  string `json:"quote_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
