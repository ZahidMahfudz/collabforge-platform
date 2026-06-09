# Authentication API

## Register

### Endpoint
```
POST /auth/register
```

### Authentication
```
Tidak Diperlukan
```

### Request Body
```
{
    "first_name" : "Zahid",
    "mid_name" : "Muhammad",
    "last_name" : "Mahfudz",
    "username" : "zaemahfudz",
    "email" : "zaemahfudz@gmail.com",
    "password" : "Password@123",
}
```

### Validation
|Field|Rule|
|---|---|
|first_name|required, Min=3, max=50|
|mid_name| max=50|
|last_name|required, Min=3, max=50|
|username|required, Min=3, max=50|
|email|required, Must be valid email|
|password|required, must contain uppercase, lowercase, number, and special character|

### success Response
```
201 created

{
    "success": true,
    "message": "register success",
    "data": {
        "id": "usr_RnU27hXU7EVajMQ3",
        "name": "Zahid Muhammad Mahfudz",
        "email": "zaemahfudz@gmail.com"
    }
}
```

### Error Response
```
409 Confilct

{
    "success": false,
    "message": "email sudah ada",
    "details": "EMAIL_ALREADY_EXISTS"
}
```

### Notes
- masih sementara, masih bisa diganti

<br>

## Login
Masuk kedalam sistem dan membuat autentikasi oengguna

### Endpoint
```
POST /auth/login
```

### Authentication
```
Tidak Diperlukan
```

### Request Body
```
{
    "email" : "zaemahfudz@gmail.com",
    "password" : "Password@123",
}
```

### Validation
|Field|Rule|
|---|---|
|email|required|
|password|required|

### success Response
```
200 OK

{
    "success": true,
    "message": "login success",
    "data": {
        "id": "usr_RnU27hXU7EVajMQ3",
        "first_name": "Zahid",
        "last_name": "Mahfudz",
        "mid_name": "Muhammad",
        "username": "zaemahfudz",
        "email": "zaemahfudz@gmail.com",
        "access_token": "v4.local.{token}"
    }
}
```

### Cookie

|name|Value|Domain|Path|Expires|HttpOnly|Secure|
|---|---|---|---|---|---|---|
|refresh_token|v4.local.{token}|localhost|/|session|true|false|

### Error Response
```
401 Unauthorize

{
    "success": false,
    "message": "email atau password salah",
    "details": "INVALID_CREDENTIALS"
}
```

```
500 Internal Server Error

{
    "success": false,
    "message": "gagal menghasilkan token",
    "details": "FAILED_TO_GENERATE_TOKEN"
}
```

### Notes
cookie setting akan berubah saat production nanti

<br>

## Refresh

### Endpoint
```
POST /auth/refresh
```

### Authentication
```
cookie

refresh_token
```

### Request Body
```
Tidak diperlukan
```

### success Response
```
200 OK

{
    "success": true,
    "message": "refresh token success",
    "data": {
        "access_token": "v4.local.{token}"
    }
}
```

### Cookie

|name|Value|Domain|Path|Expires|HttpOnly|Secure|
|---|---|---|---|---|---|---|
|refresh_token|v4.local.{token}|localhost|/|session|true|false|

### Error Response
```
401 Unauthorize

{
    "success": false,
    "message": "refresh token tidak ditemukan",
    "details": "REFRESH_TOKEN_NOT_FOUND"
}
```
```
401 Unauthorize

{
    "success": false,
    "message": "refresh token tidak valid"",
    "details": "INVALID_REFRESH_TOKEN"
}
```

```
500 Internal Server Error

{
    "success": false,
    "message": "gagal menghasilkan token",
    "details": "FAILED_TO_GENERATE_TOKEN"
}
```

### Notes
Refresh Token yang tersimpan di cookie terganti dengan refresh token yang baru