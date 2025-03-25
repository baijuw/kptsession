# Use a Go base image to build the application
FROM golang:1.24.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . /app

# Build the Go application
RUN go build -o gohttplight

# Use a minimal Alpine Linux image for the final image
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

# Set the working directory in the final image
WORKDIR /root

# Copy the built binary from the builder stage
COPY --from=builder /app/gohttplight /root/gohttplight

# Expose the port that the application listens on
EXPOSE 8080

# Set the command to run when the container starts
CMD ["/root/gohttplight"]
