package services

import (
	"context"
	"github.com/google/uuid"
	"inventory-app/internal/domain/entities"
	"inventory-app/internal/domain/repositories"
)

// InventoryService handles inventory-related business logic
type InventoryService interface {
	ProcessStockIn(ctx context.Context, productID uuid.UUID, quantity int, reference, notes string, userID uuid.UUID) error
	ProcessStockOut(ctx context.Context, productID uuid.UUID, quantity int, reference, notes string, userID uuid.UUID) error
	AdjustStock(ctx context.Context, productID uuid.UUID, newQuantity int, notes string, userID uuid.UUID) error
	TransferStock(ctx context.Context, fromProductID, toProductID uuid.UUID, quantity int, reference, notes string, userID uuid.UUID) error
	GetLowStockAlerts(ctx context.Context) ([]*entities.Product, error)
}

type inventoryService struct {
	productRepo     repositories.ProductRepository
	transactionRepo repositories.TransactionRepository
}

// NewInventoryService creates a new inventory service
func NewInventoryService(productRepo repositories.ProductRepository, transactionRepo repositories.TransactionRepository) InventoryService {
	return &inventoryService{
		productRepo:     productRepo,
		transactionRepo: transactionRepo,
	}
}

// ProcessStockIn processes incoming stock
func (s *inventoryService) ProcessStockIn(ctx context.Context, productID uuid.UUID, quantity int, reference, notes string, userID uuid.UUID) error {
	product, err := s.productRepo.GetByID(ctx, productID)
	if err != nil {
		return err
	}

	if product == nil {
		return entities.ErrProductNotFound
	}

	if quantity <= 0 {
		return entities.ErrInvalidQuantity
	}

	// Update product stock
	err = product.UpdateStock(quantity)
	if err != nil {
		return err
	}

	// Create transaction record
	transaction := entities.NewTransaction(productID, entities.TransactionTypeIn, quantity, reference, notes, userID)

	// Save transaction
	if err := s.transactionRepo.Create(ctx, transaction); err != nil {
		return err
	}

	// Update product
	return s.productRepo.Update(ctx, product)
}

// ProcessStockOut processes outgoing stock
func (s *inventoryService) ProcessStockOut(ctx context.Context, productID uuid.UUID, quantity int, reference, notes string, userID uuid.UUID) error {
	product, err := s.productRepo.GetByID(ctx, productID)
	if err != nil {
		return err
	}

	if product == nil {
		return entities.ErrProductNotFound
	}

	if quantity <= 0 {
		return entities.ErrInvalidQuantity
	}

	// Check if sufficient stock
	if product.Stock < quantity {
		return entities.ErrInsufficientStock
	}

	// Update product stock
	err = product.UpdateStock(-quantity)
	if err != nil {
		return err
	}

	// Create transaction record
	transaction := entities.NewTransaction(productID, entities.TransactionTypeOut, quantity, reference, notes, userID)

	// Save transaction
	if err := s.transactionRepo.Create(ctx, transaction); err != nil {
		return err
	}

	// Update product
	return s.productRepo.Update(ctx, product)
}

// AdjustStock adjusts stock to a specific quantity
func (s *inventoryService) AdjustStock(ctx context.Context, productID uuid.UUID, newQuantity int, notes string, userID uuid.UUID) error {
	product, err := s.productRepo.GetByID(ctx, productID)
	if err != nil {
		return err
	}

	if product == nil {
		return entities.ErrProductNotFound
	}

	if newQuantity < 0 {
		return entities.ErrInvalidQuantity
	}

	// Calculate adjustment quantity
	adjustmentQuantity := newQuantity - product.Stock

	// Update product stock
	product.Stock = newQuantity

	// Create transaction record
	transaction := entities.NewTransaction(productID, entities.TransactionTypeAdjustment, adjustmentQuantity, "", notes, userID)

	// Save transaction
	if err := s.transactionRepo.Create(ctx, transaction); err != nil {
		return err
	}

	// Update product
	return s.productRepo.Update(ctx, product)
}

// TransferStock transfers stock between products (placeholder implementation)
func (s *inventoryService) TransferStock(ctx context.Context, fromProductID, toProductID uuid.UUID, quantity int, reference, notes string, userID uuid.UUID) error {
	// This is a simplified implementation
	// In a real scenario, you might need more complex business logic

	// Process stock out from source product
	err := s.ProcessStockOut(ctx, fromProductID, quantity, reference, "Transfer out: "+notes, userID)
	if err != nil {
		return err
	}

	// Process stock in to destination product
	return s.ProcessStockIn(ctx, toProductID, quantity, reference, "Transfer in: "+notes, userID)
}

// GetLowStockAlerts retrieves products with low stock
func (s *inventoryService) GetLowStockAlerts(ctx context.Context) ([]*entities.Product, error) {
	return s.productRepo.GetLowStockProducts(ctx)
}
