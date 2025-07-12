package main

import (
	"log"

	"checkout-system/internal/app"
)

func main() {
	a := app.NewApplication()

	log.Printf("application running...")
	if err := a.Run(); err != nil {
		log.Fatalf("failed to run application: %v", err)
	}
}
