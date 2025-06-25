package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kenee/amazon-sp-api-sdk-go/apis/listings"
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
	listingsAPI := listings.NewListingsAPI(config)

	// Create context
	ctx := context.Background()

	// Call getListingsItem
	response, err := listingsAPI.GetListingsItemSimple(ctx, "A1B2C3D4E5F6G7", "GM-ZDPI-9B4E", []string{"ATVPDKIKX0DER"})
	if err != nil {
		log.Fatalf("Exception when calling ListingsAPI->GetListingsItem: %v", err)
	}

	// Process Listings Item API response
	fmt.Println("Listings Item API Response:")
	fmt.Printf("%+v\n", response)
}
