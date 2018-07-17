package db

import "time"

type Product struct {
	CreatedAt   time.Time   `json:"-"`
	ID          string      `gorm:"primary_key" json:"id"`
	Name        string      `json:"name"`
	ReleaseDate time.Time   `json:"release_date"`
	Rates       []*Rate     `json:"rates"`
	Posters     []*Poster   `json:"posters"`
	Prices      []*Price    `json:"price"`
	Discounts   []*Discount `json:"discounts"`
}

type ProductMgr interface {
	CreateProduct(product *Product) error
	EnsureProductExists(product *Product) (*Product, error)
}

func (mgr *AppDatabaseMgr) CreateProduct(product *Product) error {
	panic("not implemented")
}

func (mgr *AppDatabaseMgr) EnsureProductExists(product *Product) (*Product, error) {
	panic("not implemented")
}
