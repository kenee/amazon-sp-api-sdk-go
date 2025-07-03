package reports

import (
	"context"
	"fmt"

	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

// ReportsAPI represents the Reports API client
type ReportsAPI struct {
	apiClient *client.APIClient
}

// NewReportsAPI creates a new Reports API client
func NewReportsAPI(config *client.Configuration) *ReportsAPI {
	return &ReportsAPI{
		apiClient: client.NewAPIClient(config),
	}
}

// CreateReport creates a new report
func (r *ReportsAPI) CreateReport(ctx context.Context, request *CreateReportRequest) (*CreateReportResponse, error) {
	// Build the full path
	path := "/reports/2021-06-30/reports"

	// Make the API call
	resp, err := r.apiClient.CallAPI(ctx, "POST", path, request, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call createReport API: %w", err)
	}

	// Parse the response
	var result CreateReportResponse
	if err := r.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse createReport response: %w", err)
	}

	return &result, nil
}

// GetReport retrieves report information
func (r *ReportsAPI) GetReport(ctx context.Context, reportId string) (*GetReportResponse, error) {
	// Build the full path
	path := fmt.Sprintf("/reports/2021-06-30/reports/%s", reportId)

	// Make the API call
	resp, err := r.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getReport API: %w", err)
	}

	// Parse the response
	var result GetReportResponse
	if err := r.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getReport response: %w", err)
	}

	return &result, nil
}

// GetReports retrieves a list of reports
func (r *ReportsAPI) GetReports(ctx context.Context, request *GetReportsRequest) (*GetReportsResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Optional parameters
	if len(request.ReportTypes) > 0 {
		params["reportTypes"] = request.ReportTypes
	}
	if len(request.MarketplaceIds) > 0 {
		params["marketplaceIds"] = request.MarketplaceIds
	}
	if len(request.ProcessingStatuses) > 0 {
		params["processingStatuses"] = request.ProcessingStatuses
	}
	if request.DataStartTime != "" {
		params["dataStartTime"] = request.DataStartTime
	}
	if request.DataEndTime != "" {
		params["dataEndTime"] = request.DataEndTime
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
	queryString := r.apiClient.BuildQueryString(params)

	// Build the full path
	path := "/reports/2021-06-30/reports"
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := r.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getReports API: %w", err)
	}

	// Parse the response
	var result GetReportsResponse
	if err := r.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getReports response: %w", err)
	}

	return &result, nil
}

// CancelReport cancels a report
func (r *ReportsAPI) CancelReport(ctx context.Context, reportId string) (*CancelReportResponse, error) {
	// Build the full path
	path := fmt.Sprintf("/reports/2021-06-30/reports/%s/cancel", reportId)

	// Make the API call
	resp, err := r.apiClient.CallAPI(ctx, "DELETE", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call cancelReport API: %w", err)
	}

	// Parse the response
	var result CancelReportResponse
	if err := r.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse cancelReport response: %w", err)
	}

	return &result, nil
}

// GetReportDocument retrieves report document information
func (r *ReportsAPI) GetReportDocument(ctx context.Context, reportDocumentId string) (*GetReportDocumentResponse, error) {
	// Build the full path
	path := fmt.Sprintf("/reports/2021-06-30/documents/%s", reportDocumentId)

	// Make the API call
	resp, err := r.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getReportDocument API: %w", err)
	}

	// Parse the response
	var result GetReportDocumentResponse
	if err := r.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getReportDocument response: %w", err)
	}

	return &result, nil
}
