package models

type Quote struct {
	ID           string `json:"id"`
	Author       string `json:"author"`
	CustomerName string `json:"customer_name"`
	Status       string `json:"status"`
}
