# Makefile for GoRedis

.PHONY: build run test clean deps

# Build the application
build:
	go build -o bin/goredis-server cmd/server/main.go

# Run the application
run:
	go run cmd/server/main.go

# Install dependencies
deps:
	go mod tidy
	go mod download

# Test the application
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Run with hot reload (requires air)
dev:
	air

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Docker build
docker-build:
	docker build -t goredis:latest .

# Docker run
docker-run:
	docker run -p 8080:8080 goredis:latest

# Help
help:
	@echo "Available commands:"
	@echo "  build       - Build the application"
	@echo "  run         - Run the application"
	@echo "  test        - Run tests"
	@echo "  clean       - Clean build artifacts"
	@echo "  deps        - Install dependencies"
	@echo "  dev         - Run with hot reload"
	@echo "  fmt         - Format code"
	@echo "  lint        - Lint code"
	@echo "  docker-build- Build Docker image"
	@echo "  docker-run  - Run Docker container"
