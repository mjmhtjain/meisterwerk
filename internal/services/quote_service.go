package services

import (
	"log"

	"github.com/google/uuid"
	"github.com/mjmhtjain/meisterwerk/internal/config"
	"github.com/mjmhtjain/meisterwerk/internal/database"
	"github.com/mjmhtjain/meisterwerk/internal/dto"
	"github.com/mjmhtjain/meisterwerk/internal/models"
	"github.com/mjmhtjain/meisterwerk/internal/repository"
)

type QuoteServiceI interface {
	CreateQuote(quote dto.CreateQuoteRequest) (dto.QuoteResponse, error)
	GetQuote(id string) (dto.QuoteResponse, error)
}

type QuoteService struct {
	productService ProductServiceI
	quoteRepo      repository.QuoteRepositoryI
}

func NewQuoteService() QuoteServiceI {
	db, err := database.NewDBClient(config.NewDatabaseConfig())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return &QuoteService{
		productService: NewProductService(),
		quoteRepo:      repository.NewQuoteRepository(db),
	}
}

func (s *QuoteService) CreateQuote(quoteReq dto.CreateQuoteRequest) (dto.QuoteResponse, error) {
	// verify the products list
	for _, p := range quoteReq.ProductList {
		_, err := s.productService.GetProduct(p)
		if err != nil {
			return dto.QuoteResponse{}, err
		}
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

	return dto.QuoteResponse{
		ID:           quote.ID,
		Author:       quote.Author,
		CustomerName: quote.CustomerName,
		Status:       quote.Status,
	}, nil
}

func (s *QuoteService) GetQuote(id string) (dto.QuoteResponse, error) {
	quote, err := s.quoteRepo.GetByID(id)
	if err != nil {
		return dto.QuoteResponse{}, err
	}

	return dto.QuoteResponse{
		ID:           quote.ID,
		Author:       quote.Author,
		CustomerName: quote.CustomerName,
		Status:       quote.Status,
	}, nil
}

func (s *QuoteService) GetAllQuotes() ([]models.Quote, error) {
	return s.quoteRepo.GetAll()
}
