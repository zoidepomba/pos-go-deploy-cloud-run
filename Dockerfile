# Use a lightweight Go image
FROM golang:1.23

# Set the working directory
WORKDIR /app

# Copy the Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main ./cmd/

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./main"]