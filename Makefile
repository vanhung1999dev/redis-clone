# Variables
BINARY_NAME = goredis
BUILD_DIR = bin
SRC = .

# Default target
.PHONY: all
all: build

# Build the binary
.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC)

# Run the server
.PHONY: run
run: build
	@echo "Running $(BINARY_NAME)..."
	@./$(BUILD_DIR)/$(BINARY_NAME) --listenAddr :5001

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Run linter (if you have golangci-lint installed)
.PHONY: lint
lint:
	@golangci-lint run
