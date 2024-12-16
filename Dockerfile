# Step 1: Use the official Go image as the base image
FROM golang:1.22 AS builder

# Step 2: Set the current working directory
WORKDIR /app

# Step 3: Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Step 4: Copy the entire application code into the container
COPY . .

# Step 5: Build the Go application
RUN go build -o main .

# Step 6: Use a smaller image for production
FROM debian:bookworm-slim

# Step 7: Install required dependencies
RUN apt-get update && apt-get install -y ca-certificates

# Step 8: Set working directory and copy the built app
WORKDIR /root/
COPY --from=builder /app/main .
COPY .env .env

# Step 9: Expose the port the app runs on
EXPOSE 8080

# Step 10: Start the application
CMD ["./main"]
