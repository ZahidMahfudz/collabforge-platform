# Plan: Logout Feature

## Overview
Implementasi fitur logout yang memungkinkan user untuk mengakhiri sesi dengan cara merevoke refresh token di database dan menghapus cookie refresh token di client.

## Flow Diagram
```
Client (POST /auth/logout)
    │
    ▼
[Extract refresh_token from Cookie]
    │
    ▼
[Controller: Logout]
    │
    ▼
[UseCase: Logout]
    ├── Verifikasi refresh token (PASETO verify)
    ├── Hash refresh token
    ├── Cari refresh token di database (FindByToken)
    ├── Cek apakah token sudah direvoke
    ├── Revoke refresh token di database (RevokeToken)
    └── Return response
    │
    ▼
[Controller: Hapus cookie refresh_token]
    │
    ▼
[Return Success Response ke Client]
```

## Task List

### Task 1: Tambah Response DTO untuk Logout
**File:** `internal/dto/response/auth_response.go`

Tambahkan struct `LogoutResponse`:
```go
type LogoutResponse struct {
    Message string `json:"message"`
}
```

### Task 2: Tambah Method Logout di UseCase
**File:** `internal/usecase/auth_usecase.go`

Tambahkan method `Logout` pada struct `AuthUseCase`:
```go
func (u *AuthUseCase) Logout(ctx context.Context, refreshToken string) (*dtoresponse.LogoutResponse, error) {
    // 1. Verifikasi refresh token menggunakan PasetoService.VerifyToken
    // 2. Cek tipe token harus "refresh"
    // 3. Hash refresh token menggunakan utils.HashToken
    // 4. Cari refresh token di database menggunakan refreshTokenRepo.FindByToken
    // 5. Cek apakah token sudah direvoke (storedToken.Revoked != nil)
    // 6. Revoke refresh token menggunakan refreshTokenRepo.RevokeToken
    // 7. Return LogoutResponse
}
```

**Detail Logic:**
- Jika verifikasi token gagal → return error `"INVALID_REFRESH_TOKEN"`
- Jika tipe token bukan "refresh" → return error `"INVALID_REFRESH_TOKEN"`
- Jika token tidak ditemukan di database → return error `"INVALID_REFRESH_TOKEN"`
- Jika token sudah direvoke → return error `"REFRESH_TOKEN_ALREADY_REVOKED"`
- Jika revoke gagal → return error `"FAILED_TO_REVOKE_TOKEN"`
- Jika sukses → return `LogoutResponse{Message: "logout berhasil"}`

### Task 3: Tambah Method Logout di Controller
**File:** `internal/controller/auth_controller.go`

Tambahkan method `Logout` pada struct `AuthController`:
```go
func (c *AuthController) Logout(ctx *fiber.Ctx) error {
    // 1. Ambil refresh_token dari cookie
    // 2. Jika kosong, return error 401 "REFRESH_TOKEN_NOT_FOUND"
    // 3. Panggil authUseCase.Logout(ctx.Context(), refreshToken)
    // 4. Handle error dari usecase:
    //    - "INVALID_REFRESH_TOKEN" → 401
    //    - "REFRESH_TOKEN_ALREADY_REVOKED" → 400
    //    - "FAILED_TO_REVOKE_TOKEN" → 500
    // 5. Hapus cookie refresh_token (set value kosong, MaxAge -1)
    // 6. Return success response 200
}
```

**Detail penghapusan cookie:**
```go
ctx.Cookie(&fiber.Cookie{
    Name:     "refresh_token",
    Value:    "",
    HTTPOnly: true,
    Secure:   false,
    SameSite: "lax",
    MaxAge:   -1, // menghapus cookie
})
```

### Task 4: Tambah Route Logout
**File:** `internal/routes/auth_routes.go`

Tambahkan endpoint:
```go
// endpoint logout
authGroup.Post("/logout", authController.Logout)
```

## Error Handling Summary

| Kondisi | Error Code | HTTP Status |
|---------|-----------|-------------|
| Cookie refresh_token kosong | `REFRESH_TOKEN_NOT_FOUND` | 401 |
| Token tidak valid (signature/expired) | `INVALID_REFRESH_TOKEN` | 401 |
| Token bukan tipe refresh | `INVALID_REFRESH_TOKEN` | 401 |
| Token tidak ditemukan di DB | `INVALID_REFRESH_TOKEN` | 401 |
| Token sudah direvoke sebelumnya | `REFRESH_TOKEN_ALREADY_REVOKED` | 400 |
| Gagal revoke di database | `FAILED_TO_REVOKE_TOKEN` | 500 |
| Sukses logout | - | 200 |

## Expected Success Response
```json
{
    "success": true,
    "message": "logout success",
    "data": {
        "message": "logout berhasil"
    }
}
```

## Expected Error Response (contoh)
```json
{
    "success": false,
    "message": "refresh token tidak valid",
    "details": "INVALID_REFRESH_TOKEN"
}
```

## Execution Order
1. Task 1 → Tambah `LogoutResponse` DTO
2. Task 2 → Tambah `Logout` method di UseCase
3. Task 3 → Tambah `Logout` method di Controller
4. Task 4 → Tambah route `/auth/logout`

## Files yang Akan Dimodifikasi
- `internal/dto/response/auth_response.go`
- `internal/usecase/auth_usecase.go`
- `internal/controller/auth_controller.go`
- `internal/routes/auth_routes.go`
