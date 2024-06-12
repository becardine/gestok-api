package errors

import "fmt"

type EntityValidationError struct {
	Field string `json:"field"`
	Rule  string `json:"rule"`
	Value string `json:"value"`
}

func NewEntityValidationError(field, rule, value string) *EntityValidationError {
	return &EntityValidationError{
		Field: field,
		Rule:  rule,
		Value: value,
	}
}

func (e *EntityValidationError) Error() string {
	return fmt.Sprintf("validation error: field '%s' failed rule '%s' with value '%s'", e.Field, e.Rule, e.Value)
}
