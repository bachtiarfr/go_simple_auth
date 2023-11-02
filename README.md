# simple-auth API

Teckstack : Golang, PostgreSQL

# Pull repository

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
