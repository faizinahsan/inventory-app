package usecases

import (
	"context"
	"github.com/google/uuid"
	"inventory-app/internal/application/dto"
	"inventory-app/internal/domain/entities"
	"inventory-app/internal/domain/repositories"
	"inventory-app/internal/domain/services"
)

// ProductUseCase handles product-related operations
type ProductUseCase interface {
	CreateProduct(ctx context.Context, req *dto.ProductRequest) (*dto.ProductResponse, error)
	GetProduct(ctx context.Context, id uuid.UUID) (*dto.ProductResponse, error)
	GetProductBySKU(ctx context.Context, sku string) (*dto.ProductResponse, error)
	UpdateProduct(ctx context.Context, id uuid.UUID, req *dto.ProductRequest) (*dto.ProductResponse, error)
	DeleteProduct(ctx context.Context, id uuid.UUID) error
	ListProducts(ctx context.Context, page, limit int) (*dto.ProductListResponse, error)
	SearchProducts(ctx context.Context, query string, page, limit int) (*dto.ProductListResponse, error)
	GetLowStockProducts(ctx context.Context) ([]dto.ProductResponse, error)
}

type productUseCase struct {
	productRepo      repositories.ProductRepository
	categoryRepo     repositories.CategoryRepository
	inventoryService services.InventoryService
}

// NewProductUseCase creates a new product use case
func NewProductUseCase(productRepo repositories.ProductRepository, categoryRepo repositories.CategoryRepository, inventoryService services.InventoryService) ProductUseCase {
	return &productUseCase{
		productRepo:      productRepo,
		categoryRepo:     categoryRepo,
		inventoryService: inventoryService,
	}
}

// CreateProduct creates a new product
func (uc *productUseCase) CreateProduct(ctx context.Context, req *dto.ProductRequest) (*dto.ProductResponse, error) {
	// Validate category exists
	_, err := uc.categoryRepo.GetByID(ctx, req.CategoryID)
	if err != nil {
		return nil, err
	}

	// Check if SKU already exists
	existingProduct, _ := uc.productRepo.GetBySKU(ctx, req.SKU)
	if existingProduct != nil {
		return nil, entities.ErrDuplicateSKU
	}

	// Create product entity
	product := entities.NewProduct(
		req.SKU,
		req.Name,
		req.Description,
		req.CategoryID,
		req.Price,
		req.Cost,
		req.MinStock,
		req.MaxStock,
	)

	// Save product
	err = uc.productRepo.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return uc.entityToResponse(product), nil
}

// GetProduct retrieves a product by ID
func (uc *productUseCase) GetProduct(ctx context.Context, id uuid.UUID) (*dto.ProductResponse, error) {
	product, err := uc.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, entities.ErrProductNotFound
	}

	return uc.entityToResponse(product), nil
}

// GetProductBySKU retrieves a product by SKU
func (uc *productUseCase) GetProductBySKU(ctx context.Context, sku string) (*dto.ProductResponse, error) {
	product, err := uc.productRepo.GetBySKU(ctx, sku)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, entities.ErrProductNotFound
	}

	return uc.entityToResponse(product), nil
}

// UpdateProduct updates an existing product
func (uc *productUseCase) UpdateProduct(ctx context.Context, id uuid.UUID, req *dto.ProductRequest) (*dto.ProductResponse, error) {
	// Get existing product
	product, err := uc.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, entities.ErrProductNotFound
	}

	// Validate category exists
	_, err = uc.categoryRepo.GetByID(ctx, req.CategoryID)
	if err != nil {
		return nil, err
	}

	// Check if SKU already exists (excluding current product)
	if req.SKU != product.SKU {
		existingProduct, _ := uc.productRepo.GetBySKU(ctx, req.SKU)
		if existingProduct != nil && existingProduct.ID != id {
			return nil, entities.ErrDuplicateSKU
		}
	}

	// Update product fields
	product.SKU = req.SKU
	product.Name = req.Name
	product.Description = req.Description
	product.CategoryID = req.CategoryID
	product.Price = req.Price
	product.Cost = req.Cost
	product.MinStock = req.MinStock
	product.MaxStock = req.MaxStock

	// Save updated product
	err = uc.productRepo.Update(ctx, product)
	if err != nil {
		return nil, err
	}

	return uc.entityToResponse(product), nil
}

// DeleteProduct deletes a product
func (uc *productUseCase) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	return uc.productRepo.Delete(ctx, id)
}

// ListProducts retrieves a paginated list of products
func (uc *productUseCase) ListProducts(ctx context.Context, page, limit int) (*dto.ProductListResponse, error) {
	offset := (page - 1) * limit
	products, err := uc.productRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	response := &dto.ProductListResponse{
		Products: make([]dto.ProductResponse, len(products)),
		Total:    len(products), // In a real implementation, you'd get the total count separately
		Page:     page,
		Limit:    limit,
	}

	for i, product := range products {
		response.Products[i] = *uc.entityToResponse(product)
	}

	return response, nil
}

// SearchProducts searches for products
func (uc *productUseCase) SearchProducts(ctx context.Context, query string, page, limit int) (*dto.ProductListResponse, error) {
	offset := (page - 1) * limit
	products, err := uc.productRepo.Search(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	response := &dto.ProductListResponse{
		Products: make([]dto.ProductResponse, len(products)),
		Total:    len(products), // In a real implementation, you'd get the total count separately
		Page:     page,
		Limit:    limit,
	}

	for i, product := range products {
		response.Products[i] = *uc.entityToResponse(product)
	}

	return response, nil
}

// GetLowStockProducts retrieves products with low stock
func (uc *productUseCase) GetLowStockProducts(ctx context.Context) ([]dto.ProductResponse, error) {
	products, err := uc.inventoryService.GetLowStockAlerts(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]dto.ProductResponse, len(products))
	for i, product := range products {
		response[i] = *uc.entityToResponse(product)
	}

	return response, nil
}

// entityToResponse converts product entity to response DTO
func (uc *productUseCase) entityToResponse(product *entities.Product) *dto.ProductResponse {
	return &dto.ProductResponse{
		ID:          product.ID,
		SKU:         product.SKU,
		Name:        product.Name,
		Description: product.Description,
		CategoryID:  product.CategoryID,
		Price:       product.Price,
		Cost:        product.Cost,
		Stock:       product.Stock,
		MinStock:    product.MinStock,
		MaxStock:    product.MaxStock,
		Status:      product.Status,
		IsLowStock:  product.IsLowStock(),
		IsOverStock: product.IsOverStock(),
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
