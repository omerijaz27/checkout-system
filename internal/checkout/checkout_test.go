package checkout

import (
	"errors"
	"testing"
)

func newTestCheckout(t *testing.T) *Checkout {
	t.Helper()
	return NewCheckout()
}

func Test_Scan_ValidSKU(t *testing.T) {
	co := newTestCheckout(t)

	err := co.Scan("A")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = co.Scan("A")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if co.scannedItems["A"] != 2 {
		t.Errorf("expected %d items, got %d", 2, co.scannedItems["A"])
	}
}

func Test_Scan_InvalidSKU(t *testing.T) {
	co := newTestCheckout(t)

	err := co.Scan("Z")
	if err == nil {
		t.Fatalf("expected error for invalid SKU, got none")
	}

	var skuErr InvalidSKUError
	if !errors.As(err, &skuErr) {
		t.Errorf("expected InvalidSKUError, got %v", err)
	}
}

func Test_Scan_EmptySKU(t *testing.T) {
	co := newTestCheckout(t)

	err := co.Scan("")
	if err == nil {
		t.Fatalf("expected error for empty SKU, got none")
	}

	if !errors.Is(err, ErrEmptySKU) {
		t.Errorf("expected ErrEmptySKU, got %v", err)
	}
}

func Test_Scan_MultipleSKUs(t *testing.T) {
	co := newTestCheckout(t)

	skus := []string{"A", "B", "A", "C", "B", "D"}

	for _, sku := range skus {
		err := co.Scan(sku)
		if err != nil {
			t.Fatalf("unexpected error while scanning %s: %v", sku, err)
		}
	}

	expected := map[string]int{
		"A": 2,
		"B": 2,
		"C": 1,
		"D": 1,
	}

	for sku, count := range expected {
		if co.scannedItems[sku] != count {
			t.Errorf("expected %d of %s, got %d", count, sku, co.scannedItems[sku])
		}
	}
}

func Test_GetTotalPrice_EmptyCart(t *testing.T) {
	co := newTestCheckout(t)

	total, err := co.GetTotalPrice()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if total != 0 {
		t.Errorf("expected total 0, got %d", total)
	}
}

func Test_GetTotalPrice_SingleItem(t *testing.T) {
	co := newTestCheckout(t)
	err := co.Scan("C")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	total, err := co.GetTotalPrice()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if total != 20 {
		t.Errorf("expected total 20, got %d", total)
	}
}

func Test_GetTotalPrice_WithOffer(t *testing.T) {
	co := newTestCheckout(t)
	skus := []string{"A", "A", "A"}
	for _, sku := range skus {
		err := co.Scan(sku)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	}

	total, err := co.GetTotalPrice()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if total != 130 {
		t.Errorf("expected total 130, got %d", total)
	}
}

func Test_GetTotalPrice_OfferAndRemaining(t *testing.T) {
	co := newTestCheckout(t)
	skus := []string{"A", "A", "A", "A", "A"}
	for _, sku := range skus {
		err := co.Scan(sku)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	}

	total, err := co.GetTotalPrice()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := 130 + (2 * 50)
	if total != expected {
		t.Errorf("expected total %d, got %d", expected, total)
	}
}

func Test_GetTotalPrice_MixedItems(t *testing.T) {
	co := newTestCheckout(t)

	skus := []string{"B", "B", "A", "A", "A", "C", "D"}
	for _, sku := range skus {
		_ = co.Scan(sku)
	}
	expected := 45 + 130 + 20 + 15

	total, err := co.GetTotalPrice()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if total != expected {
		t.Errorf("expected total %d, got %d", expected, total)
	}
}

func Test_GetTotalPrice_InvalidSKUError(t *testing.T) {
	co := newTestCheckout(t)
	co.scannedItems["Z"] = 2

	_, err := co.GetTotalPrice()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	var skuError InvalidSKUError
	if !errors.As(err, &skuError) {
		t.Errorf("expected InvalidSKUError, got %v", err)
	}
}
