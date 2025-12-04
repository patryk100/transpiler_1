package models

// ResourceConfig represents a generic resource from JSON input
type ResourceConfig struct {
	ResourceType string                 `json:"resource_type"`
	ResourceName string                 `json:"resource_name"`
	Properties   map[string]interface{} `json:"properties"`
}

// TerraformResource is the interface all resources must implement
type TerraformResource interface {
	GetResourceType() string
	GetResourceName() string
	ToHCL() string
}
