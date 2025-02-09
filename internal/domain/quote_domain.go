package domain

type Quote struct {
	ID           string
	Author       string
	CustomerName string
	ProductList  []Product
	Status       string
}

type Product struct {
	ID    string
	Name  string
	Price float64
	Tax   float64
}
