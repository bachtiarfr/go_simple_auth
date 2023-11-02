# simple-auth API

Teckstack : Golang, PostgreSQL

## Instalasi

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

### 2. Autentikasi (Register)

- **URL**: `http://127.0.0.1:8080/api/v1/register`
- **Metode**: `POST`
- **Deskripsi**: Endpoint ini digunakan untuk melakukan autentikasi.
- **Request Body**:
  ```json
  {
    "fullname": "fatur",
    "email": "test2@example.com",
    "age": 24,
    "phone_number": "0895334568841",
    "password": "password2"
  }
  ```

### 3. Get All Users

- **URL**: `http://127.0.0.1:8080/api/v1/users`
- **Metode**: `GET`
- **Deskripsi**: Endpoint ini digunakan untuk mendapatkan data semua user.
- **Response Body**:
  ```json
  {
    "code": "00",
    "message": "success",
    "data": [
      {
        "user_id": 1,
        "fullname": "Test User",
        "email": "test@example.com",
        "age": 35,
        "phone_number": "+62895334568841"
      },
      {
        "user_id": 2,
        "fullname": "fatur",
        "email": "test2@example.com",
        "age": 24,
        "phone_number": "0895334568841"
      }
    ]
  }
  ```

### 4. Get Affiliated User (based on access_token)

- **URL**: `http://127.0.0.1:8080/api/v1/user`
- **Metode**: `GET`
- **Deskripsi**: Endpoint ini digunakan untuk mendapatkan data user yang terafiliasi dengan token_access.
- **Response Body**:
  ```json
  {
    "code": "00",
    "message": "success",
    "data": {
      "user_id": 1,
      "fullname": "Test User",
      "email": "test@example.com",
      "age": 35,
      "phone_number": "+62895334568841"
    }
  }
  ```
