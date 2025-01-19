# Start from the official Go image
FROM golang:1.22 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Use a smaller base image for the runtime
FROM gcr.io/distroless/base-debian11

# Set the working directory
WORKDIR /app

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your application listens on
EXPOSE 8080

# Run the binary
CMD ["./main"]
