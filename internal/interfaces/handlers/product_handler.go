package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"inventory-app/internal/application/dto"
	"inventory-app/internal/application/usecases"
)

// ProductHandler handles product-related HTTP requests
type ProductHandler struct {
	productUseCase usecases.ProductUseCase
}

// NewProductHandler creates a new product handler
func NewProductHandler(productUseCase usecases.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUseCase: productUseCase,
	}
}

// CreateProduct handles POST /products
func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var req dto.ProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	product, err := h.productUseCase.CreateProduct(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

// GetProduct handles GET /products/:id
func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid product ID"})
	}

	product, err := h.productUseCase.GetProduct(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(product)
}

// UpdateProduct handles PUT /products/:id
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid product ID"})
	}

	var req dto.ProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	product, err := h.productUseCase.UpdateProduct(c.Context(), id, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(product)
}

// DeleteProduct handles DELETE /products/:id
func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid product ID"})
	}

	err = h.productUseCase.DeleteProduct(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// ListProducts handles GET /products
func (h *ProductHandler) ListProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	products, err := h.productUseCase.ListProducts(c.Context(), page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(products)
}

// SearchProducts handles GET /products/search
func (h *ProductHandler) SearchProducts(c *fiber.Ctx) error {
	query := c.Query("q")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	products, err := h.productUseCase.SearchProducts(c.Context(), query, page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(products)
}

// GetLowStockProducts handles GET /products/low-stock
func (h *ProductHandler) GetLowStockProducts(c *fiber.Ctx) error {
	products, err := h.productUseCase.GetLowStockProducts(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(products)
}
