package checkout

type Checkout struct {
}

func NewCheckout() *Checkout {
	return &Checkout{}
}

func (c *Checkout) Scan(sku string) error {
	return nil
}

func (c *Checkout) GetTotalPrice() (int, error) {
	return 0, nil
}
