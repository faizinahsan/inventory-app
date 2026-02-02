package entities

import "errors"

// Domain errors
var (
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrProductNotFound   = errors.New("product not found")
	ErrCategoryNotFound  = errors.New("category not found")
	ErrInvalidSKU        = errors.New("invalid SKU")
	ErrInvalidQuantity   = errors.New("invalid quantity")
	ErrDuplicateSKU      = errors.New("duplicate SKU")
)
