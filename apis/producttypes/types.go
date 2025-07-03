package producttypes

import (
	"encoding/json"
)

// GetDefinitionsProductTypeRequest represents the request parameters for getDefinitionsProductType
type GetDefinitionsProductTypeRequest struct {
	ProductType          string   `json:"productType"`
	MarketplaceIds       []string `json:"marketplaceIds"`
	SellerId             string   `json:"sellerId,omitempty"`
	ProductTypeVersion   string   `json:"productTypeVersion,omitempty"`
	Requirements         string   `json:"requirements,omitempty"`
	RequirementsEnforced string   `json:"requirementsEnforced,omitempty"`
	Locale               string   `json:"locale,omitempty"`
}

// GetDefinitionsProductTypeResponse represents the response from getDefinitionsProductType
type GetDefinitionsProductTypeResponse struct {
	MetaSchema           *SchemaLink               `json:"metaSchema,omitempty"`
	Schema               *SchemaLink               `json:"schema"`
	Requirements         string                    `json:"requirements"`
	RequirementsEnforced string                    `json:"requirementsEnforced"`
	PropertyGroups       map[string]*PropertyGroup `json:"propertyGroups,omitempty"`
	Locale               string                    `json:"locale"`
	MarketplaceIds       []string                  `json:"marketplaceIds"`
	ProductType          string                    `json:"productType"`
	DisplayName          string                    `json:"displayName"`
	ProductTypeVersion   *ProductTypeVersion       `json:"productTypeVersion"`
}

// SearchDefinitionsProductTypesRequest represents the request parameters for searchDefinitionsProductTypes
type SearchDefinitionsProductTypesRequest struct {
	MarketplaceIds []string `json:"marketplaceIds"`
	Keywords       []string `json:"keywords,omitempty"`
	ItemName       string   `json:"itemName,omitempty"`
	Locale         string   `json:"locale,omitempty"`
	SearchLocale   string   `json:"searchLocale,omitempty"`
}

// SearchDefinitionsProductTypesResponse represents the response from searchDefinitionsProductTypes
type SearchDefinitionsProductTypesResponse struct {
	ProductTypes       []*ProductType `json:"productTypes"`
	ProductTypeVersion string         `json:"productTypeVersion"`
}

// ProductType represents an Amazon product type with a definition available
type ProductType struct {
	Name           string   `json:"name"`
	DisplayName    string   `json:"displayName"`
	MarketplaceIds []string `json:"marketplaceIds"`
}

// SchemaLink represents a schema link
type SchemaLink struct {
	Link          *SchemaLinkLink `json:"link"`
	ContentType   string          `json:"contentType,omitempty"`
	SchemaVersion string          `json:"schemaVersion,omitempty"`
	Name          string          `json:"name,omitempty"`
	Description   string          `json:"description,omitempty"`
	Checksum      string          `json:"checksum,omitempty"`
	Verify        bool            `json:"verify,omitempty"`
}

// SchemaLinkLink represents a link to retrieve the schema
type SchemaLinkLink struct {
	Resource string `json:"resource"`
	Verb     string `json:"verb"`
}

// UnmarshalJSON implements custom unmarshalling for SchemaLink to handle null, object, and string
func (s *SchemaLink) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		// If the value is null, just return nil (leave struct zeroed)
		return nil
	}
	// Try to unmarshal as object first
	type Alias SchemaLink
	var alias Alias
	if err := json.Unmarshal(data, &alias); err == nil {
		*s = SchemaLink(alias)
		return nil
	}
	// If not object, try as string (for backward compatibility)
	var link string
	if err := json.Unmarshal(data, &link); err == nil {
		s.Link = &SchemaLinkLink{
			Resource: link,
			Verb:     "GET",
		}
		return nil
	}
	// Otherwise, return error
	return json.Unmarshal(data, &alias)
}

// PropertyGroup represents a property group
type PropertyGroup struct {
	Title         string   `json:"title,omitempty"`
	PropertyNames []string `json:"propertyNames,omitempty"`
}

// ProductTypeVersion represents a product type version
type ProductTypeVersion struct {
	Version          string `json:"version"`
	Latest           bool   `json:"latest"`
	ReleaseCandidate bool   `json:"releaseCandidate,omitempty"`
}
