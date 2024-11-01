# Gunakan image dasar Go untuk build
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod dan go sum untuk dependency
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh kode ke dalam image
COPY . .

# Build aplikasi
RUN go build -o main cmd/main.go

# Stage kedua untuk image yang lebih ringan
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy binary dari stage pertama
COPY --from=builder /app/main .

# Set port
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]
