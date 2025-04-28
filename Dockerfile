# Start from official Golang image
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of your code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port your app runs on
EXPOSE 9000

# Command to run the executable
CMD ["./main"]
