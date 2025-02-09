package mocks

import "github.com/mjmhtjain/meisterwerk/internal/models"

// Custom mock for ProductRepositoryI
type MockProductRepository struct {
	Products []models.Product
	Err      error
}

func (m *MockProductRepository) GetAll() ([]models.Product, error) {
	return m.Products, m.Err
}

func (m *MockProductRepository) GetByID(id string) (models.Product, error) {
	if m.Err != nil {
		return models.Product{}, m.Err
	}

	return m.Products[0], nil
}
