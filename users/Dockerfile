FROM golang:latest

LABEL maintainer="Cherbadzhi Maxim"

WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Test the Go service
RUN go test -v -mod=mod -race -timeout 30s -short ./...

# Build the Go service
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=mod .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run migrations
CMD ["/app/users", "migrate", "up"]
# Command to run the executable
CMD ["/app/users", "api"]