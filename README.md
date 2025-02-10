# Meisterwerk

A Go-based REST API service for managing quotes and orders.

## Project Structure

```
.
├── cmd
│   └── main.go
├── internal
│   ├── database
│   │   └── migrations
│   │       └── 20250209201745_order.sql
│   ├── handlers
│   │   ├── health_handler.go
│   │   └── quote_handler.go
│   ├── models
│   │   ├── order.go
│   │   ├── product.go
│   │   └── quote.go
│   ├── repository
│   │   ├── order_repository.go
│   │   ├── product_repository.go
│   │   └── quote_repository.go
│   └── services
│       ├── order_service.go
│       └── quote_service.go
├── go.mod
└── README.md
```

## API Endpoints

### Health Check
- **GET** `/health`
  - Returns the health status of the service
  - Response: `{"status": "ok"}`

### Quotes
- **POST** `/api/v1/quote`
  - Creates a new quote and responds with the quote details, product details, total price, and total tax.
  - Request Body:
    ```json
    {
        "author": "A2",
        "customer_name": "Client1",
        "product_list": ["2", "4"]
    }
    ```
  - Response: 
    ```json
    {
        "id": "9b3a04d4-fd09-4d6f-9f96-6a12e0416029",
        "author": "A2",
        "customer_name": "Client1",
        "product_list": [
            {
                "id": "2",
                "name": "Product 2",
                "price": 200,
                "tax": 15.5
            },
            {
                "id": "4",
                "name": "Product 4",
                "price": 150,
                "tax": 8
            }
        ],
        "total_price": 350,
        "total_tax": 43,
        "status": "created"
    }
    ```
  - Status: 201 Created

- **GET** `/api/v1/quotes/{id}`
  - Retrieve a specific quote by id
  - Response Body:
    ```json
    {
        "id": "52ab8926-1622-4620-87b8-a2790f2e52ab",
        "author": "A1",
        "customer_name": "Client1",
        "product_list": [
            {
                "id": "2",
                "name": "Product 2",
                "price": 200,
                "tax": 15.5
            },
            {
                "id": "4",
                "name": "Product 4",
                "price": 150,
                "tax": 8
            }
        ],
        "total_price": 350,
        "total_tax": 43,
        "status": "created"
    }
    ``` 
  - Response: Quote object with status 200

- **PUT** `/api/v1/quote/{id}/status`
  - Updates the status of a quote
  - If the status is accepted, the quote is converted to an order
  - Request Body:
    ```json
    {
        "status": "accepted"
    }
    ```
  - Response Body:
    ```json
    {
        "message": "Quote status updated successfully"
    }
    ``` 
  - Response: Status 200

### Products
- **GET** `/api/v1/all-products`
  - List all available products
  - Response Body:
    ```json
    [
        {
            "id": "1",
            "name": "Product 1",
            "price": 100,
            "tax": 10
        },
        {
            "id": "2",
            "name": "Product 2",
            "price": 200,
            "tax": 15.5
        }
    ]
    ``` 
  - Response: Array of product objects with status 200

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
