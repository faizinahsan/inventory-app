package entities

import (
	"github.com/google/uuid"
	"time"
)

// Transaction represents an inventory transaction entity
type Transaction struct {
	ID        uuid.UUID `json:"id" db:"id"`
	ProductID uuid.UUID `json:"product_id" db:"product_id"`
	Type      string    `json:"type" db:"type"` // "in", "out", "adjustment"
	Quantity  int       `json:"quantity" db:"quantity"`
	Reference string    `json:"reference" db:"reference"`
	Notes     string    `json:"notes" db:"notes"`
	CreatedBy uuid.UUID `json:"created_by" db:"created_by"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

const (
	TransactionTypeIn         = "in"
	TransactionTypeOut        = "out"
	TransactionTypeAdjustment = "adjustment"
)

// NewTransaction creates a new transaction instance
func NewTransaction(productID uuid.UUID, transactionType string, quantity int, reference, notes string, createdBy uuid.UUID) *Transaction {
	return &Transaction{
		ID:        uuid.New(),
		ProductID: productID,
		Type:      transactionType,
		Quantity:  quantity,
		Reference: reference,
		Notes:     notes,
		CreatedBy: createdBy,
		CreatedAt: time.Now(),
	}
}

// IsStockIn checks if this is a stock-in transaction
func (t *Transaction) IsStockIn() bool {
	return t.Type == TransactionTypeIn
}

// IsStockOut checks if this is a stock-out transaction
func (t *Transaction) IsStockOut() bool {
	return t.Type == TransactionTypeOut
}

// IsAdjustment checks if this is a stock adjustment transaction
func (t *Transaction) IsAdjustment() bool {
	return t.Type == TransactionTypeAdjustment
}
