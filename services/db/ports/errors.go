package db

import "fmt"

type ValidationError struct {
	Field string
	Msg   string
}
func (v *ValidationError) Error() string {
	return fmt.Sprintf("validation failed: %s: %s", v.Field, v.Msg)
}

type StringLengthError struct {
	Field string
}
func (s *StringLengthError) Error() string {
	return fmt.Sprintf("field %s: exceeded size", s.Field)
}