// Package checkout handles scanning items and calculating total price.
package checkout

var pricingRules = map[string]struct {
	unitPrice     int
	offerQuantity int
	offerPrice    int
}{
	"A": {unitPrice: 50, offerQuantity: 3, offerPrice: 130},
	"B": {unitPrice: 30, offerQuantity: 2, offerPrice: 45},
	"C": {unitPrice: 20},
	"D": {unitPrice: 15},
}

// Checkout is a struct that holds scanned items.
type Checkout struct {
	scannedItems map[string]int
}

// NewCheckout creates a new checkout instance with an empty item list.
func NewCheckout() *Checkout {
	return &Checkout{
		scannedItems: make(map[string]int),
	}
}

// Scan adds an item to the scanned list. Returns error if SKU is empty or invalid.
func (c *Checkout) Scan(sku string) error {
	if sku == "" {
		return ErrEmptySKU
	}

	if _, ok := pricingRules[sku]; !ok {
		return InvalidSKUError{SKU: sku}
	}

	c.scannedItems[sku]++
	return nil
}

// GetTotalPrice calculates the final amount including offers
func (c *Checkout) GetTotalPrice() (int, error) {
	total := 0

	for sku, qty := range c.scannedItems {
		rule, ok := pricingRules[sku]
		if !ok {
			return 0, InvalidSKUError{SKU: sku}
		}

		if rule.offerQuantity > 0 && rule.offerPrice > 0 {
			offerCount := qty / rule.offerQuantity
			remaining := qty % rule.offerQuantity
			total += offerCount*rule.offerPrice + remaining*rule.unitPrice
			continue
		}

		total += qty * rule.unitPrice
	}

	return total, nil
}
