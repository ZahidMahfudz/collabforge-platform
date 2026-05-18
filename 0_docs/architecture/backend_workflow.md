# Backend Workflow

Dokumentasi ini menjelaskan alur kerja backend pada platform CollabForge.

---

# Arsitektur Backend

Backend CollabForge menggunakan pendekatan simple clean architecture dengan pemisahan tanggung jawab pada setiap layer agar kode lebih mudah dipahami, scalable, dan maintainable.

---

# Alur Request

```text
Client
   ↓
main.go
   ↓
routes
   ↓
middleware (auth & validasi request) -> req.locals (hasil validasi disimpan di req.locals)
   ↓
controller
   ↓
req.locals (jika method post)
   ↓
usecase
   ↓
repository / external service
   ↓
usecase
   ↓
dto response
   ↓
controller
   ↓
global response
   ↓
client
```

---

# Penjelasan Setiap Layer

## 1. main.go

Entry point aplikasi.

Tugas:
- load environment
- inisialisasi logger
- koneksi database
- inisialisasi Fiber
- register middleware
- register routes
- menjalankan server
- dependecy injection

---

## 2. Routes

Layer yang bertugas mendefinisikan endpoint API dan menghubungkan endpoint ke controller.

Contoh:

```go
authGroup.Post("/register", middleware.ValidateRequest[request.RegisterRequest](), authController.Register)
```
ditambahkan juga middleware untuk validasi request dengan memanggil validasi request kemudian [ini struct dari body req yang didefinisikan didalam dto/request] lalu setelahnya diteruskan kedalam controller
---

## 3. Middleware

Middleware berjalan sebelum controller.

Digunakan untuk:
- authentication
- logging
- rate limiting
- request tracking

Contoh flow:

```text
request masuk
   ↓
middleware cek token
   ↓
valid → lanjut controller
invalid → return error
```

---

## 4. Controller

Controller adalah layer HTTP.

Tugas:
- menerima request
- parsing body/query/params
- validasi dasar request
- memanggil usecase
- mengembalikan response

Controller tidak boleh:
- query database
- business logic
- generate token

---

## 5. DTO Request

DTO Request digunakan sebagai kontrak input API.

Contoh:

```go
type RegisterRequest struct {
    Name     string
    Email    string
    Password string
}
```

DTO hanya digunakan untuk:
- struktur request
- validasi ringan

---

## 6. Usecase

Usecase adalah inti business logic aplikasi.

Semua aturan bisnis utama diletakkan di layer ini.

Contoh:
- register user
- login
- create project
- assign task
- generate portfolio

Flow umum:

```text
controller
   ↓
usecase
   ↓
repository/service
```

---

## 7. Repository

Repository bertugas berinteraksi langsung dengan database menggunakan raw SQL (pgx).

Repository hanya berisi:
- query database
- mapping data

Repository tidak boleh:
- business logic
- token generation
- hashing password

---

## 8. External Service

Service digunakan untuk integrasi layanan eksternal.

Contoh:
- Google OAuth
- AI Service
- Email Service

Flow:

```text
usecase
   ↓
external service
```

---

## 9. DTO Response

DTO Response digunakan untuk membentuk output API agar konsisten dan aman.

Tujuan:
- menyembunyikan field sensitif
- menjaga konsistensi response

---

## 10. Global Response

Semua response menggunakan wrapper global.

Contoh success response:

```json
{
  "success": true,
  "message": "success",
  "data": {}
}
```

Contoh error response:

```json
{
  "success": false,
  "message": "error message",
  "error": "ERROR_CODE"
}
```

---

# Contoh Flow Register

```text
POST /auth/register
```

Flow:

```text
Client
   ↓
Route
   ↓
Middleware
   ↓
Controller
   ↓
DTO Request
   ↓
Usecase
   ↓
Repository (cek email)
   ↓
Repository (create user)
   ↓
Generate Paseto
   ↓
DTO Response
   ↓
Global Response
   ↓
Client
```

---

# Prinsip Arsitektur

## Controller
Fokus pada HTTP layer.

## Usecase
Fokus pada business logic.

## Repository
Fokus pada database access.

## Service
Fokus pada external integration.

---

# Tujuan Arsitektur Ini

- mudah dipahami
- scalable
- maintainable
- mudah di-debug
- AI-agent friendly
- cocok untuk raw SQL
- cocok untuk modular development