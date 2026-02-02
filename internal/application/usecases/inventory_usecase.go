package usecases

import (
	"context"
	"github.com/google/uuid"
	"inventory-app/internal/application/dto"
	"inventory-app/internal/domain/entities"
	"inventory-app/internal/domain/repositories"
	"inventory-app/internal/domain/services"
)

// InventoryUseCase handles inventory-related operations
type InventoryUseCase interface {
	StockIn(ctx context.Context, req *dto.StockMovementRequest, userID uuid.UUID) error
	StockOut(ctx context.Context, req *dto.StockMovementRequest, userID uuid.UUID) error
	AdjustStock(ctx context.Context, req *dto.StockAdjustmentRequest, userID uuid.UUID) error
	GetTransactionHistory(ctx context.Context, productID uuid.UUID, page, limit int) (*dto.TransactionListResponse, error)
	GetAllTransactions(ctx context.Context, page, limit int) (*dto.TransactionListResponse, error)
}

type inventoryUseCase struct {
	inventoryService services.InventoryService
	transactionRepo  repositories.TransactionRepository
}

// NewInventoryUseCase creates a new inventory use case
func NewInventoryUseCase(inventoryService services.InventoryService, transactionRepo repositories.TransactionRepository) InventoryUseCase {
	return &inventoryUseCase{
		inventoryService: inventoryService,
		transactionRepo:  transactionRepo,
	}
}

// StockIn processes incoming stock
func (uc *inventoryUseCase) StockIn(ctx context.Context, req *dto.StockMovementRequest, userID uuid.UUID) error {
	return uc.inventoryService.ProcessStockIn(ctx, req.ProductID, req.Quantity, req.Reference, req.Notes, userID)
}

// StockOut processes outgoing stock
func (uc *inventoryUseCase) StockOut(ctx context.Context, req *dto.StockMovementRequest, userID uuid.UUID) error {
	return uc.inventoryService.ProcessStockOut(ctx, req.ProductID, req.Quantity, req.Reference, req.Notes, userID)
}

// AdjustStock adjusts stock to a specific quantity
func (uc *inventoryUseCase) AdjustStock(ctx context.Context, req *dto.StockAdjustmentRequest, userID uuid.UUID) error {
	return uc.inventoryService.AdjustStock(ctx, req.ProductID, req.NewQuantity, req.Notes, userID)
}

// GetTransactionHistory retrieves transaction history for a product
func (uc *inventoryUseCase) GetTransactionHistory(ctx context.Context, productID uuid.UUID, page, limit int) (*dto.TransactionListResponse, error) {
	offset := (page - 1) * limit
	transactions, err := uc.transactionRepo.GetByProductID(ctx, productID, limit, offset)
	if err != nil {
		return nil, err
	}

	response := &dto.TransactionListResponse{
		Transactions: make([]dto.TransactionResponse, len(transactions)),
		Total:        len(transactions), // In a real implementation, you'd get the total count separately
		Page:         page,
		Limit:        limit,
	}

	for i, transaction := range transactions {
		response.Transactions[i] = uc.entityToResponse(transaction)
	}

	return response, nil
}

// GetAllTransactions retrieves all transactions
func (uc *inventoryUseCase) GetAllTransactions(ctx context.Context, page, limit int) (*dto.TransactionListResponse, error) {
	offset := (page - 1) * limit
	transactions, err := uc.transactionRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	response := &dto.TransactionListResponse{
		Transactions: make([]dto.TransactionResponse, len(transactions)),
		Total:        len(transactions), // In a real implementation, you'd get the total count separately
		Page:         page,
		Limit:        limit,
	}

	for i, transaction := range transactions {
		response.Transactions[i] = uc.entityToResponse(transaction)
	}

	return response, nil
}

// entityToResponse converts transaction entity to response DTO
func (uc *inventoryUseCase) entityToResponse(transaction *entities.Transaction) dto.TransactionResponse {
	return dto.TransactionResponse{
		ID:        transaction.ID,
		ProductID: transaction.ProductID,
		Type:      transaction.Type,
		Quantity:  transaction.Quantity,
		Reference: transaction.Reference,
		Notes:     transaction.Notes,
		CreatedBy: transaction.CreatedBy,
		CreatedAt: transaction.CreatedAt,
	}
}
