# Start from a Debian-based image with the Go 1.16 version installed
FROM golang:1.22

# Create a directory inside the container to store all our application and then make it the working directory.
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o backend .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./backend"]
