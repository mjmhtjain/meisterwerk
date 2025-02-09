package services

import (
	"github.com/mjmhtjain/meisterwerk/internal/dto"
	"github.com/mjmhtjain/meisterwerk/internal/repository"
)

type ProductServiceI interface {
	GetAllProducts() ([]dto.ProductResponse, error)
	GetProduct(id string) (dto.ProductResponse, error)
}

type ProductService struct {
	productRepo repository.ProductRepositoryI
}

func NewProductService() ProductServiceI {

	return &ProductService{
		productRepo: repository.NewProductRepository(),
	}
}

func (s *ProductService) GetAllProducts() ([]dto.ProductResponse, error) {
	products, err := s.productRepo.GetAll()
	if err != nil {
		return nil, err
	}

	productResponses := make([]dto.ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = dto.ProductResponse{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
			Tax:   product.Tax,
		}
	}

	return productResponses, nil
}

func (s *ProductService) GetProduct(id string) (dto.ProductResponse, error) {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	return dto.ProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
		Tax:   product.Tax,
	}, nil
}
