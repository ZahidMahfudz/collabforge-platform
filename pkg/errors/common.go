package errors

import "net/http"

var (
	ErrInternalServer    = New("Internal Server Error", "", http.StatusInternalServerError)
	ErrNotFound          = New("Resource Not Found", "", http.StatusNotFound)
	ErrUnauthorized      = New("Unauthorized", "", http.StatusUnauthorized)
	ErrForbidden         = New("Forbidden", "", http.StatusForbidden)
	ErrBadRequest        = New("Bad Request", "", http.StatusBadRequest)
	ErrUserNotFound      = New("User Not Found", "", http.StatusNotFound)
	ErrEmailAlreadyUsed  = New("Email Already Used", "", http.StatusBadRequest)
	ErrInvalidPassword   = New("Invalid Password", "", http.StatusUnauthorized)
	ErrInvalidToken      = New("Invalid Token", "", http.StatusUnauthorized)
	ErrTokenExpired      = New("Token Expired", "", http.StatusUnauthorized)
	ErrTokenNotValid     = New("Token Not Valid", "", http.StatusUnauthorized)
	ErrTokenNotFound     = New("Token Not Found", "", http.StatusUnauthorized)
	ErrTokenNotGenerated = New("Token Not Generated", "", http.StatusUnauthorized)
)
