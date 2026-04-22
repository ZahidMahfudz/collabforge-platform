# AUTH REGISTER SPEC

## TUJUAN

Implementasi fitur register:

1. Register manual (email + password)
2. Register dengan Google (langsung dianggap login)
3. Menghasilkan PASETO token setelah berhasil

---

## ATURAN WAJIB

* Gunakan clean architecture (controller → usecase → repository)
* Gunakan pgx (raw SQL)
* Gunakan bcrypt untuk hash password
* Gunakan ID custom (prefix + random)
* Prefix user: "usr"
* Gunakan pkg/errors untuk error
* Gunakan pkg/response untuk response

DILARANG:

* business logic di controller
* query database di controller

---

## GENERATE ID

FORMAT:
prefix_random

CONTOH:
usr_xxxxxxxxxxxxxxxx

RULE:

* prefix: 3 karakter
* separator: "_"
* random: 16 karakter
* total panjang: 20 karakter
* karakter: a-zA-Z0-9
* gunakan: crypto/rand
* tidak boleh: auto increment / math random

IMPLEMENTASI:

* buat di: pkg/utils/id_generator.go
* function:
  GenerateID(prefix string, length int) (string, error)

---

## DATABASE

TABLE: users

FIELD:

* id (VARCHAR(20), primary key)
* name (VARCHAR(255), NOT NULL)
* email (VARCHAR(255), NOT NULL, UNIQUE)
* password (VARCHAR(255), NULL)
* provider (VARCHAR(20), NOT NULL, default 'local')
* provider_id (VARCHAR(255), NULL)
* created_at (TIMESTAMP, NOT NULL, default CURRENT_TIMESTAMP)
* updated_at (TIMESTAMP, NOT NULL, default CURRENT_TIMESTAMP)

---

## MIGRATION SQL

```sql
CREATE TABLE users (
    id VARCHAR(20) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255),
    provider VARCHAR(20) NOT NULL DEFAULT 'local',
    provider_id VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);

ALTER TABLE users
ADD CONSTRAINT users_provider_check
CHECK (provider IN ('local', 'google'));
```

---

## ENTITY

```go
type User struct {
    ID         string
    Name       string
    Email      string
    Password   *string
    Provider   string
    ProviderID *string
    CreatedAt  time.Time
    UpdatedAt  time.Time
}
```

---

## ENDPOINT

### 1. REGISTER MANUAL

POST /auth/register

REQUEST:
{
"name": "string",
"email": "string",
"password": "string"
}

FLOW:

1. Validasi input (required, email valid, password >= 6)
2. Cek email sudah ada atau belum

   * jika ada → EMAIL_ALREADY_EXISTS
3. Hash password (bcrypt)
4. Generate ID (prefix "usr")
5. Simpan user (provider = "local")
6. Generate PASETO token
7. Return token + user

---

### 2. REGISTER GOOGLE

POST /auth/google

REQUEST:
{
"id_token": "string"
}

FLOW:

1. Verifikasi id_token ke Google
2. Ambil data:

   * email
   * name
   * provider_id
3. Cek user berdasarkan email

   * jika sudah ada → gunakan user tersebut
   * jika belum:
     a. generate ID
     b. simpan user (password NULL, provider = "google")
4. Generate PASETO token
5. Return token + user

---

## TOKEN (PASETO)

Gunakan PASETO (Platform-Agnostic Security Token)

VERSI:

* v4.local

PAYLOAD:
{
"user_id": "string",
"email": "string",
"exp": "timestamp"
}

RULE:

* secret key dari .env
* minimal 32 karakter
* expired: 24 jam
* gunakan encryption (local mode)

---

## RESPONSE (WAJIB FORMAT GLOBAL)

SUCCESS:
{
"success": true,
"message": "register success",
"data": {
"token": "string",
"user": {
"id": "string",
"name": "string",
"email": "string"
}
}
}

ERROR:
{
"success": false,
"message": "error message",
"error": "ERROR_CODE"
}

---

## ERROR YANG DIGUNAKAN

* EMAIL_ALREADY_EXISTS
* BAD_REQUEST
* INTERNAL_SERVER_ERROR

---

## STRUKTUR YANG HARUS DIBUAT

* entity: user
* repository interface: user_repository
* repository implementation (pgx)
* usecase: auth_usecase
* controller: auth_controller
* route: auth_route

---

## SECURITY

* password wajib di-hash (bcrypt)
* jangan simpan password plain text
* jangan return password ke client
* wajib verifikasi token Google
* jangan expose error internal

---

## OUTPUT

* PASETO token
* user (id, name, email)

---

## DONE

* register manual berhasil
* email tidak bisa duplikat
* password tersimpan dalam bentuk hash
* register google berhasil
* PASETO token berhasil dibuat
* response sesuai format global
* tidak ada business logic di controller
