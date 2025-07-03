package producttypes

import (
	"context"
	"fmt"

	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

// ProductTypesAPI represents the Product Types API client
type ProductTypesAPI struct {
	apiClient *client.APIClient
}

// NewProductTypesAPI creates a new Product Types API client
func NewProductTypesAPI(config *client.Configuration) *ProductTypesAPI {
	return &ProductTypesAPI{
		apiClient: client.NewAPIClient(config),
	}
}

// GetDefinitionsProductType retrieves a product type definition
func (p *ProductTypesAPI) GetDefinitionsProductType(ctx context.Context, request *GetDefinitionsProductTypeRequest) (*GetDefinitionsProductTypeResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Required parameters
	params["marketplaceIds"] = request.MarketplaceIds

	// Optional parameters
	if request.SellerId != "" {
		params["sellerId"] = request.SellerId
	}
	if request.ProductTypeVersion != "" {
		params["productTypeVersion"] = request.ProductTypeVersion
	}
	if request.Requirements != "" {
		params["requirements"] = request.Requirements
	}
	if request.RequirementsEnforced != "" {
		params["requirementsEnforced"] = request.RequirementsEnforced
	}
	if request.Locale != "" {
		params["locale"] = request.Locale
	}

	// Build query string
	queryString := p.apiClient.BuildQueryString(params)

	// Build the full path
	path := fmt.Sprintf("/definitions/2020-09-01/productTypes/%s", request.ProductType)
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := p.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getDefinitionsProductType API: %w", err)
	}

	// Parse the response
	var result GetDefinitionsProductTypeResponse
	if err := p.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getDefinitionsProductType response: %w", err)
	}

	return &result, nil
}

// GetDefinitionsProductTypeSimple is a simplified version of GetDefinitionsProductType for common use cases
func (p *ProductTypesAPI) GetDefinitionsProductTypeSimple(ctx context.Context, productType string, marketplaceIds []string) (*GetDefinitionsProductTypeResponse, error) {
	request := &GetDefinitionsProductTypeRequest{
		ProductType:          productType,
		MarketplaceIds:       marketplaceIds,
		ProductTypeVersion:   "LATEST",
		Requirements:         "LISTING",
		RequirementsEnforced: "ENFORCED",
		Locale:               "en_US",
	}
	return p.GetDefinitionsProductType(ctx, request)
}

// SearchDefinitionsProductTypes searches for product type definitions
func (p *ProductTypesAPI) SearchDefinitionsProductTypes(ctx context.Context, request *SearchDefinitionsProductTypesRequest) (*SearchDefinitionsProductTypesResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Required parameters
	params["marketplaceIds"] = request.MarketplaceIds

	// Optional parameters
	if len(request.Keywords) > 0 {
		params["keywords"] = request.Keywords
	}
	if request.ItemName != "" {
		params["itemName"] = request.ItemName
	}
	if request.Locale != "" {
		params["locale"] = request.Locale
	}
	if request.SearchLocale != "" {
		params["searchLocale"] = request.SearchLocale
	}

	// Build query string
	queryString := p.apiClient.BuildQueryString(params)

	// Build the full path
	path := "/definitions/2020-09-01/productTypes"
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := p.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call searchDefinitionsProductTypes API: %w", err)
	}

	// Parse the response
	var result SearchDefinitionsProductTypesResponse
	if err := p.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse searchDefinitionsProductTypes response: %w", err)
	}

	return &result, nil
}

// SearchDefinitionsProductTypesSimple is a simplified version of SearchDefinitionsProductTypes for common use cases
func (p *ProductTypesAPI) SearchDefinitionsProductTypesSimple(ctx context.Context, marketplaceIds []string, keywords []string) (*SearchDefinitionsProductTypesResponse, error) {
	request := &SearchDefinitionsProductTypesRequest{
		MarketplaceIds: marketplaceIds,
		Keywords:       keywords,
		Locale:         "en_US",
	}
	return p.SearchDefinitionsProductTypes(ctx, request)
}
