# Use the official golang image as base image
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN make build

# Use a lightweight base image
FROM alpine:latest

# Set the working directory in the container
WORKDIR /app

# Copy the binary from the builder stage to the final stage
COPY --from=builder /app/go-api .

# Expose the port the application listens on
EXPOSE 8090

# Set application port to the same expose port
ENV PORT=:8090

# Run the binary
CMD ["./go-api"]
