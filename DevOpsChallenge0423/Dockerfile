# Use the official Golang image as the base image
FROM golang:1.20.3

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files, if you have them, to fetch dependencies
# COPY go.mod go.sum ./

# Install dependencies
# RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o server .

# Expose the server port
EXPOSE 8083

# Run the server
CMD ["./server"]
