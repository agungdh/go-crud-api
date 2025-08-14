# --- Stage 1: Build ---
FROM golang:1.25 AS builder

# Set destination for COPY commands
WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build with optimizations (no CGO, smaller binary)
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# --- Stage 2: Run ---
FROM alpine:latest

# Set working directory inside container
WORKDIR /root/

# Copy binary from builder stage
COPY --from=builder /app/main .

# Set executable
CMD ["./main"]
    