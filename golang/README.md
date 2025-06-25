# Amazon Selling Partner API SDK - Go

This is a Go implementation of the Amazon Selling Partner API SDK, forked from the official [Amazon SP-API SDK](https://github.com/amzn/selling-partner-api-sdk).

## 🚀 Features

- **Complete SP-API Support**: Orders, Catalog, and Listings APIs
- **LWA Authentication**: Login with Amazon authentication with token caching
- **Production Ready**: Clean, production-ready code without debug logs
- **Comprehensive Examples**: Working examples for all supported APIs
- **Error Handling**: Proper error handling and response parsing
- **Configuration Management**: Flexible configuration options

## 📦 Installation

```bash
go get github.com/kenee/amazon-sp-api-sdk-go
```

## 🔧 Prerequisites

Before using this SDK, you must:

1. Register as a Selling Partner API developer
2. Create an application in Seller Central
3. Generate LWA credentials (Client ID, Client Secret, Refresh Token)

For detailed setup instructions, see the [SP-API Registration Overview](https://developer-docs.amazon.com/sp-api/docs/registration-overview).

## 🛠️ Quick Start

### 1. Set up environment variables

Create a `.env` file in your project root:

```env
SP_API_CLIENT_ID=your_client_id
SP_API_CLIENT_SECRET=your_client_secret
SP_API_REFRESH_TOKEN=your_refresh_token
SP_API_ENDPOINT=https://api.amazon.com/auth/o2/token
SP_API_ENDPOINT_HOST=https://sandbox.sellingpartnerapi-na.amazon.com
```

### 2. Basic usage

```go
package main

import (
    "context"
    "log"
    "os"

    "github.com/kenee/amazon-sp-api-sdk-go/auth"
    "github.com/kenee/amazon-sp-api-sdk-go/client"
    "github.com/kenee/amazon-sp-api-sdk-go/apis/orders"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    godotenv.Load()

    // Set up credentials
    credentials := &auth.LWAAuthorizationCredentials{
        ClientID:     os.Getenv("SP_API_CLIENT_ID"),
        ClientSecret: os.Getenv("SP_API_CLIENT_SECRET"),
        RefreshToken: os.Getenv("SP_API_REFRESH_TOKEN"),
        Endpoint:     os.Getenv("SP_API_ENDPOINT"),
    }

    // Initialize configuration
    config := client.NewConfigurationWithCredentials(credentials)
    config.SetHost(os.Getenv("SP_API_ENDPOINT_HOST"))

    // Create API client
    ordersAPI := orders.NewOrdersAPI(config)

    // Make API call
    ctx := context.Background()
    response, err := ordersAPI.GetOrdersSimple(ctx, []string{"ATVPDKIKX0DER"}, "2023-01-01T00:00:00Z")
    if err != nil {
        log.Fatal(err)
    }

    // Process response
    if response.Payload != nil && len(response.Payload.Orders) > 0 {
        for _, order := range response.Payload.Orders {
            log.Printf("Order ID: %s", order.AmazonOrderId)
        }
    }
}
```

## 📚 API Examples

### Orders API

```go
// Get orders with filters
request := &orders.GetOrdersRequest{
    MarketplaceIds: []string{"ATVPDKIKX0DER"},
    CreatedAfter:   "2023-01-01T00:00:00Z",
    OrderStatuses:  []string{"Shipped", "Unshipped"},
}

response, err := ordersAPI.GetOrders(ctx, request)
```

### Catalog API

```go
// Get catalog item by ASIN
request := &catalog.GetCatalogItemRequest{
    ASIN:           "B071VG5N9D",
    MarketplaceIds: []string{"ATVPDKIKX0DER"},
    IncludedData:   []string{"summaries", "attributes", "images"},
}

response, err := catalogAPI.GetCatalogItem(ctx, request)
```

### Listings API

```go
// Get listings item
request := &listings.GetListingsItemRequest{
    SellerId:       "A1B2C3D4E5F6G7",
    SKU:            "GM-ZDPI-9B4E",
    MarketplaceIds: []string{"ATVPDKIKX0DER"},
    IncludedData:   []string{"summaries", "offers", "fulfillmentAvailability"},
}

response, err := listingsAPI.GetListingsItem(ctx, request)
```

## 🔐 Authentication

The SDK uses Login with Amazon (LWA) for authentication. Token caching is supported to improve performance:

```go
// Enable token caching
cache := auth.NewMemoryTokenCache()
config := client.NewConfigurationWithCredentialsAndCache(credentials, cache)
```

## 🌍 Environment Support

The SDK supports different SP-API environments:

- **Sandbox**: `https://sandbox.sellingpartnerapi-na.amazon.com`
- **Production NA**: `https://sellingpartnerapi-na.amazon.com`
- **Production EU**: `https://sellingpartnerapi-eu.amazon.com`
- **Production FE**: `https://sellingpartnerapi-fe.amazon.com`

## 📁 Project Structure

```
golang/
├── sdk/
│   ├── apis/           # API implementations
│   │   ├── orders/     # Orders API
│   │   ├── catalog/    # Catalog API
│   │   └── listings/   # Listings API
│   ├── auth/           # Authentication
│   ├── client/         # HTTP client
│   └── examples/       # Usage examples
└── README.md
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## 🔗 Links

- [Amazon Selling Partner API Documentation](https://developer-docs.amazon.com/sp-api/)
- [Official SP-API SDK](https://github.com/amzn/selling-partner-api-sdk)
- [SP-API Registration Overview](https://developer-docs.amazon.com/sp-api/docs/registration-overview)

## ⚠️ Disclaimer

This is an unofficial implementation and is not affiliated with Amazon. Use at your own risk. 