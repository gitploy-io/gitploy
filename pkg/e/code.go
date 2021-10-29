package e

import (
	"fmt"
)

const (
	// ErrorCodeMergeConflict is that the ref can't be merged into the main branch.
	ErrorCodeMergeConflict ErrorCode = "merge_conflict"

	// ErrorCodeLicenseDecode is that the license.
	ErrorCodeLicenseDecode ErrorCode = "license_decode"

	ErrorCodeInternalError ErrorCode = "internal_error"
)

type (
	ErrorStatus int
	ErrorCode   string

	Error struct {
		Code ErrorCode
		Wrap error
	}
)

func NewError(code ErrorCode, wrap error) *Error {
	return &Error{
		Code: code,
		Wrap: wrap,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %s, message: %s : %s", e.Code, GetMessage(e.Code), e.Wrap)
}
