package mocks

import "github.com/mjmhtjain/meisterwerk/internal/dto"

// MockQuoteService is a custom mock implementation of QuoteServiceI
type MockQuoteService struct {
	CreateQuoteFunc       func(req dto.CreateQuoteRequest) (dto.QuoteResponse, error)
	GetQuoteFunc          func(id string) (dto.QuoteResponse, error)
	UpdateQuoteStatusFunc func(id string, status dto.QuoteStatus) error
}

func (m *MockQuoteService) CreateQuote(req dto.CreateQuoteRequest) (dto.QuoteResponse, error) {

	return m.CreateQuoteFunc(req)
}

func (m *MockQuoteService) GetQuote(id string) (dto.QuoteResponse, error) {
	return m.GetQuoteFunc(id)
}

func (m *MockQuoteService) UpdateQuoteStatus(id string, status dto.QuoteStatus) error {
	return m.UpdateQuoteStatusFunc(id, status)
}
