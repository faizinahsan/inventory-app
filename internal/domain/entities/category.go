package entities

import (
	"github.com/google/uuid"
	"time"
)

// Category represents a product category entity
type Category struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Description string     `json:"description" db:"description"`
	ParentID    *uuid.UUID `json:"parent_id" db:"parent_id"`
	Status      string     `json:"status" db:"status"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

// NewCategory creates a new category instance
func NewCategory(name, description string, parentID *uuid.UUID) *Category {
	return &Category{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		ParentID:    parentID,
		Status:      "active",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// IsSubCategory checks if this is a subcategory
func (c *Category) IsSubCategory() bool {
	return c.ParentID != nil
}

// Deactivate marks the category as inactive
func (c *Category) Deactivate() {
	c.Status = "inactive"
	c.UpdatedAt = time.Now()
}

// Activate marks the category as active
func (c *Category) Activate() {
	c.Status = "active"
	c.UpdatedAt = time.Now()
}
