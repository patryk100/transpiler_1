# Variables
BINARY_NAME=transpiler_1
ENTRY_POINT=./cmd/transpiler_1

# Default target (runs when you just type 'make')
all: build

# Build the application
build:
	@echo "Building..."
	go build -o build/$(BINARY_NAME) $(ENTRY_POINT)

# Run the application (development)
run:
	go run $(ENTRY_POINT)

# Run tests
test:
	go test ./...

# Clean up binary
clean:
	@echo "Cleaning..."
	go clean
	rm -f $(BINARY_NAME)

# Format code
fmt:
	go fmt ./...

.PHONY: all build run test clean fmt
