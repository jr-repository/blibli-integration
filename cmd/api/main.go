package main

import (
	"fmt"
	"log"

	"github.com/bytecorner/blibli-integration/internal/blibli"
	"github.com/bytecorner/blibli-integration/internal/config"
)

// MAIN EXECUTION AND TESTING
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Fatal error loading config: %v", err)
	}

	blibliClient := blibli.NewClient(cfg)

	fmt.Println("--- Starting Blibli API Integration Test ---")
	fmt.Println("Attempting to hit Product List endpoint...")

	products, err := blibliClient.GetProductList(0, 10)
	if err != nil {
		log.Fatalf("Error retrieving products: %v\nCheck your credentials and network connection.", err)
	}

	fmt.Printf("Successfully retrieved page %d of %d.\n", products.Number, products.TotalPages)
	fmt.Printf("Total Elements: %d\n\n", products.TotalElements)

	for i, product := range products.Content {
		fmt.Printf("[%d] %s (SKU: %s)\n", i+1, product.ProductName, product.ProductSku)
		fmt.Printf("    Price: %.2f | Stock: %d | Status: %s\n", product.Price, product.Stock, product.ProductStatus)
	}

	fmt.Println("--- Integration Test Completed Successfully ---")
}