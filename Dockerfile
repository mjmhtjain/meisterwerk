# Builder stage
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/meisterwerk ./cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/meisterwerk .

# Copy any config files if needed (add these later if required)
# COPY config/ config/

# Expose the port your application runs on
ENV PORT=8080
EXPOSE 8080

# Run the binary
CMD ["./meisterwerk"] 