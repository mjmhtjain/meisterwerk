package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/mjmhtjain/meisterwerk/internal/models"
)

type QuoteRepositoryI interface {
	Create(quote *models.Quote) error
	CreateQuoteProductMap(quoteID string, productID string) error
	GetByID(id string) (*models.Quote, error)
	Update(quote *models.Quote) error
	GetAll() ([]models.Quote, error)
}

type QuoteRepository struct {
	db *sql.DB
}

func NewQuoteRepository(db *sql.DB) QuoteRepositoryI {
	return &QuoteRepository{
		db: db,
	}
}

func (r *QuoteRepository) Create(quote *models.Quote) error {
	query := `
		INSERT INTO quote (id, author, customer_name, status)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.Exec(query, quote.ID, quote.Author, quote.CustomerName, quote.Status)
	return err
}

func (r *QuoteRepository) CreateQuoteProductMap(quoteID string, productID string) error {
	query := `
		INSERT INTO quote_product_map (id, quote_fk, product_fk)
		VALUES ($1, $2, $3)
	`
	_, err := r.db.Exec(query, uuid.New().String(), quoteID, productID)
	return err
}

func (r *QuoteRepository) GetByID(id string) (*models.Quote, error) {
	var quote models.Quote
	query := `
		SELECT id, author, customer_name, status
		FROM quote
		WHERE id = $1
	`

	err := r.db.QueryRow(query, id).Scan(
		&quote.ID,
		&quote.Author,
		&quote.CustomerName,
		&quote.Status,
	)
	if err != nil {
		return nil, err
	}

	return &quote, nil
}

func (r *QuoteRepository) Update(quote *models.Quote) error {
	query := `
		UPDATE quote
		SET author = $1, customer_name = $2, status = $3
		WHERE id = $4
	`
	_, err := r.db.Exec(query, quote.Author, quote.CustomerName, quote.Status, quote.ID)
	return err
}

func (r *QuoteRepository) GetAll() ([]models.Quote, error) {
	query := `
		SELECT id, author, customer_name, status
		FROM quote
	`

	rows, err := r.db.Query(query)
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
