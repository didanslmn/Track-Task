package shared

import (
	"errors"
	"fmt"
	"net/http"
)

type ErrorCode string

const (
	CodeValidation   ErrorCode = "VALIDATION_ERROR"
	CodeUnauthorized ErrorCode = "UNAUTHORIZED"
	CodeNotFound     ErrorCode = "NOT_FOUND"
	CodeInternal     ErrorCode = "INTERNAL_ERROR"
	CodeBadRequest   ErrorCode = "BAD_REQUEST"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type AppError struct {
	HTTPStatus int
	Code       ErrorCode
	Message    string
	Details    []ValidationError
	RawErr     error
}

func (e *AppError) Error() string {
	if e.RawErr != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Code, e.Message, e.RawErr)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *AppError) Unwarp() error {
	return e.RawErr
}

func (e AppError) WithError(err error) *AppError {
	e.RawErr = err
	return &e
}

func (e AppError) WithDetail(detail []ValidationError) *AppError {
	e.Details = detail
	return &e
}

func (e AppError) WithMessage(message string) *AppError {
	e.Message = message
	return &e
}

// predifined error code

var (
	ErrValidation = AppError{
		Code: CodeValidation, Message: "Validasi gagal", HTTPStatus: http.StatusBadRequest,
	}
	ErrUnautorized = AppError{
		Code: CodeUnauthorized, Message: "Autentikasi diperlukan", HTTPStatus: http.StatusUnauthorized,
	}
	ErrNotFound = AppError{
		Code: CodeNotFound, Message: "Resource Tidak Ditekukan", HTTPStatus: http.StatusNotFound,
	}
	ErrInternal = AppError{
		Code: CodeInternal, Message: "Terjadi Kesalahan Internal", HTTPStatus: http.StatusInternalServerError,
	}
	ErrBadRequest = AppError{
		Code: CodeBadRequest, Message: "Permintaan tidak valid", HTTPStatus: http.StatusBadRequest,
	}
)

// Constructor function for common HTTP errors

func NewValidationError(details []ValidationError) *AppError {
	return ErrValidation.WithDetail(details)
}

func NewNotFoundError(resource string) *AppError {
	return ErrNotFound.WithMessage(resource)
}

func WarpInternal(err error) *AppError {
	return ErrInternal.WithError(err)
}

// Error-checking helpers

func IsAppError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr)
}

func AsAppError(err error) (*AppError, bool) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr, true
	}
	return nil, false
}
