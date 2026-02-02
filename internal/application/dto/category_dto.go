package dto

import (
	"github.com/google/uuid"
	"time"
)

// CategoryRequest represents a category creation/update request
type CategoryRequest struct {
	Name        string     `json:"name" binding:"required"`
	Description string     `json:"description"`
	ParentID    *uuid.UUID `json:"parent_id"`
}

// CategoryResponse represents a category response
type CategoryResponse struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ParentID    *uuid.UUID `json:"parent_id"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// CategoryTreeResponse represents a hierarchical category response
type CategoryTreeResponse struct {
	ID          uuid.UUID              `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Status      string                 `json:"status"`
	Children    []CategoryTreeResponse `json:"children"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}
