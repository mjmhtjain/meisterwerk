package services

import (
	"github.com/mjmhtjain/meisterwerk/internal/models"
	"github.com/mjmhtjain/meisterwerk/internal/repository"
)

type OrderServiceI interface {
	CreateOrder(order models.Order) (models.Order, error)
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
