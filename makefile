.PHONY: default run build test docs clean

# variables
APP_NAME = "gestock-api"

# tasks
default: run-with-docs

run:
	@echo "Running application..."
	@go run cmd/server/main.go
run-with-docs:
	@echo "Running application with docs..."
	@swag init
	@go run cmd/server/main.go
build:
	@echo "Building binary..."
	@go build -o $(APP_NAME) cmd/server/main.go
test:
	@echo "Running tests..."
	@go test -v ./...
docs:
	@echo "Generating docs..."
	@swag init
clean:
	@echo "Cleaning up..."
	@rm -rf $(APP_NAME) ./docs
tidy:
	@echo "Tidying up..."
	@go mod tidy