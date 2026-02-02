package dto

import (
	"github.com/google/uuid"
	"time"
)

// ProductRequest represents a product creation/update request
type ProductRequest struct {
	SKU         string    `json:"sku" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	CategoryID  uuid.UUID `json:"category_id" binding:"required"`
	Price       float64   `json:"price" binding:"required,min=0"`
	Cost        float64   `json:"cost" binding:"required,min=0"`
	MinStock    int       `json:"min_stock" binding:"min=0"`
	MaxStock    int       `json:"max_stock" binding:"min=0"`
}

// ProductResponse represents a product response
type ProductResponse struct {
	ID          uuid.UUID `json:"id"`
	SKU         string    `json:"sku"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryID  uuid.UUID `json:"category_id"`
	Price       float64   `json:"price"`
	Cost        float64   `json:"cost"`
	Stock       int       `json:"stock"`
	MinStock    int       `json:"min_stock"`
	MaxStock    int       `json:"max_stock"`
	Status      string    `json:"status"`
	IsLowStock  bool      `json:"is_low_stock"`
	IsOverStock bool      `json:"is_over_stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ProductListResponse represents a paginated list of products
type ProductListResponse struct {
	Products []ProductResponse `json:"products"`
	Total    int               `json:"total"`
	Page     int               `json:"page"`
	Limit    int               `json:"limit"`
}
