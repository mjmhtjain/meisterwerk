package services

import (
	"errors"
	"testing"

	"github.com/mjmhtjain/meisterwerk/internal/dto"
	"github.com/mjmhtjain/meisterwerk/internal/models"
	"github.com/mjmhtjain/meisterwerk/internal/repository"
	repoMocks "github.com/mjmhtjain/meisterwerk/internal/repository/mocks"
	serviceMocks "github.com/mjmhtjain/meisterwerk/internal/services/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateQuote(t *testing.T) {
	tests := []struct {
		name               string
		quoteReq           dto.CreateQuoteRequest
		mockRepo           repository.QuoteRepositoryI
		mockProductService ProductServiceI
		expectedError      bool
	}{
		{
			name: "Success",
			quoteReq: dto.CreateQuoteRequest{
				Author:       "John Doe",
				CustomerName: "Martha Smith",
				ProductList:  []string{"product1", "product2"},
			},
			mockRepo: &repoMocks.MockQuoteRepository{
				Quotes: []models.Quote{},
			},
			mockProductService: &serviceMocks.MockProductService{
				Products: []dto.ProductResponse{
					{ID: "product1"},
					{ID: "product2"},
				},
				Product: dto.ProductResponse{ID: "product1"},
				Err:     nil,
			},
			expectedError: false,
		},
		{
			name: "Repository Error when creating quote",
			quoteReq: dto.CreateQuoteRequest{
				Author:       "John Doe",
				CustomerName: "Martha Smith",
				ProductList:  []string{"product1", "product2"},
			},
			mockRepo: &repoMocks.MockQuoteRepository{
				Quotes: []models.Quote{},
				Err:    errors.New("repository error"),
			},
			mockProductService: &serviceMocks.MockProductService{
				Products: []dto.ProductResponse{},
				Err:      nil,
			},
			expectedError: true,
		},
		{
			name: "Product Service Error when getting product",
			quoteReq: dto.CreateQuoteRequest{
				Author:       "John Doe",
				CustomerName: "Martha Smith",
				ProductList:  []string{"product1", "product2"},
			},
			mockRepo: &repoMocks.MockQuoteRepository{
				Quotes: []models.Quote{},
				Err:    nil,
			},
			mockProductService: &serviceMocks.MockProductService{
				Products: []dto.ProductResponse{},
				Err:      errors.New("product service error"),
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quoteService := &QuoteService{
				productService: tt.mockProductService,
				quoteRepo:      tt.mockRepo,
			}

			quoteResponse, err := quoteService.CreateQuote(tt.quoteReq)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Empty(t, quoteResponse.ID)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, quoteResponse.ID)
				assert.Equal(t, tt.quoteReq.Author, quoteResponse.Author)
				assert.Equal(t, tt.quoteReq.CustomerName, quoteResponse.CustomerName)
				assert.Equal(t, "created", quoteResponse.Status)
			}
		})
	}
}
