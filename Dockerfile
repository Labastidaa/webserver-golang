# Base image to run container with Go App
FROM golang:1.22.1-alpine3.19  AS builder

WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Cleanup unused modules in go.mod - minimize size of module cache and docker image
RUN go mod tidy

# Copy the rest of the application code
COPY . ./

# Compiles Go App into binary.
RUN go build -o /app/go-blckchn ./cmd/server/

# Start a new base image for running the application
FROM alpine:3.19

# Installs the necessary certificates for SSL/TLS communication, ensuring that your Go application can make secure network requests. 
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage to the new image
COPY --from=builder /app/go-blckchn /app/go-blckchn

# Set the working directory
WORKDIR /app

# Set a non-root user (recommended for security)
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Set environment variables and expose port
ENV COINMARKETCAP_API_KEY=${COINMARKETCAP_API_KEY}
EXPOSE 8081

# Run the binary when the container starts
CMD ["/app/go-blckchn"]
