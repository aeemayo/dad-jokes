# Start with a golang base image for building
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Install git and ca-certificates (needed for dependencies and HTTPS)
RUN apk add --no-cache git ca-certificates

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
# CGO_ENABLED=0 creates a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Start a new stage with a small alpine image
FROM alpine:latest

# Install ca-certificates to ensure we can make HTTPS requests (OpenRouter, Teneo)
RUN apk add --no-cache ca-certificates

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy .env file if it exists (Optional: usually secrets are set in Render dashboard)
# COPY .env . 

# Command to run the executable
CMD ["./main"]
