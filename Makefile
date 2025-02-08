# Makefile for Meisterwerk

# Variables
IMAGE_NAME = meisterwerk
PORT = 8080

# Build the Docker image
build:
	docker build -t $(IMAGE_NAME) .

# Run the Docker container
run:
	docker run -d -p $(PORT):$(PORT) $(IMAGE_NAME)

# Stop the Docker container (if running)
stop:
	docker ps -q --filter "ancestor=$(IMAGE_NAME)" | xargs -r docker stop

# Remove the Docker container (if exists)
clean:
	docker rm -f $$(docker ps -aq --filter "ancestor=$(IMAGE_NAME)")

# Access the application
access:
	curl http://localhost:$(PORT)/health

# Remove unused Docker data
prune:
	docker system prune -a

# Start all services with docker-compose
compose-up:
	docker-compose up -d

# Stop all services
compose-down:
	docker-compose down

# View logs
compose-logs:
	docker-compose logs -f

# Rebuild and restart services
compose-rebuild:
	docker-compose up -d --build

# Clean up volumes
compose-clean:
	docker-compose down -v

# Help command
help:
	@echo "Makefile commands:"
	@echo "  make build          - Build the Docker image"
	@echo "  make run           - Run the Docker container"
	@echo "  make stop          - Stop the running Docker container"
	@echo "  make clean         - Remove the Docker container"
	@echo "  make access        - Access the health endpoint"
	@echo "  make prune         - Remove unused Docker data"
	@echo "  make compose-up    - Start all services with docker-compose"
	@echo "  make compose-down  - Stop all services"
	@echo "  make compose-logs  - View logs"
	@echo "  make compose-rebuild - Rebuild and restart services"
	@echo "  make compose-clean - Clean up volumes"
	@echo "  make help          - Show this help message"