# Use the latest Go version as a builder
FROM golang:latest AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Ensure the binary is built for the correct architecture and is executable
RUN GOOS=linux GOARCH=amd64 go build -o pokedox && chmod +x pokedox

# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Install libc (needed for some Go binaries)
RUN apk add --no-cache libc6-compat

# Copy the compiled binary from the builder stage
COPY --from=builder /app/pokedox .

# Ensure the binary has execution permissions
RUN chmod +x pokedox

# Set the entry point to run the application
ENTRYPOINT ["./pokedox"]
