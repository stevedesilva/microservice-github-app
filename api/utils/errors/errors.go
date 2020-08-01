package errors

import "net/http"

// ApiError exposed to client
type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	AStatus  int    `json:"status"`
	AMessage string `json:"message"`
	AnError  string `json:"error,omitempty"`
}

func (s *apiError) Status() int {
	return s.AStatus
}

func (s *apiError) Message() string {
	return s.AMessage
}

func (s *apiError) Error() string {
	return s.AnError
}

// NewNotFoundError error
func NewNotFoundError(message string) ApiError {
	return &apiError{AStatus: http.StatusNotFound, AMessage: message}
}

// NewInternalServerError error
func NewInternalServerError(message string) ApiError {
	return &apiError{AStatus: http.StatusInternalServerError, AMessage: message}
}

// NewBadRequestError error
func NewBadRequestError(message string) ApiError {
	return &apiError{AStatus: http.StatusBadRequest, AMessage: message}
}

// NewApiError error
func NewApiError(status int, message string) ApiError {
	return &apiError{
		AStatus:  status,
		AMessage: message,
	}
}
