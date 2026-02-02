package valueobjects

import (
	"errors"
	"regexp"
	"strings"
)

// SKU represents a Stock Keeping Unit value object
type SKU struct {
	value string
}

// NewSKU creates a new SKU value object
func NewSKU(value string) (*SKU, error) {
	if err := validateSKU(value); err != nil {
		return nil, err
	}
	return &SKU{value: strings.ToUpper(value)}, nil
}

// Value returns the SKU value
func (s *SKU) Value() string {
	return s.value
}

// String implements the Stringer interface
func (s *SKU) String() string {
	return s.value
}

// Equals checks if two SKUs are equal
func (s *SKU) Equals(other *SKU) bool {
	return s.value == other.value
}

// validateSKU validates SKU format
func validateSKU(value string) error {
	if value == "" {
		return errors.New("SKU cannot be empty")
	}

	// SKU format: alphanumeric, dashes, underscores allowed, 3-20 characters
	matched, err := regexp.MatchString(`^[A-Za-z0-9_-]{3,20}$`, value)
	if err != nil {
		return err
	}

	if !matched {
		return errors.New("invalid SKU format: must be 3-20 characters, alphanumeric, dashes, and underscores only")
	}

	return nil
}
