package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kenee/amazon-sp-api-sdk-go/apis/catalog"
	"github.com/kenee/amazon-sp-api-sdk-go/auth"
	"github.com/kenee/amazon-sp-api-sdk-go/client"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file in parent directory
	envPath := filepath.Join("..", "..", "..", ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Set up LWA credentials
	credentials := &auth.LWAAuthorizationCredentials{
		ClientID:     os.Getenv("SP_API_CLIENT_ID"),
		ClientSecret: os.Getenv("SP_API_CLIENT_SECRET"),
		RefreshToken: os.Getenv("SP_API_REFRESH_TOKEN"),
		Endpoint:     os.Getenv("SP_API_ENDPOINT"),
	}

	// Initialize config and set SP-API endpoint region
	config := client.NewConfigurationWithCredentials(credentials)
	config.SetHost(os.Getenv("SP_API_ENDPOINT_HOST"))

	// Create API instance
	catalogAPI := catalog.NewCatalogAPI(config)

	// Create context
	ctx := context.Background()

	// Call getCatalogItem
	response, err := catalogAPI.GetCatalogItemSimple(ctx, "B071VG5N9D", []string{"ATVPDKIKX0DER"})
	if err != nil {
		log.Fatalf("Exception when calling CatalogAPI->GetCatalogItem: %v", err)
	}

	// Process Catalog API response
	fmt.Println("Catalog Item API Response:")
	fmt.Printf("%+v\n", response)
}

// maskString masks sensitive information for display
func maskString(s string) string {
	if len(s) <= 8 {
		return "***"
	}
	return s[:4] + "***" + s[len(s)-4:]
}
