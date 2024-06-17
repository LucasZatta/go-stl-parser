# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	
	@go build -o main cmd/main/main.go

# Run the application
run:
	@go run cmd/main/main.go


# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main


.PHONY: all build run test clean
