package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// GenerateID generates a random ID
func GenerateID(prefix string) string {
	bytes := make([]byte, 8)
	rand.Read(bytes)
	return fmt.Sprintf("%s_%s", prefix, hex.EncodeToString(bytes))
}

// ValidateRequired checks if required fields are not empty
func ValidateRequired(fields map[string]string) error {
	var missing []string
	for field, value := range fields {
		if strings.TrimSpace(value) == "" {
			missing = append(missing, field)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("required fields are missing: %s", strings.Join(missing, ", "))
	}

	return nil
}

// FormatCurrency formats a float64 as currency
func FormatCurrency(amount float64) string {
	return fmt.Sprintf("$%.2f", amount)
}

// ParseDateRange parses date range strings
func ParseDateRange(startDate, endDate string) (time.Time, time.Time, error) {
	layout := "2006-01-02"

	start, err := time.Parse(layout, startDate)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid start date format: %v", err)
	}

	end, err := time.Parse(layout, endDate)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid end date format: %v", err)
	}

	if end.Before(start) {
		return time.Time{}, time.Time{}, fmt.Errorf("end date must be after start date")
	}

	return start, end, nil
}

// Contains checks if a slice contains a specific item
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// Paginate calculates pagination parameters
func Paginate(page, limit, total int) (offset int, hasNext bool) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset = (page - 1) * limit
	hasNext = offset+limit < total

	return offset, hasNext
}
