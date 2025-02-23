# Stage 1: Build the Go application
FROM golang:1.21.1-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main .

# Stage 2: Run the Go application
FROM alpine:latest

WORKDIR /root/

# Install kubectl
RUN apk add --no-cache curl && \
    curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && \
    install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl && \
    rm kubectl

# Copy the built application from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/prod.env .env 

# Expose port 8082
EXPOSE 8082

# Run the application
CMD ["./main"]