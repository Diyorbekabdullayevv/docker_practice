# -------- Build Stage --------
FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server .

# -------- Run Stage --------
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]


# # Use the official Golang image
# FROM golang:1.22

# # Set working directory inside container
# WORKDIR /app

# # Copy go.mod and go.sum first (if available)
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy the rest of the source code
# COPY . .

# # Build the Go app (name it "server")
# RUN go build -o server .

# # Tell Docker which port the app listens on
# EXPOSE 8080

# # Run the app
# CMD ["./server"]