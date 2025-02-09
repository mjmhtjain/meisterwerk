package mocks

import (
	"github.com/mjmhtjain/meisterwerk/internal/dto"
)

// MockProductService is a simple mock implementation
type MockProductService struct {
	Products []dto.ProductResponse
	Product  dto.ProductResponse
	Err      error
}

func (m *MockProductService) GetAllProducts() ([]dto.ProductResponse, error) {
	return m.Products, m.Err
}

func (m *MockProductService) GetProduct(id string) (dto.ProductResponse, error) {
	return m.Product, m.Err
}
