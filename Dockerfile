# Builder Stage
FROM golang:1.23.2 AS builder

# Set the working directory in the builder
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application with static linking
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ution ./cmd/ution.go

# Production Stage
FROM alpine:latest

# Set the working directory for the final image
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/ution .

# Ensure the binary has executable permissions
RUN chmod +x ./ution

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./ution"]
