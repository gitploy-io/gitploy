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
	// ErrorCodeDeploymentInvalid is the payload is invalid.
	ErrorCodeDeploymentInvalid ErrorCode = "deployment_invalid"
	// ErrorCodeDeploymentLocked is when the environment is locked.
	ErrorCodeDeploymentLocked ErrorCode = "deployment_locked"
	// ErrorCodeDeploymentUnapproved is when the deployment is not approved.
	ErrorCodeDeploymentNotApproved ErrorCode = "deployment_not_approved"
	// ErrorCodeDeploymentStatusNotWaiting is the status must be 'waiting' to create a remote deployment.
	ErrorCodeDeploymentStatusInvalid ErrorCode = "deployment_status_invalid"
	// ErrorCodeDeploymentUndeployable is that the merge conflict occurs or a commit status has failed.
	ErrorCodeDeploymentUndeployable ErrorCode = "deployment_undeployable"

	// ErrorCodeLicenseDecode is the error when the license is decoded.
	ErrorCodeLicenseDecode ErrorCode = "license_decode"
	// ErrorCodeLicenseRequired is that the license is required.
	ErrorCodeLicenseRequired ErrorCode = "license_required"

	// General purpose error codes.
	ErrorCodeInvalidRequest      ErrorCode = "invalid_request"
	ErrorPermissionRequired      ErrorCode = "permission_required"
	ErrorCodeNotFound            ErrorCode = "not_found"
	ErrorCodeUnprocessableEntity ErrorCode = "unprocessable_entity"
	ErrorCodeInternalError       ErrorCode = "internal_error"
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
