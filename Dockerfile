# Gunakan image Golang sebagai base image
FROM golang:1.20-alpine

# Set working directory di dalam container
WORKDIR /app

# Copy isi proyek Anda ke dalam working directory di dalam container
COPY . .

# Build proyek Golang
RUN go build -o main .

# Command yang akan dijalankan ketika container dijalankan
CMD ["./main"]