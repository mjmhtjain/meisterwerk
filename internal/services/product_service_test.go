package services

import (
	"testing"

	"github.com/mjmhtjain/meisterwerk/internal/dto"
	"github.com/mjmhtjain/meisterwerk/internal/models"
	"github.com/mjmhtjain/meisterwerk/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAllProducts(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := &mocks.MockProductRepository{
			Products: []models.Product{
				{ID: "1", Name: "Product 1", Price: 100, Tax: 10},
				{ID: "2", Name: "Product 2", Price: 200, Tax: 20},
			},
			Err: nil,
		}
		service := &ProductService{productRepo: mockRepo}

		result, err := service.GetAllProducts()

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "Product 1", result[0].Name)
		assert.Equal(t, "Product 2", result[1].Name)
	})

	t.Run("Empty product list", func(t *testing.T) {
		mockRepo := &mocks.MockProductRepository{
			Products: []models.Product{},
			Err:      nil,
		}
		service := &ProductService{productRepo: mockRepo}

		result, err := service.GetAllProducts()

		assert.NoError(t, err)
		assert.Len(t, result, 0)
	})

	t.Run("Error scenario", func(t *testing.T) {
		mockRepo := &mocks.MockProductRepository{
			Products: nil,
			Err:      assert.AnError,
		}
		service := &ProductService{productRepo: mockRepo}

		result, err := service.GetAllProducts()

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestGetProduct(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := &mocks.MockProductRepository{
			Products: []models.Product{
				{ID: "1", Name: "Product 1", Price: 100, Tax: 10},
			},
			Err: nil,
		}
		service := &ProductService{productRepo: mockRepo}

		result, err := service.GetProduct("1")

		assert.NoError(t, err)
		assert.Equal(t, "1", result.ID)
		assert.Equal(t, "Product 1", result.Name)
		assert.Equal(t, 100.0, result.Price)
		assert.Equal(t, 10.0, result.Tax)
	})

	t.Run("Product not found", func(t *testing.T) {
		mockRepo := &mocks.MockProductRepository{
			Products: []models.Product{},
			Err:      assert.AnError,
		}
		service := &ProductService{productRepo: mockRepo}

		result, err := service.GetProduct("2")

		assert.Error(t, err)
		assert.Equal(t, dto.ProductResponse{}, result)
	})

	t.Run("Error scenario", func(t *testing.T) {
		mockRepo := &mocks.MockProductRepository{
			Products: nil,
			Err:      assert.AnError,
		}
		service := &ProductService{productRepo: mockRepo}

		result, err := service.GetProduct("1")

		assert.Error(t, err)
		assert.Equal(t, dto.ProductResponse{}, result)
	})
}
