package models

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	TaxRate     float64 `json:"taxRate"`
}
