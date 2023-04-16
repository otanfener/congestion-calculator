# Use a Golang builder container to build the binary
FROM golang:1.20-alpine AS builder

# Set the working directory
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod tidy

# Copy the source code
COPY . .

# Build the binary with CGO disabled
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o congestion-calculator ./cmd/

# Create a new image with only the binary
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder container
COPY --from=builder /app/congestion-calculator .


EXPOSE 3000
# Run the binary
CMD ["./congestion-calculator"]
