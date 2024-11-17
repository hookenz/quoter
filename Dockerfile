# Step 1: Build stage
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Build the Go app
RUN go build -o quoter .

# Step 2: Final stage
FROM scratch

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/quoter /app/quoter

# Command to run the executable
CMD ["./quoter"]

