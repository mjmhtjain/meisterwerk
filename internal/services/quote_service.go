package services

import (
	"log"

	"github.com/google/uuid"
	"github.com/mjmhtjain/meisterwerk/internal/config"
	"github.com/mjmhtjain/meisterwerk/internal/database"
	"github.com/mjmhtjain/meisterwerk/internal/dto"
	"github.com/mjmhtjain/meisterwerk/internal/models"
	"gorm.io/gorm"
)

type QuoteServiceI interface {
	CreateQuote(quote dto.CreateQuoteRequest) (dto.QuoteResponse, error)
	GetQuote(id string) (dto.QuoteResponse, error)
	UpdateQuote(quote *models.Quote) error
	GetAllQuotes() ([]models.Quote, error)
}

type QuoteService struct {
	db *gorm.DB
}

func NewQuoteService() QuoteServiceI {
	// Initialize database
	db, err := database.NewDBClient(config.NewDatabaseConfig())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return &QuoteService{
		db: db,
	}
}

func (s *QuoteService) CreateQuote(req dto.CreateQuoteRequest) (dto.QuoteResponse, error) {
	quote := models.Quote{
		ID:           uuid.New().String(),
		Author:       req.Author,
		CustomerName: req.CustomerName,
		Status:       "created",
	}

	if err := s.db.Create(quote).Error; err != nil {
		return dto.QuoteResponse{}, err
	}

	return dto.QuoteResponse{
		ID:           quote.ID,
		Author:       quote.Author,
		CustomerName: quote.CustomerName,
		Status:       quote.Status,
	}, nil
}

func (s *QuoteService) GetQuote(id string) (dto.QuoteResponse, error) {
	var quote models.Quote
	if err := s.db.Where("id = ?", id).First(&quote).Error; err != nil {
		return dto.QuoteResponse{}, err
	}

	return dto.QuoteResponse{
		ID:           quote.ID,
		Author:       quote.Author,
		CustomerName: quote.CustomerName,
		Status:       quote.Status,
	}, nil
}

func (s *QuoteService) UpdateQuote(quote *models.Quote) error {
	return s.db.Save(quote).Error
}

func (s *QuoteService) GetAllQuotes() ([]models.Quote, error) {
	var quotes []models.Quote
	if err := s.db.Find(&quotes).Error; err != nil {
		return nil, err
	}
	return quotes, nil
}
