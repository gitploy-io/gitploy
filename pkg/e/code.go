package e

import (
	"fmt"
)

const (
	// ErrorCodeDeploymentConflict is the deployment number is conflicted.
	ErrorCodeDeploymentConflict ErrorCode = "deployment_conflict"
	// ErrorCodeDeploymentInvalid is the payload is invalid.
	ErrorCodeDeploymentInvalid ErrorCode = "deployment_invalid"
	// ErrorCodeDeploymentLocked is when the environment is locked.
	ErrorCodeDeploymentLocked ErrorCode = "deployment_locked"
	// ErrorCodeDeploymentUndeployable is that the merge conflict occurs or a commit status has failed.
	ErrorCodeDeploymentUndeployable ErrorCode = "deployment_undeployable"

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

func (e *Error) Unwrap() error {
	return e.Wrap
}
