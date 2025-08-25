package errs

import (
	"net/http"
)

type MessageErr interface {
	Message() string
	Status() int
	Error() string
}

type ErrorData struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

var (

	// Auth errors
	ErrUserNotAuthorized = NewUnauthorizedError("Anda tidak bisa mengakses halaman ini")
	ErrUserNotVerified   = NewUnauthorizedError("Akun belum divalidasi. Silakan hubungi Super Admin untuk memvalidasi akun anda")
	ErrUserAlreadyExists = NewBadRequest("User sudah terdaftar")
	// User Errors
	ErrUserNotFound = NewNotFoundError("User ini tidak ditemukan")

	// Common errors
	ErrServer             = NewInternalServerError("Terjadi Kesalahan")
	ErrValid              = NewBadRequest("input tidak valid/user tidak ditemukan")
	ErrRequestBody        = NewBadRequest("Request body tidak valid")
	ErrNamaSatkerNotFound = NewNotFoundError("Nama Satker tidak ditemukan")
	ErrLoginFailed        = NewNotFoundError("Nama Satker atau password tidak ditemukan")
	ErrDecryptFailed      = NewNotFoundError("Dekripsi password gagal")
	ErrForbidden          = NewForbiddenError("Permintaan tidak diizinkan")

	// Data Lelang Errors
	ErrDataLelangNotFound = NewNotFoundError("Data lelang tidak ditemukan")
)

func (e *ErrorData) Message() string {
	return e.ErrMessage
}

func (e *ErrorData) Status() int {
	return e.ErrStatus
}

func (e *ErrorData) Error() string {
	return e.ErrError
}

func NewUnauthorizedError(message string) MessageErr {
	return &ErrorData{
		ErrStatus:  http.StatusForbidden,
		ErrMessage: message,
		ErrError:   "NOT_AUTHORIZED",
	}
}

func NewUnauthenticatedError(message string) MessageErr {
	return &ErrorData{
		ErrStatus:  http.StatusUnauthorized,
		ErrMessage: message,
		ErrError:   "NOT_AUTHENTICATED",
	}
}

func NewNotFoundError(message string) MessageErr {
	return &ErrorData{
		ErrStatus:  http.StatusNotFound,
		ErrMessage: message,
		ErrError:   "NOT_FOUND",
	}
}

func NewBadRequest(message string) MessageErr {
	return &ErrorData{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: message,
		ErrError:   "BAD_REQUEST",
	}
}

func NewInternalServerError(message string) MessageErr {
	return &ErrorData{
		ErrStatus:  http.StatusInternalServerError, //500
		ErrMessage: message,
		ErrError:   "INTERNAL_SERVER_ERROR",
	}
}

func NewUnprocessibleEntityError(message string) MessageErr {
	return &ErrorData{
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrMessage: message,
		ErrError:   "INVALID_REQUEST_BODY",
	}
}

func NewTooManyRequestsError(message string) MessageErr {
	return &ErrorData{
		ErrStatus:  http.StatusTooManyRequests,
		ErrMessage: message,
		ErrError:   "TOO_MANY_REQUEST",
	}
}

func NewForbiddenError(message string) MessageErr {
	return &ErrorData{
		ErrStatus:  http.StatusForbidden,
		ErrMessage: message,
		ErrError:   "FORBIDDEN",
	}
}
