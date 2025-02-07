package service

import (
	"errors"
	"sync"

	"github.com/mjmhtjain/meisterwerk/internal/models"
)

type ProductService struct {
	products map[string]models.Product
	mutex    sync.RWMutex
}

func NewProductService() *ProductService {
	return &ProductService{
		products: make(map[string]models.Product),
	}
}

func (s *ProductService) GetAllProducts() []models.Product {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	products := make([]models.Product, 0, len(s.products))
	for _, product := range s.products {
		products = append(products, product)
	}
	return products
}

func (s *ProductService) GetProduct(id string) (models.Product, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	product, exists := s.products[id]
	if !exists {
		return models.Product{}, errors.New("product not found")
	}
	return product, nil
}

func (s *ProductService) CreateProduct(product models.Product) (models.Product, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// In a real application, you would generate a proper UUID
	product.ID = "generated-id"
	s.products[product.ID] = product
	return product, nil
}

func (s *ProductService) UpdateProduct(id string, product models.Product) (models.Product, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.products[id]; !exists {
		return models.Product{}, errors.New("product not found")
	}

	product.ID = id
	s.products[id] = product
	return product, nil
}

func (s *ProductService) DeleteProduct(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.products[id]; !exists {
		return errors.New("product not found")
	}

	delete(s.products, id)
	return nil
}
