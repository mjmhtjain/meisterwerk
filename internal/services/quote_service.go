package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/mjmhtjain/meisterwerk/internal/dto"
	"github.com/mjmhtjain/meisterwerk/internal/models"
	"github.com/mjmhtjain/meisterwerk/internal/repository"
)

type QuoteServiceI interface {
	CreateQuote(quote dto.CreateQuoteRequest) (dto.QuoteResponse, error)
	GetQuote(id string) (dto.QuoteResponse, error)
	UpdateQuoteStatus(id string, status dto.QuoteStatus) error
}

type QuoteService struct {
	productService ProductServiceI
	quoteRepo      repository.QuoteRepositoryI
	orderService   OrderServiceI
}

func NewQuoteService() QuoteServiceI {

	return &QuoteService{
		productService: NewProductService(),
		quoteRepo:      repository.NewQuoteRepository(),
		orderService:   NewOrderService(),
	}
}

func (s *QuoteService) CreateQuote(quoteReq dto.CreateQuoteRequest) (dto.QuoteResponse, error) {
	// verify the products list
	products := []dto.ProductResponse{}

	for _, p := range quoteReq.ProductList {
		product, err := s.productService.GetProduct(p)
		if err != nil {
			return dto.QuoteResponse{}, err
		}
		products = append(products, product)
	}

	quote := models.Quote{
		ID:           uuid.New().String(),
		Author:       quoteReq.Author,
		CustomerName: quoteReq.CustomerName,
		Status:       "created",
	}

	if err := s.quoteRepo.Create(quote); err != nil {
		return dto.QuoteResponse{}, err
	}

	for _, p := range quoteReq.ProductList {
		if err := s.quoteRepo.CreateQuoteProductMap(quote.ID, p); err != nil {
			return dto.QuoteResponse{}, err
		}
	}

	// calculate the total price and tax
	totalPrice := 0.0
	totalTax := 0.0

	for _, p := range products {
		totalPrice += p.Price
		totalTax += p.Price * (p.Tax / 100)
	}

	return dto.QuoteResponse{
		ID:           quote.ID,
		Author:       quote.Author,
		CustomerName: quote.CustomerName,
		ProductList:  products,
		TotalPrice:   totalPrice,
		TotalTax:     totalTax,
		Status:       quote.Status,
	}, nil
}

func (s *QuoteService) GetQuote(id string) (dto.QuoteResponse, error) {
	quote, err := s.quoteRepo.GetByID(id)
	if err != nil {
		return dto.QuoteResponse{}, err
	}

	products, err := s.quoteRepo.GetProductsByQuoteID(id)
	if err != nil {
		return dto.QuoteResponse{}, err
	}

	productsResponse := []dto.ProductResponse{}

	for _, p := range products {
		productsResponse = append(productsResponse, dto.ProductResponse{
			ID:    p.ID,
			Name:  p.Name,
			Price: p.Price,
			Tax:   p.Tax,
		})
	}

	// calculate the total price and tax
	totalPrice := 0.0
	totalTax := 0.0

	for _, p := range products {
		totalPrice += p.Price
		totalTax += p.Price * (p.Tax / 100)
	}

	quoteResponse := dto.QuoteResponse{
		ID:           quote.ID,
		Author:       quote.Author,
		CustomerName: quote.CustomerName,
		ProductList:  productsResponse,
		TotalPrice:   totalPrice,
		TotalTax:     totalTax,
		Status:       quote.Status,
	}

	return quoteResponse, nil
}

func (s *QuoteService) GetAllQuotes() ([]models.Quote, error) {
	return s.quoteRepo.GetAll()
}

func (s *QuoteService) UpdateQuoteStatus(id string, status dto.QuoteStatus) error {
	err := s.quoteRepo.UpdateQuoteStatus(id, status)
	if err != nil {
		return err
	}

	// if the quote is accepted, create an order
	if status == dto.QuoteStatusAccepted {
		order := models.Order{
			ID:        uuid.New().String(),
			Status:    dto.OrderStatusCreated.String(),
			QuoteFK:   id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		_, err := s.orderService.CreateOrder(order)
		if err != nil {
			return err
		}
	}

	return nil
}
