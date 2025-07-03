package feeds

import "time"

// CreateFeedRequest represents the request for creating a feed
type CreateFeedRequest struct {
	FeedType            string       `json:"feedType"`
	MarketplaceIds      []string     `json:"marketplaceIds"`
	InputFeedDocumentId string       `json:"inputFeedDocumentId"`
	FeedOptions         *FeedOptions `json:"feedOptions,omitempty"`
}

// FeedOptions represents feed options
type FeedOptions struct {
	// Add specific feed options as needed
}

// CreateFeedResponse represents the response from creating a feed
type CreateFeedResponse struct {
	FeedId string `json:"feedId"`
}

// GetFeedResponse represents the response from getting a feed
type GetFeedResponse struct {
	FeedId               string     `json:"feedId"`
	FeedType             string     `json:"feedType"`
	MarketplaceIds       []string   `json:"marketplaceIds"`
	CreatedTime          time.Time  `json:"createdTime"`
	ProcessingStatus     string     `json:"processingStatus"`
	ProcessingStartTime  *time.Time `json:"processingStartTime,omitempty"`
	ProcessingEndTime    *time.Time `json:"processingEndTime,omitempty"`
	ResultFeedDocumentId *string    `json:"resultFeedDocumentId,omitempty"`
	FeedDocumentId       string     `json:"feedDocumentId"`
}

// GetFeedsRequest represents the request for getting feeds
type GetFeedsRequest struct {
	FeedTypes          []string `json:"feedTypes,omitempty"`
	MarketplaceIds     []string `json:"marketplaceIds,omitempty"`
	ProcessingStatuses []string `json:"processingStatuses,omitempty"`
	CreatedSince       string   `json:"createdSince,omitempty"`
	CreatedUntil       string   `json:"createdUntil,omitempty"`
	NextToken          string   `json:"nextToken,omitempty"`
	PageSize           int      `json:"pageSize,omitempty"`
}

// GetFeedsResponse represents the response from getting feeds
type GetFeedsResponse struct {
	Feeds     []GetFeedResponse `json:"feeds"`
	NextToken string            `json:"nextToken,omitempty"`
}

// CancelFeedResponse represents the response from canceling a feed
type CancelFeedResponse struct {
	FeedId string `json:"feedId"`
}

// CreateFeedDocumentRequest represents the request for creating a feed document
type CreateFeedDocumentRequest struct {
	ContentType string `json:"contentType"`
}

// CreateFeedDocumentResponse represents the response from creating a feed document
type CreateFeedDocumentResponse struct {
	FeedDocumentId string `json:"feedDocumentId"`
	Url            string `json:"url"`
}

// GetFeedDocumentResponse represents the response from getting a feed document
type GetFeedDocumentResponse struct {
	FeedDocumentId       string `json:"feedDocumentId"`
	Url                  string `json:"url"`
	CompressionAlgorithm string `json:"compressionAlgorithm,omitempty"`
}
