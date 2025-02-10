package mocks

import (
	"github.com/mjmhtjain/meisterwerk/internal/dto"
	"github.com/mjmhtjain/meisterwerk/internal/models"
)

// Custom mock for QuoteRepository
type MockQuoteRepository struct {
	Quotes   []models.Quote
	Err      error
	Products []models.Product
}

func (m *MockQuoteRepository) Create(quote models.Quote) error {
	m.Quotes = append(m.Quotes, quote)
	return m.Err
}

func (m *MockQuoteRepository) CreateQuoteProductMap(quoteID string, productID string) error {
	// Simulate successful mapping
	return m.Err
}

func (m *MockQuoteRepository) GetByID(id string) (models.Quote, error) {
	if m.Err != nil {
		return models.Quote{}, m.Err
	}

	return m.Quotes[0], nil
}

func (m *MockQuoteRepository) GetAll() ([]models.Quote, error) {
	return m.Quotes, m.Err
}

func (m *MockQuoteRepository) UpdateQuoteStatus(id string, status dto.QuoteStatus) error {
	return m.Err
}

func (m *MockQuoteRepository) GetProductsByQuoteID(id string) ([]models.Product, error) {
	return m.Products, m.Err
}
