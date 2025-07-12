package app

import "checkout-system/internal/checkout"

type ICheckout interface {
	Scan(sku string) error
	GetTotalPrice() (int, error)
}

type Application struct {
	checkout ICheckout
}

func NewApplication() *Application {
	return &Application{
		checkout: checkout.NewCheckout(),
	}
}

func (a *Application) Run() {}
