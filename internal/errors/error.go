package errors

import "errors"

var (
	ErrInvalidID        = errors.New("invalid id")
	ErrNotFound         = errors.New("not found")
	ErrorInternalServer = errors.New("internal server error")
	ErrorInvalidRequest = errors.New("invalid request")
)

type ErrorHandler struct {
	Message string `json:"message"`
}
