# Meisterwerk

A Go-based REST API service for managing quotes and orders.

## Project Structure

```
.
├── cmd
│   └── main.go
├── internal
│   ├── handlers
│   │   ├── health_handler.go
│   │   └── quote_handler.go
│   ├── models
│   │   └── quote.go
│   └── service
│       └── quote_service.go
├── go.mod
└── README.md
```

## Running the Application with Docker

To run the application using Docker, follow these steps:

1. **Build the Docker image**:
   Navigate to the root of the project directory and run the following command to build the Docker image:

   ```bash
   docker build -t meisterwerk .
   ```

2. **Run the Docker container**:
   After the image is built, you can run the container with the following command:

   ```bash
   docker run -p 8080:8080 meisterwerk
   ```

   This command maps port 8080 of the container to port 8080 on your host machine.

3. **Access the application**:
   You can access the application by navigating to `http://localhost:8080/health` in your web browser or using a tool like `curl`:

   ```bash
   curl http://localhost:8080/health
   ```

   You should receive a JSON response indicating the status of the application.

## Using the Makefile

The project includes a Makefile to simplify common tasks. Here are the available commands:

- **Build the Docker image**:
  ```bash
  make build
  ```

- **Run the Docker container**:
  ```bash
  make run
  ```

- **Stop the running Docker container**:
  ```bash
  make stop
  ```

- **Remove the Docker container**:
  ```bash
  make clean
  ```

- **Access the health endpoint**:
  ```bash
  make access
  ```

- **Show help message**:
  ```bash
  make help
  ```

You can run these commands from the root of the project directory.
