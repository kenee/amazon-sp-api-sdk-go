package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kenee/amazon-sp-api-sdk-go/apis/orders"
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
	ordersAPI := orders.NewOrdersAPI(config)

	// Create context
	ctx := context.Background()

	// Call getOrders
	response, err := ordersAPI.GetOrdersSimple(ctx, []string{"ATVPDKIKX0DER"}, "2023-01-01T00:00:00Z")
	if err != nil {
		log.Fatalf("Exception when calling OrdersAPI->GetOrders: %v", err)
	}

	// Process Orders API response
	fmt.Println("Order API Response:")
	fmt.Printf("%+v\n", response)

	// Display order information
	if response.Payload != nil && len(response.Payload.Orders) > 0 {
		fmt.Printf("Number of orders: %d\n", len(response.Payload.Orders))
		for i, order := range response.Payload.Orders {
			fmt.Printf("Order %d: %s\n", i+1, order.AmazonOrderId)
		}
	} else {
		fmt.Println("No orders found")
	}
}

// maskString masks sensitive information for display
func maskString(s string) string {
	if len(s) <= 8 {
		return "***"
	}
	return s[:4] + "***" + s[len(s)-4:]
}
