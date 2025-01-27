# Stage 1: Build the Go application
FROM golang:1.21.1-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main .

# Stage 2: Run the Go application
FROM alpine:latest

WORKDIR /root/

# Copy the built application from the builder stage
COPY --from=builder /app/main .

# Expose port 8082
EXPOSE 8082

# Run the application
CMD ["./main"]