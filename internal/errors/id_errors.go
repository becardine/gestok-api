package errors

import "fmt"

type ErrInvalidIDType struct {
	Message string
}

func (e *ErrInvalidIDType) Error() string {
	return fmt.Sprintf("invalid ID type: %s", e.Message)
}

func NewErrInvalidIDType(message string) *ErrInvalidIDType {
	return &ErrInvalidIDType{
		Message: message,
	}
}
