package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"inventory-app/internal/domain/entities"
	"inventory-app/internal/domain/repositories"
	"inventory-app/internal/infrastructure/database"
)

type productRepository struct {
	db *database.DB
}

// NewProductRepository creates a new product repository
func NewProductRepository(db *database.DB) repositories.ProductRepository {
	return &productRepository{db: db}
}

// Create creates a new product
func (r *productRepository) Create(ctx context.Context, product *entities.Product) error {
	query := `
		INSERT INTO products (id, sku, name, description, category_id, price, cost, stock, min_stock, max_stock, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`

	_, err := r.db.ExecContext(ctx, query,
		product.ID, product.SKU, product.Name, product.Description, product.CategoryID,
		product.Price, product.Cost, product.Stock, product.MinStock, product.MaxStock,
		product.Status, product.CreatedAt, product.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}

	return nil
}

// GetByID retrieves a product by ID
func (r *productRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Product, error) {
	query := `
		SELECT id, sku, name, description, category_id, price, cost, stock, min_stock, max_stock, status, created_at, updated_at
		FROM products WHERE id = $1
	`

	product := &entities.Product{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&product.ID, &product.SKU, &product.Name, &product.Description, &product.CategoryID,
		&product.Price, &product.Cost, &product.Stock, &product.MinStock, &product.MaxStock,
		&product.Status, &product.CreatedAt, &product.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get product by ID: %w", err)
	}

	return product, nil
}

// GetBySKU retrieves a product by SKU
func (r *productRepository) GetBySKU(ctx context.Context, sku string) (*entities.Product, error) {
	query := `
		SELECT id, sku, name, description, category_id, price, cost, stock, min_stock, max_stock, status, created_at, updated_at
		FROM products WHERE sku = $1
	`

	product := &entities.Product{}
	err := r.db.QueryRowContext(ctx, query, sku).Scan(
		&product.ID, &product.SKU, &product.Name, &product.Description, &product.CategoryID,
		&product.Price, &product.Cost, &product.Stock, &product.MinStock, &product.MaxStock,
		&product.Status, &product.CreatedAt, &product.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get product by SKU: %w", err)
	}

	return product, nil
}

// GetAll retrieves all products with pagination
func (r *productRepository) GetAll(ctx context.Context, limit, offset int) ([]*entities.Product, error) {
	query := `
		SELECT id, sku, name, description, category_id, price, cost, stock, min_stock, max_stock, status, created_at, updated_at
		FROM products ORDER BY created_at DESC LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get all products: %w", err)
	}
	defer rows.Close()

	var products []*entities.Product
	for rows.Next() {
		product := &entities.Product{}
		err := rows.Scan(
			&product.ID, &product.SKU, &product.Name, &product.Description, &product.CategoryID,
			&product.Price, &product.Cost, &product.Stock, &product.MinStock, &product.MaxStock,
			&product.Status, &product.CreatedAt, &product.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}

	return products, nil
}

// GetByCategory retrieves products by category with pagination
func (r *productRepository) GetByCategory(ctx context.Context, categoryID uuid.UUID, limit, offset int) ([]*entities.Product, error) {
	query := `
		SELECT id, sku, name, description, category_id, price, cost, stock, min_stock, max_stock, status, created_at, updated_at
		FROM products WHERE category_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, categoryID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get products by category: %w", err)
	}
	defer rows.Close()

	var products []*entities.Product
	for rows.Next() {
		product := &entities.Product{}
		err := rows.Scan(
			&product.ID, &product.SKU, &product.Name, &product.Description, &product.CategoryID,
			&product.Price, &product.Cost, &product.Stock, &product.MinStock, &product.MaxStock,
			&product.Status, &product.CreatedAt, &product.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}

	return products, nil
}

// Update updates a product
func (r *productRepository) Update(ctx context.Context, product *entities.Product) error {
	query := `
		UPDATE products 
		SET sku = $2, name = $3, description = $4, category_id = $5, price = $6, cost = $7, 
		    stock = $8, min_stock = $9, max_stock = $10, status = $11, updated_at = $12
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query,
		product.ID, product.SKU, product.Name, product.Description, product.CategoryID,
		product.Price, product.Cost, product.Stock, product.MinStock, product.MaxStock,
		product.Status, product.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	return nil
}

// Delete deletes a product
func (r *productRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM products WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}

// GetLowStockProducts retrieves products with low stock
func (r *productRepository) GetLowStockProducts(ctx context.Context) ([]*entities.Product, error) {
	query := `
		SELECT id, sku, name, description, category_id, price, cost, stock, min_stock, max_stock, status, created_at, updated_at
		FROM products WHERE stock <= min_stock AND status = 'active' ORDER BY stock ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get low stock products: %w", err)
	}
	defer rows.Close()

	var products []*entities.Product
	for rows.Next() {
		product := &entities.Product{}
		err := rows.Scan(
			&product.ID, &product.SKU, &product.Name, &product.Description, &product.CategoryID,
			&product.Price, &product.Cost, &product.Stock, &product.MinStock, &product.MaxStock,
			&product.Status, &product.CreatedAt, &product.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}

	return products, nil
}

// Search searches for products
func (r *productRepository) Search(ctx context.Context, query string, limit, offset int) ([]*entities.Product, error) {
	searchQuery := `
		SELECT id, sku, name, description, category_id, price, cost, stock, min_stock, max_stock, status, created_at, updated_at
		FROM products 
		WHERE (name ILIKE $1 OR description ILIKE $1 OR sku ILIKE $1) AND status = 'active'
		ORDER BY name ASC LIMIT $2 OFFSET $3
	`

	searchTerm := "%" + strings.ToLower(query) + "%"
	rows, err := r.db.QueryContext(ctx, searchQuery, searchTerm, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to search products: %w", err)
	}
	defer rows.Close()

	var products []*entities.Product
	for rows.Next() {
		product := &entities.Product{}
		err := rows.Scan(
			&product.ID, &product.SKU, &product.Name, &product.Description, &product.CategoryID,
			&product.Price, &product.Cost, &product.Stock, &product.MinStock, &product.MaxStock,
			&product.Status, &product.CreatedAt, &product.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}

	return products, nil
}
