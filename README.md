# checkout-system

A small Go app that scans SKUs like A, B, C, D and gives the total. It also handles offers "3 for 130" or "2 for 45".

## How it works

- Each SKU has a price
- Some SKUs have offers (e.g A: 3 for 130)
- Scan items in any order
- It adds them up with offers applied properly

## Run 
`go run ./cmd/main.go`

## Test
`go test ./...`

## Requirements

- Go v1.24.0+
