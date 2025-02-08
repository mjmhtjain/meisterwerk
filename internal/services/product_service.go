package services

import (
	"database/sql"
	"log"

	"github.com/mjmhtjain/meisterwerk/internal/config"
	"github.com/mjmhtjain/meisterwerk/internal/database"
	"github.com/mjmhtjain/meisterwerk/internal/dto"
	"github.com/mjmhtjain/meisterwerk/internal/models"
)

type ProductServiceI interface {
	GetAllProducts() ([]dto.ProductResponse, error)
}

type ProductService struct {
	db *sql.DB
}

func NewProductService() ProductServiceI {
	db, err := database.NewDBClient(config.NewDatabaseConfig())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return &ProductService{
		db: db,
	}
}

func (s *ProductService) GetAllProducts() ([]dto.ProductResponse, error) {
	rows, err := s.db.Query("SELECT id, name, price, tax FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []dto.ProductResponse{}
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Tax)
		if err != nil {
			return nil, err
		}
		products = append(products, dto.ProductResponse{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
			Tax:   product.Tax,
		})
	}

	return products, nil
}
