package repository

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/mjmhtjain/meisterwerk/internal/config"
	"github.com/mjmhtjain/meisterwerk/internal/database"
	"github.com/mjmhtjain/meisterwerk/internal/dto"
	"github.com/mjmhtjain/meisterwerk/internal/models"
)

type QuoteRepositoryI interface {
	Create(quote models.Quote) error
	CreateQuoteProductMap(quoteID string, productID string) error
	GetProductsByQuoteID(id string) ([]models.Product, error)
	GetByID(id string) (models.Quote, error)
	UpdateQuoteStatus(id string, status dto.QuoteStatus) error
	GetAll() ([]models.Quote, error)
}

type QuoteRepository struct {
	db *sql.DB
}

func NewQuoteRepository() QuoteRepositoryI {
	db, err := database.NewDBClient(config.NewDatabaseConfig())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return &QuoteRepository{
		db: db,
	}
}

func (r *QuoteRepository) Create(quote models.Quote) error {
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

func (r *QuoteRepository) GetByID(id string) (models.Quote, error) {
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
		return models.Quote{}, err
	}

	return quote, nil
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

func (r *QuoteRepository) GetProductsByQuoteID(id string) ([]models.Product, error) {
	query := `
		SELECT p.id, p.name, p.price, p.tax
		FROM product p
		INNER JOIN quote_product_map qpm ON p.id = qpm.product_fk
		WHERE qpm.quote_fk = $1
	`

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Tax)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, rows.Err()
}

func (r *QuoteRepository) UpdateQuoteStatus(id string, status dto.QuoteStatus) error {
	query := `
		UPDATE quote
		SET status = $1
		WHERE id = $2
	`
	_, err := r.db.Exec(query, status.String(), id)
	return err
}
