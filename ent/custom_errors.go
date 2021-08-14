package ent

import (
	"errors"
	"fmt"
)

type EagerLoadingError struct {
	Edge string
}

func (e *EagerLoadingError) Error() string {
	return fmt.Sprintf("The %s edge is not found.", e.Edge)
}

func IsConfigNotFoundError(err error) bool {
	var e *EagerLoadingError
	return errors.As(err, &e)
}
