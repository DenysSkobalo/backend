.PHONY: run build up down clean

# Run the Go server
run:
	@echo "Running the Go server..."
	@go run ./cmd

# Build the Go project
build:
	@echo "Building the Go project..."
	@go build -o ./bin/main ./cmd

# Start the development environment with Docker Compose
up:
	@echo "Starting Docker containers..."
	@docker-compose -f ./deployments/docker-compose.yml up -d

# Stop and remove the Docker Compose development environment
down:
	@echo "Stopping and cleaning Docker containers..."
	@docker-compose -f ./deployments/docker-compose.yml down

# Clean up any binaries or other artifacts created by the build
clean:
	@echo "Cleaning up..."
	@rm -rf ./bin
