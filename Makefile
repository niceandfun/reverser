# Project configuration
PROJECT_NAME := reverser

# Directories
PROTO_DIR := proto
BUILD_DIR := build
CMD_DIR := cmd
PKG_DIR := pkg

# Go configuration
GO := go
GOPATH := $(shell go env GOPATH)
GOBIN := $(GOPATH)/bin

# Protoc configuration
PROTOC := protoc
PROTO_GO_PLUGIN := protoc-gen-go
PROTO_GRPC_PLUGIN := protoc-gen-go-grpc

# Find all proto files
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)
PROTO_GO_FILES := $(PROTO_FILES:$(PROTO_DIR)/%.proto=$(BUILD_DIR)/%.pb.go)
GRPC_GO_FILES := $(PROTO_FILES:$(PROTO_DIR)/%.proto=$(BUILD_DIR)/%.grpc.pb.go)

.PHONY: all clean deps proto build test run

all: deps proto build

# Install dependencies
deps:
	@echo "Installing dependencies..."
	$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	$(GO) mod tidy

# Create build directory
$(BUILD_DIR):
	@mkdir -p $(BUILD_DIR)

# Compile proto files
proto: $(BUILD_DIR)
	@echo "Generating protobuf code..."
	@for file in $(PROTO_FILES); do \
		$(PROTOC) -I$(PROTO_DIR) \
			--go_out=$(BUILD_DIR) --go_opt=paths=source_relative \
			--go-grpc_out=$(BUILD_DIR) --go-grpc_opt=paths=source_relative \
			$$file; \
	done

# Build the project
build: proto
	@echo "Building project..."
	$(GO) build -o $(BUILD_DIR)/$(PROJECT_NAME) ./cmd/...

# Run the application
run: build
	./$(BUILD_DIR)/$(PROJECT_NAME)

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR)
