package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      error  `json:"error,omitempty"`
}

func NewHTTPError(w http.ResponseWriter, statusCode int, message string, err error) {
	httpError := HTTPError{
		StatusCode: statusCode,
		Message:    message,
		Error:      err,
	}

	w.WriteHeader(statusCode)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(httpError)
}

func (e HTTPError) ErrorMessage() string {
	if e.Error != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Error)
	}
	return e.Message
}
