// Package app wires everything together and runs the flow.
package app

import (
	"fmt"

	"checkout-system/internal/checkout"
)

// ICheckout is the interface for checkout.
type ICheckout interface {
	Scan(sku string) error
	GetTotalPrice() (int, error)
}

// Application struct holds everything needed to run the app.
type Application struct {
	checkout ICheckout
}

// NewApplication sets up the app with a basic checkout instance.
func NewApplication() *Application {
	return &Application{
		checkout: checkout.NewCheckout(),
	}
}

// Run simulates scanning some SKUs and prints the final total.
func (a *Application) Run() error {
	skus := []string{"B", "A", "B"}
	for _, sku := range skus {
		if err := a.checkout.Scan(sku); err != nil {
			return fmt.Errorf("failed to scan SKU %s: %w", sku, err)
		}
	}

	total, err := a.checkout.GetTotalPrice()
	if err != nil {
		return fmt.Errorf("failed to calculate total: %w", err)
	}

	fmt.Printf("total is: %d\n", total)
	return nil
}
