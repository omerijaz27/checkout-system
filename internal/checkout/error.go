package checkout

import (
	"errors"
	"fmt"
)

var (
	ErrEmptySKU = errors.New("SKU is empty")
)

// InvalidSKUError is returned when an unknown SKU is scanned
type InvalidSKUError struct {
	SKU string
}

func (e InvalidSKUError) Error() string {
	return fmt.Sprintf("invalid SKU: %s", e.SKU)
}
