# Start with a base image that has Go installed
FROM golang:1.23.1 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY main.go .
COPY go.mod .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o random_number_server .

# Use a lightweight base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the executable from the builder stage
COPY --from=builder /app/random_number_server .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./random_number_server"]

