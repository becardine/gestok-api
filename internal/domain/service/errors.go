package service

import "fmt"

type ErrNotFound struct {
	Entity string
	ID     interface{}
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("%s with ID %v not found", e.Entity, e.ID)
}
