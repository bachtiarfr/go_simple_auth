# Gunakan image resmi Golang sebagai dasar
FROM golang:1.18

# Set working directory di dalam kontainer
WORKDIR /app

# Salin file-file yang diperlukan ke dalam kontainer
COPY . .

# Kompilasi aplikasi Golang
RUN go build -o main

# Port yang akan digunakan oleh aplikasi
EXPOSE 8080

# Menjalankan aplikasi saat kontainer dijalankan
CMD ["./main"]
