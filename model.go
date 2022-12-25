package go_product_api

import (
	"database/sql"
	"errors"
)

type Product struct {
	ID    int     `json:"ID"`
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
}

func (p *Product) getProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (p *Product) updateProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (p *Product) deleteProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (p *Product) createProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func getProducts(db *sql.DB) ([]Product, error) {
	return nil, errors.New("Not implemented")
}