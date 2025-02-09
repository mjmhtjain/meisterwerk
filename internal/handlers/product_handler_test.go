package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/meisterwerk/internal/dto"
	"github.com/mjmhtjain/meisterwerk/internal/services"
	"github.com/mjmhtjain/meisterwerk/internal/services/mocks"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter(mockService services.ProductServiceI) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	handler := NewProductHandler(mockService)

	router.GET("/products", handler.GetAllProducts)

	return router
}

func TestGetAllProducts(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockService := &mocks.MockProductService{
			Products: []dto.ProductResponse{
				{ID: "1", Name: "Product 1", Price: 100, Tax: 10},
				{ID: "2", Name: "Product 2", Price: 200, Tax: 20},
			},
		}
		router := setupTestRouter(mockService)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/products", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []dto.ProductResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, len(mockService.Products), len(response))
		assert.Equal(t, mockService.Products[0].ID, response[0].ID)
		assert.Equal(t, mockService.Products[1].ID, response[1].ID)
	})

	t.Run("NoProducts", func(t *testing.T) {
		mockService := &mocks.MockProductService{
			Products: []dto.ProductResponse{},
		}
		router := setupTestRouter(mockService)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/products", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []dto.ProductResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(response))
	})

	t.Run("ServiceError", func(t *testing.T) {
		mockService := &mocks.MockProductService{
			Err: errors.New("service error"),
		}
		router := setupTestRouter(mockService)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/products", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("ResponseStructure", func(t *testing.T) {
		mockService := &mocks.MockProductService{
			Products: []dto.ProductResponse{
				{ID: "1", Name: "Product 1", Price: 100, Tax: 10},
			},
		}
		router := setupTestRouter(mockService)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/products", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []dto.ProductResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(response))
		assert.Equal(t, "1", response[0].ID)
		assert.Equal(t, "Product 1", response[0].Name)
	})
}
