package e

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	// ErrorCodeConfigInvalid is that an error occurs when it parse the file.
	ErrorCodeConfigInvalid ErrorCode = "config_parse_error"
	// ErrorCodeConfigUndefinedEnv is that the environment is not defined in the configuration file.
	ErrorCodeConfigUndefinedEnv ErrorCode = "config_undefined_env"

	// ErrorCodeDeploymentConflict is the deployment number is conflicted.
	ErrorCodeDeploymentConflict ErrorCode = "deployment_conflict"
	// ErrorCodeDeploymentInvalid is the payload is invalid.
	ErrorCodeDeploymentInvalid ErrorCode = "deployment_invalid"
	// ErrorCodeDeploymentLocked is when the environment is locked.
	ErrorCodeDeploymentLocked ErrorCode = "deployment_locked"
	// ErrorCodeDeploymentFrozen is when the time in in the freeze window.
	ErrorCodeDeploymentFrozen ErrorCode = "deployment_frozen"
	// ErrorCodeDeploymentUnapproved is when the deployment is not approved.
	ErrorCodeDeploymentNotApproved ErrorCode = "deployment_not_approved"
	// ErrorCodeDeploymentSerialization is the serialization error.
	ErrorCodeDeploymentSerialization ErrorCode = "deployment_serialization"
	// ErrorCodeDeploymentStatusNotWaiting is the status must be 'waiting' to create a remote deployment.
	ErrorCodeDeploymentStatusInvalid ErrorCode = "deployment_status_invalid"

	// ErrorCodeEntityNotFound is the entity is not found.
	// Entity is a resource of store or scm.
	ErrorCodeEntityNotFound ErrorCode = "entity_not_found"
	// ErrorCodeEntityUnprocessable is the entity is unprocessable.
	ErrorCodeEntityUnprocessable ErrorCode = "entity_unprocessable"

	// ErrorCodeInternalError is the internal error couldn't be handled.
	ErrorCodeInternalError ErrorCode = "internal_error"

	// ErrorCodeLockAlreadyExist is that the environment is already locked.
	ErrorCodeLockAlreadyExist ErrorCode = "lock_already_exist"

	// ErrorCodeLicenseDecode is the error when the license is decoded.
	ErrorCodeLicenseDecode ErrorCode = "license_decode"
	// ErrorCodeLicenseRequired is that the license is required.
	ErrorCodeLicenseRequired ErrorCode = "license_required"

	// ErrorCodeParameterInvalid is a parameter of a request is invalid.
	ErrorCodeParameterInvalid ErrorCode = "parameter_invalid"

	// ErrorPermissionRequired is the permission is required to access.
	ErrorPermissionRequired ErrorCode = "permission_required"

	// ErrorRepoUniqueName is the repository name must be unique.
	ErrorRepoUniqueName ErrorCode = "repo_unique_name"
)

type (
	ErrorStatus int
	ErrorCode   string

	Error struct {
		Code    ErrorCode
		Message string
		Wrap    error

		httpCode int
	}
)

func NewError(code ErrorCode, wrap error) *Error {
	return &Error{
		Code:     code,
		Message:  GetMessage(code),
		Wrap:     wrap,
		httpCode: mapHTTPCode(code),
	}
}

func NewErrorWithMessage(code ErrorCode, message string, wrap error) *Error {
	return &Error{
		Code:     code,
		Message:  message,
		Wrap:     wrap,
		httpCode: mapHTTPCode(code),
	}
}

func (e *Error) Error() string {
	msgs := []string{
		fmt.Sprintf("Code: %s", e.Code),
		fmt.Sprintf("Message: %s", e.Message),
	}

	if e.Wrap != nil {
		msgs = append(msgs, fmt.Sprintf("Wrap: %s", e.Wrap))
	}

	return strings.Join(msgs, ", ")
}

// GetHTTPCode returns the HTTP code.
func (e *Error) GetHTTPCode() int {
	return e.httpCode
}

// SetHTTPCode sets the HTTP code manually.
func (e *Error) SetHTTPCode(code int) {
	e.httpCode = code
}

func (e *Error) Unwrap() error {
	return e.Wrap
}

func mapHTTPCode(code ErrorCode) int {
	httpCode, ok := httpCodes[code]
	if !ok {
		return http.StatusInternalServerError
	}

	return httpCode
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
