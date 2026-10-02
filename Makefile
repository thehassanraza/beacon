.PHONY: dev build install clean test test-verbose

test:
	@go test ./...

test-verbose:
	@go test ./... -v

dev:
	@echo "Starting development server..."
	@go run . & pnpm --dir frontend dev

build:
	@echo "Building frontend..."
	@pnpm --dir frontend build
	@echo "Building backend..."
	@go build -o beacon .

install: build
	@echo "Installing beacon to /usr/local/bin (may require sudo)..."
	@install -m 755 beacon /usr/local/bin/beacon
	@echo "Installed: beacon (run 'beacon help')"

docker-build:
	@echo "Building Docker image..."
	@docker build -t beacon:latest .

up:
	@echo "Starting Beacon and test containers..."
	@touch app.db
	@docker-compose up --build -d

down:
	@echo "Stopping containers..."
	@docker-compose down

clean:
	@rm -rf beacon frontend/dist
	@echo "Cleaned up."
