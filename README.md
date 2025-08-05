# GoRedis - A Lightweight Redis-Compatible Server in Go

GoRedis is a minimal Redis-compatible TCP server built using Go. It supports basic `GET` and `SET` operations, allowing interaction with clients using the Redis protocol. It is primarily designed for learning and experimentation purposes.

---

## ✨ Features

- Redis-compatible protocol parsing using [tidwall/resp](https://github.com/tidwall/resp)
- Supports basic commands: `HELLO`, `CLIENT`, `SET`, and `GET`
- In-memory key-value store with mutex protection
- Concurrency-safe via goroutines and channels
- Integration tested with the official `go-redis` client

---

## 📁 Project Structure

```
.
├── main.go            # Entry point
├── server.go          # TCP server and message loop
├── peer.go            # Peer connection lifecycle
├── proto.go           # RESP protocol handling and command parsing
├── keyval.go          # In-memory key-value store
├── server_test.go     # Helper types for testing
├── redis_test.go      # Integration test using go-redis client
├── Makefile           # Build and run helper commands
```

---

## 🚀 Getting Started

### Prerequisites

- Go 1.18+
- Git

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/goredis.git
   cd goredis
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Build the project:

   ```bash
   make build
   ```

---

## ▶️ Running the Server

To start the GoRedis server:

```bash
make run
```

The server will listen on port `:5001` by default.

---

## 🧪 Testing

Run tests using:

```bash
make test
```

The tests include:

- Starting the server on `:5001`
- Using the official `go-redis` client to connect
- Verifying `SET` and `GET` functionality

---

## 🛠️ Makefile Commands

This project includes a `Makefile` to simplify common tasks:

```makefile
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

# Run linter (optional)
.PHONY: lint
lint:
	@golangci-lint run
```

---

## 💬 Supported Commands

| Command         | Description                      |
|------------------|----------------------------------|
| `HELLO`         | RESP handshake (returns map)     |
| `CLIENT`        | Dummy handler, responds with OK  |
| `SET key val`   | Stores a value for a key         |
| `GET key`       | Retrieves value for a key        |

---

## 🛡 License

MIT License. See [LICENSE](LICENSE) for more details.

---

## 🙏 Acknowledgements

- [tidwall/resp](https://github.com/tidwall/resp) – RESP protocol parser
- [go-redis/redis](https://github.com/redis/go-redis) – Official Go Redis client used for testing
