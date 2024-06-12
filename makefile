.PHONY: default run build test docs clean tidy create-migration migrate-up migrate-down sqlc

# variables
APP_NAME = "gestok-api"
name= "create_products_stocks_table"

# tasks
default: run-with-docs

run:
	@echo "Running application..."
	@go generate ./... && go run cmd/server/main.go
run-with-docs:
	@echo "Running application with docs..."
	@swag init -g cmd/server/main.go
	@go run cmd/server/main.go
build:
	@echo "Building binary..."
	@go build -o $(APP_NAME) cmd/server/main.go
test:
	@echo "Running tests..."
	@go test -v ./...
docs:
	@echo "Generating docs..."
	@swag init -g cmd/server/main.go
clean:
	@echo "Cleaning up..."
	@rm -rf $(APP_NAME) ./docs
tidy:
	@echo "Tidying up..."
	@go mod tidy
create-migration:
	@echo "Creating migration..."
	migrate create -ext=sql -dir=internal/infra/migrations -seq $(name)
migrate-up:
	@echo "Applying migrations..."
	migrate -database=$(DATABASE_URL) -path=internal/infra/migrations up
migrate-down:
	@echo "Rolling back migrations..."
	migrate -database=$(database) -path=sql/migrations down
sqlc:
	@echo "Generating SQLC..."
	sqlc generate