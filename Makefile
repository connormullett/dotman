.PHONY: build clean install test run help

# Binary name
BINARY_NAME=dotman

# Build directory
BUILD_DIR=bin

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod

# Build the project
build:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v

install:

# Install dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# Run tests
test:
	$(GOTEST) -v ./...

# Clean build files
clean:
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)

# Install binary to GOPATH/bin
install:
	$(GOCMD) install -v
