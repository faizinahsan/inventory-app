package dto

import (
	"github.com/google/uuid"
	"time"
)

// TransactionRequest represents a transaction request
type TransactionRequest struct {
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	Type      string    `json:"type" binding:"required,oneof=in out adjustment"`
	Quantity  int       `json:"quantity" binding:"required"`
	Reference string    `json:"reference"`
	Notes     string    `json:"notes"`
}

// TransactionResponse represents a transaction response
type TransactionResponse struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	Type      string    `json:"type"`
	Quantity  int       `json:"quantity"`
	Reference string    `json:"reference"`
	Notes     string    `json:"notes"`
	CreatedBy uuid.UUID `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

// StockMovementRequest represents a stock movement request
type StockMovementRequest struct {
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required"`
	Reference string    `json:"reference"`
	Notes     string    `json:"notes"`
}

// StockAdjustmentRequest represents a stock adjustment request
type StockAdjustmentRequest struct {
	ProductID   uuid.UUID `json:"product_id" binding:"required"`
	NewQuantity int       `json:"new_quantity" binding:"required,min=0"`
	Notes       string    `json:"notes"`
}

// TransactionListResponse represents a paginated list of transactions
type TransactionListResponse struct {
	Transactions []TransactionResponse `json:"transactions"`
	Total        int                   `json:"total"`
	Page         int                   `json:"page"`
	Limit        int                   `json:"limit"`
}
