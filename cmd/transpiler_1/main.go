package main

import (
	"fmt"
	"os"

	"github.com/patryk100/transpiler_1/internal/parser"
)

func main() {
	// Check if a file path was provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: transpiler_1 <json-file>")
		fmt.Println("Example: transpiler_1 examples/s3_bucket.json")
		os.Exit(1)
	}

	filePath := os.Args[1]

	// Create a new parser
	p := parser.NewParser()

	// Parse the JSON file
	resource, err := p.ParseFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing file: %v\n", err)
		os.Exit(1)
	}

	// Generate Terraform HCL
	hcl := resource.ToHCL()

	// Output the generated Terraform code
	fmt.Println("# Generated Terraform Configuration")
	fmt.Println()
	fmt.Print(hcl)
}
