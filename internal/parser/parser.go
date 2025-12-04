package parser

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/patryk100/transpiler_1/internal/models"
)

// Parser handles JSON to Resource conversion
type Parser struct{}

// NewParser creates a new parser instance
func NewParser() *Parser {
	return &Parser{}
}

// ParseFile reads a JSON file and converts it to a TerraformResource
func (p *Parser) ParseFile(filepath string) (models.TerraformResource, error) {
	// Read the file
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return p.ParseJSON(data)
}

// ParseJSON converts JSON bytes into a TerraformResource
func (p *Parser) ParseJSON(data []byte) (models.TerraformResource, error) {
	// First, parse into generic ResourceConfig to determine type
	var config models.ResourceConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Based on resource_type, create the appropriate resource
	switch config.ResourceType {
	case "aws_s3_bucket":
		return p.parseS3Bucket(config)
	default:
		return nil, fmt.Errorf("unsupported resource type: %s", config.ResourceType)
	}
}

// parseS3Bucket converts ResourceConfig to S3Bucket
func (p *Parser) parseS3Bucket(config models.ResourceConfig) (*models.S3Bucket, error) {
	bucket := &models.S3Bucket{
		ResourceName: config.ResourceName,
	}

	// Extract bucket name (required)
	if bucketName, ok := config.Properties["bucket"].(string); ok {
		bucket.Bucket = bucketName
	} else {
		return nil, fmt.Errorf("bucket property is required and must be a string")
	}

	// Extract ACL (optional)
	if acl, ok := config.Properties["acl"].(string); ok {
		bucket.ACL = acl
	}

	// Extract versioning (optional)
	if versioning, ok := config.Properties["versioning_enabled"].(bool); ok {
		bucket.VersioningEnabled = versioning
	}

	// Extract tags (optional)
	if tags, ok := config.Properties["tags"].(map[string]interface{}); ok {
		bucket.Tags = make(map[string]string)
		for key, value := range tags {
			if strValue, ok := value.(string); ok {
				bucket.Tags[key] = strValue
			}
		}
	}

	return bucket, nil
}
