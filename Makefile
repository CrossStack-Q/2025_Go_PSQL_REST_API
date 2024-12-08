# Makefile

# Define variables
GO = go
MAIN_FILE = ./cmd/main.go

# Default target (run the app)
run:
	$(GO) run $(MAIN_FILE)

# Alternatively, you can define other common targets like build or clean:

build:
	$(GO) build -o bin/restapi $(MAIN_FILE)

clean:
	rm -rf ./bin/restapi

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@, $(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
