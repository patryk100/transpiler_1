package models

import (
	"fmt"
	"strings"
)

// S3Bucket represents an AWS S3 bucket resource
type S3Bucket struct {
	ResourceName      string            `json:"resource_name"`
	Bucket            string            `json:"bucket"`
	ACL               string            `json:"acl,omitempty"`
	VersioningEnabled bool              `json:"versioning_enabled,omitempty"`
	Tags              map[string]string `json:"tags,omitempty"`
}

// GetResourceType returns the Terraform resource type
func (s *S3Bucket) GetResourceType() string {
	return "aws_s3_bucket"
}

// GetResourceName returns the resource identifier name
func (s *S3Bucket) GetResourceName() string {
	return s.ResourceName
}

// ToHCL generates Terraform HCL code for the S3 bucket
func (s *S3Bucket) ToHCL() string {
	var hcl strings.Builder

	// Resource declaration
	hcl.WriteString(fmt.Sprintf("resource \"%s\" \"%s\" {\n", s.GetResourceType(), s.ResourceName))

	// Required: bucket name
	hcl.WriteString(fmt.Sprintf("  bucket = \"%s\"\n", s.Bucket))

	// Optional: ACL
	if s.ACL != "" {
		hcl.WriteString(fmt.Sprintf("  acl    = \"%s\"\n", s.ACL))
	}

	// Optional: Versioning (as a nested block)
	if s.VersioningEnabled {
		hcl.WriteString("\n  versioning {\n")
		hcl.WriteString("    enabled = true\n")
		hcl.WriteString("  }\n")
	}

	// Optional: Tags
	if len(s.Tags) > 0 {
		hcl.WriteString("\n  tags = {\n")
		for key, value := range s.Tags {
			hcl.WriteString(fmt.Sprintf("    %s = \"%s\"\n", key, value))
		}
		hcl.WriteString("  }\n")
	}

	hcl.WriteString("}\n")

	return hcl.String()
}
