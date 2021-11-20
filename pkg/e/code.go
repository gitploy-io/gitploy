package e

import (
	"errors"
	"fmt"
)

const (
	// ErrorCodeConfigParseError is that an error occurs when it parse the file.
	ErrorCodeConfigParseError ErrorCode = "config_parse_error"
	// ErrorCodeConfigRegexpError is the regexp(re2) is invalid.
	ErrorCodeConfigRegexpError ErrorCode = "config_regexp_error"

	// ErrorCodeDeploymentConflict is the deployment number is conflicted.
	ErrorCodeDeploymentConflict ErrorCode = "deployment_conflict"
	// ErrorCodeDeploymentInvalid is the payload is invalid when it posts a remote deployment.
	ErrorCodeDeploymentInvalid ErrorCode = "deployment_invalid"
	// ErrorCodeDeploymentLocked is when the environment is locked.
	ErrorCodeDeploymentLocked ErrorCode = "deployment_locked"
	// ErrorCodeDeploymentUnapproved is when the deployment is not approved.
	ErrorCodeDeploymentNotApproved ErrorCode = "deployment_not_approved"
	// ErrorCodeDeploymentStatusNotWaiting is the status must be 'waiting' to create a remote deployment.
	ErrorCodeDeploymentStatusInvalid ErrorCode = "deployment_status_invalid"

	// ErrorCodeLicenseDecode is the error when the license is decoded.
	ErrorCodeLicenseDecode ErrorCode = "license_decode"
	// ErrorCodeLicenseRequired is that the license is required.
	ErrorCodeLicenseRequired ErrorCode = "license_required"

	// ErrorCodeEntityNotFound is the entity is not found.
	ErrorCodeEntityNotFound ErrorCode = "entity_not_found"
	// ErrorCodeEntityUnprocessable is the entity is unprocessable.
	ErrorCodeEntityUnprocessable ErrorCode = "entity_unprocessable"

	// ErrorCodeParameterInvalid is a parameter of a request is invalid.
	ErrorCodeParameterInvalid ErrorCode = "parameter_invalid"

	// ErrorPermissionRequired is the permission is required to access.
	ErrorPermissionRequired ErrorCode = "permission_required"

	// ErrorCodeInternalError is the internal error couldn't be handled.
	ErrorCodeInternalError ErrorCode = "internal_error"
)

type (
	ErrorStatus int
	ErrorCode   string

	Error struct {
		Code    ErrorCode
		Message string
		Wrap    error
	}
)

func NewError(code ErrorCode, wrap error) *Error {
	return &Error{
		Code:    code,
		Message: GetMessage(code),
		Wrap:    wrap,
	}
}

func NewErrorWithMessage(code ErrorCode, message string, wrap error) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Wrap:    wrap,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %s, message: %s, wrap: %s", e.Code, e.Message, e.Wrap)
}

func (e *Error) Unwrap() error {
	return e.Wrap
}

func IsError(err error) bool {
	var ge *Error
	return errors.As(err, &ge)
}

// HasErrorCode verify the type of error and the code.
func HasErrorCode(err error, codes ...ErrorCode) bool {
	var ge *Error
	if !errors.As(err, &ge) {
		return false
	}

	for _, code := range codes {
		if ge.Code == code {
			return true
		}
	}

	return false
}
