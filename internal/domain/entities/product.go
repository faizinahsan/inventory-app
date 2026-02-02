package entities

import (
	"github.com/google/uuid"
	"time"
)

// Product represents a product entity in the inventory domain
type Product struct {
	ID          uuid.UUID `json:"id" db:"id"`
	SKU         string    `json:"sku" db:"sku"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CategoryID  uuid.UUID `json:"category_id" db:"category_id"`
	Price       float64   `json:"price" db:"price"`
	Cost        float64   `json:"cost" db:"cost"`
	Stock       int       `json:"stock" db:"stock"`
	MinStock    int       `json:"min_stock" db:"min_stock"`
	MaxStock    int       `json:"max_stock" db:"max_stock"`
	Status      string    `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// NewProduct creates a new product instance
func NewProduct(sku, name, description string, categoryID uuid.UUID, price, cost float64, minStock, maxStock int) *Product {
	return &Product{
		ID:          uuid.New(),
		SKU:         sku,
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
		Price:       price,
		Cost:        cost,
		Stock:       0,
		MinStock:    minStock,
		MaxStock:    maxStock,
		Status:      "active",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// IsLowStock checks if the product stock is below minimum threshold
func (p *Product) IsLowStock() bool {
	return p.Stock <= p.MinStock
}

// IsOverStock checks if the product stock exceeds maximum threshold
func (p *Product) IsOverStock() bool {
	return p.Stock >= p.MaxStock
}

// UpdateStock updates the product stock quantity
func (p *Product) UpdateStock(quantity int) error {
	if p.Stock+quantity < 0 {
		return ErrInsufficientStock
	}
	p.Stock += quantity
	p.UpdatedAt = time.Now()
	return nil
}

// Deactivate marks the product as inactive
func (p *Product) Deactivate() {
	p.Status = "inactive"
	p.UpdatedAt = time.Now()
}

// Activate marks the product as active
func (p *Product) Activate() {
	p.Status = "active"
	p.UpdatedAt = time.Now()
}
