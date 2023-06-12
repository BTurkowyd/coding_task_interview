# Start from a base Golang image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Expose the port 3000
EXPOSE 3000

# Build the application
RUN go build -o main .


# Set the command to run the application
CMD ["./main"]
