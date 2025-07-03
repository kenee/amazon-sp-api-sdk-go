package reports

import "time"

// CreateReportRequest represents the request for creating a report
type CreateReportRequest struct {
	ReportType     string         `json:"reportType"`
	MarketplaceIds []string       `json:"marketplaceIds"`
	DataStartTime  string         `json:"dataStartTime,omitempty"`
	DataEndTime    string         `json:"dataEndTime,omitempty"`
	ReportOptions  *ReportOptions `json:"reportOptions,omitempty"`
}

// ReportOptions represents report options
type ReportOptions struct {
	// Add specific report options as needed
}

// CreateReportResponse represents the response from creating a report
type CreateReportResponse struct {
	ReportId string `json:"reportId"`
}

// GetReportResponse represents the response from getting a report
type GetReportResponse struct {
	ReportId            string     `json:"reportId"`
	ReportType          string     `json:"reportType"`
	MarketplaceIds      []string   `json:"marketplaceIds"`
	DataStartTime       time.Time  `json:"dataStartTime,omitempty"`
	DataEndTime         time.Time  `json:"dataEndTime,omitempty"`
	CreatedTime         time.Time  `json:"createdTime"`
	ProcessingStatus    string     `json:"processingStatus"`
	ProcessingStartTime *time.Time `json:"processingStartTime,omitempty"`
	ProcessingEndTime   *time.Time `json:"processingEndTime,omitempty"`
	ReportDocumentId    *string    `json:"reportDocumentId,omitempty"`
}

// GetReportsRequest represents the request for getting reports
type GetReportsRequest struct {
	ReportTypes        []string `json:"reportTypes,omitempty"`
	MarketplaceIds     []string `json:"marketplaceIds,omitempty"`
	ProcessingStatuses []string `json:"processingStatuses,omitempty"`
	DataStartTime      string   `json:"dataStartTime,omitempty"`
	DataEndTime        string   `json:"dataEndTime,omitempty"`
	CreatedSince       string   `json:"createdSince,omitempty"`
	CreatedUntil       string   `json:"createdUntil,omitempty"`
	NextToken          string   `json:"nextToken,omitempty"`
	PageSize           int      `json:"pageSize,omitempty"`
}

// GetReportsResponse represents the response from getting reports
type GetReportsResponse struct {
	Reports   []GetReportResponse `json:"reports"`
	NextToken string              `json:"nextToken,omitempty"`
}

// CancelReportResponse represents the response from canceling a report
type CancelReportResponse struct {
	ReportId string `json:"reportId"`
}

// GetReportDocumentResponse represents the response from getting a report document
type GetReportDocumentResponse struct {
	ReportDocumentId     string `json:"reportDocumentId"`
	Url                  string `json:"url"`
	CompressionAlgorithm string `json:"compressionAlgorithm,omitempty"`
}
