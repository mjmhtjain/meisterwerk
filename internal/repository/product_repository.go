package repository

import (
	"database/sql"
	"log"

	"github.com/mjmhtjain/meisterwerk/internal/config"
	"github.com/mjmhtjain/meisterwerk/internal/database"
	"github.com/mjmhtjain/meisterwerk/internal/models"
)

type ProductRepositoryI interface {
	GetAll() ([]models.Product, error)
	GetByID(id string) (*models.Product, error)
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository() ProductRepositoryI {
	db, err := database.NewDBClient(config.NewDatabaseConfig())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	rows, err := r.db.Query("SELECT id, name, price, tax FROM product")
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

	return products, nil
}

func (r *ProductRepository) GetByID(id string) (*models.Product, error) {
	var product models.Product
	err := r.db.QueryRow("SELECT id, name, price, tax FROM product WHERE id = $1", id).
		Scan(&product.ID, &product.Name, &product.Price, &product.Tax)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}
