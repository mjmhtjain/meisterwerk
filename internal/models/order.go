package models

import "time"

type Order struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	QuoteFK   string    `json:"quote_fk"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
