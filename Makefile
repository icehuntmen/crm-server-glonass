# Define project name
PROJECT_NAME := service

# Define directories
BIN_DIR := bin
SRC_DIRS := ./src/cmd ./src/internal ./src/pkg

# Define output binary
BINARY := $(BIN_DIR)/$(PROJECT_NAME)

# Define Go command
GO := go

.PHONY: all build

# Default target: build the project
all: build

# Build the project
build: $(BINARY)

# Build binary
$(BINARY): $(SRC_DIRS)
	mkdir -p $(BIN_DIR)
	$(GO) build -o $(BINARY) ./src/cmd/main

run:
	./bin/service

# Run tests
test:
	$(GO) test -v ./...


# Clean build artifacts
clean:
	rm -rf $(BIN_DIR)

docker :
	docker-compose -f docker/docker-compose.yaml up -d --build

