# Gunakan image Golang sebagai base image
FROM golang:1.20-alpine

# Set working directory di dalam container
WORKDIR /app

# Copy isi proyek Anda ke dalam working directory di dalam container
COPY . .
COPY .env .env

# Build proyek Golang
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command yang akan dijalankan ketika container dijalankan
CMD ["./main"]