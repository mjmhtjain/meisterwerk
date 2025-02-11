package services

import (
	"time"

	"github.com/mjmhtjain/meisterwerk/internal/dto"
	"github.com/mjmhtjain/meisterwerk/internal/models"
	"github.com/mjmhtjain/meisterwerk/internal/repository"
)

type OrderServiceI interface {
	CreateOrder(order models.Order) (models.Order, error)
	GetAllOrders() ([]dto.OrderResponse, error)
}

type OrderService struct {
	orderRepo repository.OrderRepositoryI
}

func NewOrderService() *OrderService {
	return &OrderService{
		orderRepo: repository.NewOrderRepository(),
	}
}

func (s *OrderService) CreateOrder(order models.Order) (models.Order, error) {
	return s.orderRepo.CreateOrder(order)
}

func (s *OrderService) GetAllOrders() ([]dto.OrderResponse, error) {
	orders, err := s.orderRepo.GetAll()
	if err != nil {
		return nil, err
	}

	orderResponses := []dto.OrderResponse{}
	for _, order := range orders {
		orderResponses = append(orderResponses, dto.OrderResponse{
			ID:        order.ID,
			Quote_ID:  order.QuoteFK,
			Status:    order.Status,
			CreatedAt: order.CreatedAt.Format(time.RFC3339),
			UpdatedAt: order.UpdatedAt.Format(time.RFC3339),
		})
	}
	return orderResponses, nil
}
