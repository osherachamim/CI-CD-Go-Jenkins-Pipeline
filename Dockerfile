# Use the latest official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY hello-world.go .

# Build the Go application
RUN go build -o hello-world hello-world.go

# Expose the application’s port
EXPOSE 8081

# Command to run the executable
CMD ["./hello-world"]
