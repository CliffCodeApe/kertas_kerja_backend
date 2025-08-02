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
	// Document errors
	ErrDocumentNotFound = NewNotFoundError("Dokumen tidak ditemukan")
	ErrInvalidDocID     = NewBadRequest("Dokumen tidak valid")

	// Auth errors
	ErrUserNotAuthorized = NewUnauthorizedError("Anda tidak bisa mengakses halaman ini")
	ErrUserNotVerified   = NewUnauthorizedError("Akun belum diverifikasi. Silakan lakukan verifikasi email anda")
	// User Errors
	ErrUserNotFound = NewNotFoundError("User ini tidak ditemukan")

	// Cover errors
	ErrCoverNotFound = NewNotFoundError("Cover tidak ditemukan")

	// Abstract errors
	ErrAbstractNotFound = NewNotFoundError("Abstract tidak ditemukan")

	// Acknowledgement errors
	ErrAcknowledgementNotFound = NewNotFoundError("Acknowledgement tidak ditemukan")

	// Chapter errors
	ErrChapterUnauthorized = NewUnauthorizedError("Anda tidak bisa mengakses chapter ini")
	ErrChapterNotFound     = NewNotFoundError("Chapter tidak ditemukan")
	ErrInvalidChapterID    = NewBadRequest("Chapter tidak valid")
	ErrChapterTypeNotMatch = NewBadRequest("Parent object tidak valid")
	ErrFirstChapters       = NewBadRequest("Chapter induk tidak ada pada dokumen ini")
	// ErrPreviousObjectNotValid = NewBadRequest("Previous object tidak valid")
	ErrPreviousObjectInvalid      = NewBadRequest("Previous object tidak valid")
	ErrPreviousObjectTypeNotMatch = NewBadRequest("Tipe previous object tidak cocok")

	// Previous Object errors
	ErrSubAndChapterMismatch = NewBadRequest("Sub chapter/point tidak bisa memiliki chapter sebagai previous object")
	ErrParentObjectMismatch  = NewBadRequest("Parent dari object tidak setara dengan parent dari previous object")
	ErrChapterAndSubMismatch = NewBadRequest("Chapter tidak bisa memiliki sub chapter/point sebagai previous object")

	// Table errors
	ErrTableNotFound = NewNotFoundError("Table tidak ditemukan")

	// Picture errors
	ErrPictureNotFound = NewNotFoundError("Gambar tidak ditemukan")

	// Text errors
	ErrTextNotFound = NewNotFoundError("Text tidak ditemukan")

	// Reset Password Error
	ErrTokenNotFound = NewNotFoundError("Token tidak di temukan")
	ErrTokenExpired  = NewUnauthorizedError("Token telah kadaluarsa")
	ErrTokenLimit    = NewTooManyRequestsError("Anda telah melewati batas request")
	ErrTokenNotValid = NewUnauthorizedError("Token tidak Valid")

	// Common errors
	ErrServer             = NewInternalServerError("Terjadi Kesalahan")
	ErrValid              = NewBadRequest("input tidak valid/user tidak ditemukan")
	ErrRequestBody        = NewBadRequest("Request body tidak valid")
	ErrNamaSatkerNotFound = NewNotFoundError("Nama Satker tidak ditemukan")
	ErrLoginFailed        = NewNotFoundError("Nama Satker atau password tidak ditemukan")
	ErrDecryptFailed      = NewNotFoundError("Dekripsi password gagal")

	// storage errors
	ErrInvalidFileType  = NewBadRequest("file type is not allowed")
	ErrFileSizeTooLarge = NewBadRequest("file size is too large")
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
