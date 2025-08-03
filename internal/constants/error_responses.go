package constants

import "net/http"

type ErrorResponse struct {
	Code    string `json:"error_code"`
	Message string `json:"error_message"`
	Status  int    `json:"-"`
}

// Error definitions
var (
	ErrInvalidPayload = ErrorResponse{
		Code:    "INVALID_PAYLOAD",
		Message: "Invalid request payload",
		Status:  http.StatusBadRequest,
	}

	ErrDBInsertFailed = ErrorResponse{
		Code:    "DB_INSERT_FAILED",
		Message: "Failed to insert order into the database",
		Status:  http.StatusInternalServerError,
	}

	ErrJWTGeneration = ErrorResponse{
		Code:    "JWT_GENERATION_FAILED",
		Message: "Unable to generate JWT token",
		Status:  http.StatusInternalServerError,
	}

	ErrSendWhatsApp = ErrorResponse{
		Code:    "WHATSAPP_SEND_FAILED",
		Message: "Failed to send WhatsApp message",
		Status:  http.StatusInternalServerError,
	}

	ErrInternal = ErrorResponse{
		Code:    "INTERNAL_ERROR",
		Message: "Something went wrong",
		Status:  http.StatusInternalServerError,
	}
)
