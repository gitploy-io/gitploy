package errors

import (
	"errors"
	"fmt"
)

type RefNotFoundError struct {
	Ref string
}

func (e *RefNotFoundError) Error() string {
	return fmt.Sprintf("%s is not found.", e.Ref)
}

func IsRefNotFoundError(err error) bool {
	var e *RefNotFoundError
	return errors.As(err, &e)
}
