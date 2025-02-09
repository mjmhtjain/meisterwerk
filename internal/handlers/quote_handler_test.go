package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/meisterwerk/internal/dto"
	"github.com/mjmhtjain/meisterwerk/internal/handlers/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateQuote(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		request        dto.CreateQuoteRequest
		setupMock      func(*mocks.MockQuoteService)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "Success",
			request: dto.CreateQuoteRequest{
				Author:       "author123",
				CustomerName: "customer123",
				ProductList:  []string{"product1", "product2"},
			},
			setupMock: func(m *mocks.MockQuoteService) {
				m.CreateQuoteFunc = func(req dto.CreateQuoteRequest) (dto.QuoteResponse, error) {
					return dto.QuoteResponse{
						ID:           "quote123",
						Author:       "author123",
						CustomerName: req.CustomerName,
						Status:       "pending",
					}, nil
				}
			},
			expectedStatus: http.StatusCreated,
			expectedBody: dto.QuoteResponse{
				ID:           "quote123",
				Author:       "author123",
				CustomerName: "customer123",
				Status:       "pending",
			},
		},
		{
			name:           "Invalid Request Body",
			request:        dto.CreateQuoteRequest{},
			setupMock:      func(m *mocks.MockQuoteService) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: gin.H{
				"error": "Key: 'CreateQuoteRequest.Author' Error:Field validation for 'Author' failed on the 'required' tag\nKey: 'CreateQuoteRequest.CustomerName' Error:Field validation for 'CustomerName' failed on the 'required' tag\nKey: 'CreateQuoteRequest.ProductList' Error:Field validation for 'ProductList' failed on the 'required' tag",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockService := &mocks.MockQuoteService{}
			if tt.setupMock != nil {
				tt.setupMock(mockService)
			}

			handler := NewQuoteHandler(mockService)
			router := gin.New()
			router.POST("/quotes", handler.CreateQuote)

			// Create request
			requestBody, _ := json.Marshal(tt.request)
			req, _ := http.NewRequest(http.MethodPost, "/quotes", bytes.NewBuffer(requestBody))
			req.Header.Set("Content-Type", "application/json")

			// Create response recorder
			w := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusCreated {
				var response dto.QuoteResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				expectedJSON, _ := json.Marshal(tt.expectedBody)
				actualJSON, _ := json.Marshal(response)
				assert.JSONEq(t, string(expectedJSON), string(actualJSON))
			} else {
				var response gin.H
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, response)
			}
		})
	}
}
