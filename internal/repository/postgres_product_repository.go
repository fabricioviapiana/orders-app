package repository

import (
	"database/sql"

	"github.com/fabricioviapiana/orders-app/internal/domain"
)

type postgresProductRepository struct {
	db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) *postgresProductRepository {
	return &postgresProductRepository{
		db: db,
	}
}

func (r *postgresProductRepository) List() ([]domain.Product, error) {
	query := `SELECT id, name, price FROM products`
	rows, err := r.db.Query(query)
	if err != nil {
		return []domain.Product{}, nil
	}

	var products []domain.Product
	for rows.Next() {
		var product domain.Product

		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return products, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return products, err
	}

	return products, nil
}

func (r *postgresProductRepository) Create(name string, price float64) (domain.Product, error) {
	query := `INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id`
	var id string
	if err := r.db.QueryRow(query, name, price).Scan(&id); err != nil {
		return domain.Product{}, err
	}
	return domain.Product{
		ID:    id,
		Name:  name,
		Price: price,
	}, nil
}

func (r *postgresProductRepository) FindByID(id string) (domain.Product, error) {
	query := `SELECT id, name, price FROM products WHERE id = $1`
	var product domain.Product
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return product, err
	}
	return product, nil
}
