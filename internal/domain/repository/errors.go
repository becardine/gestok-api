package repository

import "fmt"

type ErrProductNotFound struct {
	ProductID int
}

func (e *ErrProductNotFound) Error() string {
	return fmt.Sprintf("product with ID %d not found", e.ProductID)
}
