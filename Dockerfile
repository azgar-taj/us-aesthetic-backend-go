# Use the official Go image as the base image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod ./

# Download and install the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/service

# Expose the port that the application listens on
EXPOSE 8080

# Set the entry point for the container
CMD ["./app"]