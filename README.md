# simple-auth API

Teckstack : Golang, PostgreSQL

## Pull repository

Clone terlebih dahulu dari link repository

```bash
git clone https://github.com/bachtiarfr/go_simple_auth.git
```

Setelah berhasil diclone ke folder anda, masuk kedalam folder project nya kemudian jalankan perintah migration untuk migration table nya, makesure sudah setup database nya dengan benar, database menggunakan postgresql

## Database Migration

untuk migration up jalankan

```bash
go run cmd/migration.go up
```

untuk migration down jalankan

```bash
go run cmd/migration.go down
```

## List endpoin API

### 1. Autentikasi (Login)

- **URL**: `http://127.0.0.1:8080/api/v1/login`
- **Metode**: `POST`
- **Deskripsi**: Endpoint ini digunakan untuk melakukan autentikasi.
- **Request Body**:
  ```json
  {
    "email": "test@example.com",
    "password": "password"
  }
  ```

````
### 2. Autentikasi (Register)

- **URL**: `http://127.0.0.1:8080/api/v1/register`
- **Metode**: `POST`
- **Deskripsi**: Endpoint ini digunakan untuk melakukan autentikasi.
- **Request Body**:
  ```json
  {
    "email": "test@example.com",
    "password": "password"
  }
````
