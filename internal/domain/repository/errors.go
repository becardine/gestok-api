package repository

import "fmt"

type ErrProductNotFound struct {
	ProductID string
}

func (e *ErrProductNotFound) Error() string {
	return fmt.Sprintf("product with ID %s not found", e.ProductID)
}
