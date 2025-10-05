package sql_storage

import (
	"database/sql"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
)

type ProductRepository struct {
	storage *SQLStorage
}

func (pr *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	product := &model.Product{}
	
	if err := pr.storage.db.QueryRow(
		"SELECT id, title, description, price, image, weight, amount FROM products WHERE id = $1",
		id,
	).Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.Image, &product.Weight, &product.Amount); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.ErrRecordNotFound
		}
		return nil, err
	}

	return product, nil
}

func (pr *ProductRepository) GetProducts() ([]*model.Product, error) {
	rows, err := pr.storage.db.Query(
		"SELECT id, title, description, price, image, weight, amount FROM products ORDER BY title",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*model.Product
	for rows.Next() {
		product := &model.Product{}
		if err := rows.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.Image, &product.Weight, &product.Amount); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}