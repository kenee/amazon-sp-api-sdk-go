package feeds

import (
	"context"
	"fmt"

	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

// FeedsAPI represents the Feeds API client
type FeedsAPI struct {
	apiClient *client.APIClient
}

// NewFeedsAPI creates a new Feeds API client
func NewFeedsAPI(config *client.Configuration) *FeedsAPI {
	return &FeedsAPI{
		apiClient: client.NewAPIClient(config),
	}
}

// CreateFeed creates a new feed
func (f *FeedsAPI) CreateFeed(ctx context.Context, request *CreateFeedRequest) (*CreateFeedResponse, error) {
	// Build the full path
	path := "/feeds/2021-06-30/feeds"

	// Make the API call
	resp, err := f.apiClient.CallAPI(ctx, "POST", path, request, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call createFeed API: %w", err)
	}

	// Parse the response
	var result CreateFeedResponse
	if err := f.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse createFeed response: %w", err)
	}

	return &result, nil
}

// GetFeed retrieves feed information
func (f *FeedsAPI) GetFeed(ctx context.Context, feedId string) (*GetFeedResponse, error) {
	// Build the full path
	path := fmt.Sprintf("/feeds/2021-06-30/feeds/%s", feedId)

	// Make the API call
	resp, err := f.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getFeed API: %w", err)
	}

	// Parse the response
	var result GetFeedResponse
	if err := f.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getFeed response: %w", err)
	}

	return &result, nil
}

// GetFeeds retrieves a list of feeds
func (f *FeedsAPI) GetFeeds(ctx context.Context, request *GetFeedsRequest) (*GetFeedsResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Optional parameters
	if len(request.FeedTypes) > 0 {
		params["feedTypes"] = request.FeedTypes
	}
	if len(request.MarketplaceIds) > 0 {
		params["marketplaceIds"] = request.MarketplaceIds
	}
	if len(request.ProcessingStatuses) > 0 {
		params["processingStatuses"] = request.ProcessingStatuses
	}
	if request.CreatedSince != "" {
		params["createdSince"] = request.CreatedSince
	}
	if request.CreatedUntil != "" {
		params["createdUntil"] = request.CreatedUntil
	}
	if request.NextToken != "" {
		params["nextToken"] = request.NextToken
	}
	if request.PageSize > 0 {
		params["pageSize"] = request.PageSize
	}

	// Build query string
	queryString := f.apiClient.BuildQueryString(params)

	// Build the full path
	path := "/feeds/2021-06-30/feeds"
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := f.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getFeeds API: %w", err)
	}

	// Parse the response
	var result GetFeedsResponse
	if err := f.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getFeeds response: %w", err)
	}

	return &result, nil
}

// CancelFeed cancels a feed
func (f *FeedsAPI) CancelFeed(ctx context.Context, feedId string) (*CancelFeedResponse, error) {
	// Build the full path
	path := fmt.Sprintf("/feeds/2021-06-30/feeds/%s/cancel", feedId)

	// Make the API call
	resp, err := f.apiClient.CallAPI(ctx, "DELETE", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call cancelFeed API: %w", err)
	}

	// Parse the response
	var result CancelFeedResponse
	if err := f.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse cancelFeed response: %w", err)
	}

	return &result, nil
}

// CreateFeedDocument creates a feed document
func (f *FeedsAPI) CreateFeedDocument(ctx context.Context, request *CreateFeedDocumentRequest) (*CreateFeedDocumentResponse, error) {
	// Build the full path
	path := "/feeds/2021-06-30/documents"

	// Make the API call
	resp, err := f.apiClient.CallAPI(ctx, "POST", path, request, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call createFeedDocument API: %w", err)
	}

	// Parse the response
	var result CreateFeedDocumentResponse
	if err := f.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse createFeedDocument response: %w", err)
	}

	return &result, nil
}

// GetFeedDocument retrieves feed document information
func (f *FeedsAPI) GetFeedDocument(ctx context.Context, feedDocumentId string) (*GetFeedDocumentResponse, error) {
	// Build the full path
	path := fmt.Sprintf("/feeds/2021-06-30/documents/%s", feedDocumentId)

	// Make the API call
	resp, err := f.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getFeedDocument API: %w", err)
	}

	// Parse the response
	var result GetFeedDocumentResponse
	if err := f.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getFeedDocument response: %w", err)
	}

	return &result, nil
}
