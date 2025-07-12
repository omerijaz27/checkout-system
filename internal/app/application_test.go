package app

import (
	"errors"
	"testing"
)

type MockCheckout struct {
	ScanFunc          func(sku string) error
	GetTotalPriceFunc func() (int, error)
}

func (m *MockCheckout) Scan(sku string) error {
	if m.ScanFunc != nil {
		return m.ScanFunc(sku)
	}
	return nil
}

func (m *MockCheckout) GetTotalPrice() (int, error) {
	if m.GetTotalPriceFunc != nil {
		return m.GetTotalPriceFunc()
	}
	return 0, nil
}

func Test_Run(t *testing.T) {
	mock := &MockCheckout{
		ScanFunc: func(sku string) error {
			return nil
		},
		GetTotalPriceFunc: func() (int, error) {
			return 95, nil
		},
	}

	app := &Application{checkout: mock}
	err := app.Run()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func Test_Run_FailsIfScanFails(t *testing.T) {
	mock := &MockCheckout{
		ScanFunc: func(sku string) error {
			if sku == "B" {
				return errors.New("scan error")
			}
			return nil
		},
		GetTotalPriceFunc: func() (int, error) {
			return 0, nil
		},
	}

	app := &Application{checkout: mock}
	err := app.Run()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
