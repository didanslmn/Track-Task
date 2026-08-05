package shared

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

//-------------------------------------
// Standard API Response Structure
//-------------------------------------

type Meta struct {
	Page       int   `json:"page,omitempty"`
	Limit      int   `json:"limit,omitempty"`
	Total      int64 `json:"total,omitempty"`
	TotalPages int   `json:"total_pages,omitempty"`
}

type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Meta    *Meta  `json:"meta,omitempty"`
}

type ErrorResponse struct {
	Success bool              `json:"success"`
	Error   ErrorResponseItem `json:"error"`
}

type ErrorResponseItem struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Details []ValidationError `json:"detail,omitempty"`
}

// -----------------------------------
// JSON response writer function
// -----------------------------------
func WriteJSON(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	if body == nil {
		return
	}
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		http.Error(w, `{"success":false,"error":{"message":"failed to encode response"}}`, http.StatusInternalServerError)
	}
}

// -----------------------------
// Success helpers
// -----------------------------
func Success(w http.ResponseWriter, data any) {
	WriteJSON(w, http.StatusOK, &SuccessResponse{
		Success: true,
		Data:    data,
	})
}

func SuccessWithMessage(w http.ResponseWriter, message string) {
	WriteJSON(w, http.StatusOK, &SuccessResponse{
		Success: true,
		Message: message,
	})
}

func SuccessWithMeta(w http.ResponseWriter, data any, meta *Meta) {
	WriteJSON(w, http.StatusOK, &SuccessResponse{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}

func Created(w http.ResponseWriter, data any) {
	WriteJSON(w, http.StatusCreated, &SuccessResponse{
		Success: true,
		Data:    data,
	})
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

//-----------------------------------
// Error helpers
//-----------------------------------

func Error(w http.ResponseWriter, err error) {
	appErr, ok := AsAppError(err)
	if !ok {
		appErr = WarpInternal(err)
	}

	WriteJSON(w, appErr.HTTPStatus, &ErrorResponse{
		Success: false,
		Error: ErrorResponseItem{
			Code:    string(appErr.Code),
			Message: appErr.Message,
			Details: appErr.Details,
		},
	})
}

func ErrorWithStatus(w http.ResponseWriter, statusCode int, code string, message string) {
	WriteJSON(w, statusCode, ErrorResponse{
		Success: false,
		Error: ErrorResponseItem{
			Code:    code,
			Message: message,
		},
	})
}

// Convinence helper - menggunakan Error() untuk http status code tertentu

func BadRequest(w http.ResponseWriter, message string) {
	Error(w, ErrBadRequest.WithMessage(message))
}

func Unauthorized(w http.ResponseWriter, message string) {
	Error(w, ErrUnautorized.WithMessage(message))
}

func NotFound(w http.ResponseWriter, message string) {
	Error(w, ErrNotFound.WithMessage(message))
}

func InternalServerError(w http.ResponseWriter, message string) {
	Error(w, ErrInternal.WithMessage(message))
}

//-------------------------------
// Paignation Helper
//-------------------------------

func BuildMeta(total int64, page, limit int) Meta {
	if limit == 0 {
		limit = 10
	}
	if page == 0 {
		page = 1
	}

	totalPage := total / int64(limit)
	if total%int64(limit) > 0 {
		totalPage++
	}

	return Meta{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: int(totalPage),
	}
}

// Parse json request body

// DecodeJSON
func DecodeJSON(r *http.Request, target any) error {
	if r.Body == nil {
		return ErrBadRequest.WithMessage("Request body is required")
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(target); err != nil {
		if errors.Is(err, io.EOF) {
			return ErrBadRequest.WithMessage("Request body cannot be empty")
		}
		return ErrBadRequest.WithError(err)
	}

	return nil
}
