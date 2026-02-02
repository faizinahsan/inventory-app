package repositories

import (
	"context"
	"github.com/google/uuid"
	"inventory-app/internal/domain/entities"
)

// CategoryRepository defines the interface for category persistence operations
type CategoryRepository interface {
	Create(ctx context.Context, category *entities.Category) error
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Category, error)
	GetAll(ctx context.Context) ([]*entities.Category, error)
	GetByParentID(ctx context.Context, parentID uuid.UUID) ([]*entities.Category, error)
	GetRootCategories(ctx context.Context) ([]*entities.Category, error)
	Update(ctx context.Context, category *entities.Category) error
	Delete(ctx context.Context, id uuid.UUID) error
}
