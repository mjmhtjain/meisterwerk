package repository

import (
	"database/sql"
	"log"

	"github.com/mjmhtjain/meisterwerk/internal/config"
	"github.com/mjmhtjain/meisterwerk/internal/database"
	"github.com/mjmhtjain/meisterwerk/internal/models"
)

type OrderRepositoryI interface {
	CreateOrder(order models.Order) (models.Order, error)
	GetAll() ([]models.Order, error)
}

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository() *OrderRepository {
	db, err := database.NewDBClient(config.NewDatabaseConfig())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) CreateOrder(order models.Order) (models.Order, error) {
	query := `
		INSERT INTO "order" (id, status, quote_fk, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(query, order.ID, order.Status, order.QuoteFK, order.CreatedAt, order.UpdatedAt)
	return order, err
}

func (r *OrderRepository) GetAll() ([]models.Order, error) {
	query := `
		SELECT id, status, quote_fk, created_at, updated_at FROM "order"
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []models.Order{}
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.Status, &order.QuoteFK, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
