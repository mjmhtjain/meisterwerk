package services

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/mjmhtjain/meisterwerk/internal/config"
	"github.com/mjmhtjain/meisterwerk/internal/database"
	"github.com/mjmhtjain/meisterwerk/internal/dto"
	"github.com/mjmhtjain/meisterwerk/internal/models"
)

type QuoteServiceI interface {
	CreateQuote(quote dto.CreateQuoteRequest) (dto.QuoteResponse, error)
	GetQuote(id string) (dto.QuoteResponse, error)
	UpdateQuote(quote *models.Quote) error
	GetAllQuotes() ([]models.Quote, error)
}

type QuoteService struct {
	productService ProductServiceI
	db             *sql.DB
}

func NewQuoteService() QuoteServiceI {
	db, err := database.NewDBClient(config.NewDatabaseConfig())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	prodService := NewProductService()

	return &QuoteService{
		productService: prodService,
		db:             db,
	}
}

func (s *QuoteService) CreateQuote(quoteReq dto.CreateQuoteRequest) (dto.QuoteResponse, error) {
	// verify the products list
	for _, p := range quoteReq.ProductList {
		_, err := s.productService.GetProduct(p)
		if err != nil {
			// throw error
		}
	}

	quote := models.Quote{
		ID:           uuid.New().String(),
		Author:       quoteReq.Author,
		CustomerName: quoteReq.CustomerName,
		Status:       "created",
	}

	query := `
		INSERT INTO quote (id, author, customer_name, status)
		VALUES ($1, $2, $3, $4)
	`

	query2 := `
	INSERT INTO quote_product_map (id, quote_fk, product_fk)
	VALUES ($1, $2, $3)
`

	_, err := s.db.Exec(query, quote.ID, quote.Author, quote.CustomerName, quote.Status)
	if err != nil {
		return dto.QuoteResponse{}, err
	}

	for _, p := range quoteReq.ProductList {
		_, err = s.db.Exec(query2, uuid.New().String(), quote.ID, p)
		if err != nil {
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
	var quote models.Quote
	query := `
		SELECT id, author, customer_name, status
		FROM quote
		WHERE id = $1
	`

	err := s.db.QueryRow(query, id).Scan(
		&quote.ID,
		&quote.Author,
		&quote.CustomerName,
		&quote.Status,
	)
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

func (s *QuoteService) UpdateQuote(quote *models.Quote) error {
	query := `
		UPDATE quote
		SET author = $1, customer_name = $2, status = $3
		WHERE id = $4
	`

	_, err := s.db.Exec(query, quote.Author, quote.CustomerName, quote.Status, quote.ID)
	return err
}

func (s *QuoteService) GetAllQuotes() ([]models.Quote, error) {
	query := `
		SELECT id, author, customer_name, status
		FROM quote
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []models.Quote
	for rows.Next() {
		var quote models.Quote
		err := rows.Scan(
			&quote.ID,
			&quote.Author,
			&quote.CustomerName,
			&quote.Status,
		)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, quote)
	}

	return quotes, rows.Err()
}
