package e

import "net/http"

var messages = map[ErrorCode]string{
	ErrorCodeConfigInvalid:           "The configuration is invalid.",
	ErrorCodeConfigUndefinedEnv:      "The environment is not defined in the configuration.",
	ErrorCodeDeploymentConflict:      "The conflict occurs, please retry.",
	ErrorCodeDeploymentInvalid:       "The validation has failed.",
	ErrorCodeDeploymentLocked:        "The environment is locked.",
	ErrorCodeDeploymentFrozen:        "It is in the deploy freeze window.",
	ErrorCodeDeploymentNotApproved:   "The deployment is not approved.",
	ErrorCodeDeploymentSerialization: "There is a running deployment.",
	ErrorCodeDeploymentStatusInvalid: "The deployment status is invalid.",
	ErrorCodeEntityNotFound:          "It is not found.",
	ErrorCodeEntityUnprocessable:     "Invalid request payload.",
	ErrorCodeInternalError:           "Server internal error.",
	ErrorCodeLockAlreadyExist:        "The environment is already locked",
	ErrorCodeLicenseDecode:           "Decoding the license is failed.",
	ErrorCodeLicenseRequired:         "The license is required.",
	ErrorCodeParameterInvalid:        "Invalid request parameter.",
	ErrorPermissionRequired:          "The permission is required.",
	ErrorRepoUniqueName:              "The same repository name already exists.",
}

func GetMessage(code ErrorCode) string {
	message, ok := messages[code]
	if !ok {
		return string(code)
	}

	return message
}

var httpCodes = map[ErrorCode]int{
	ErrorCodeConfigInvalid:           http.StatusUnprocessableEntity,
	ErrorCodeConfigUndefinedEnv:      http.StatusUnprocessableEntity,
	ErrorCodeDeploymentConflict:      http.StatusUnprocessableEntity,
	ErrorCodeDeploymentInvalid:       http.StatusUnprocessableEntity,
	ErrorCodeDeploymentLocked:        http.StatusUnprocessableEntity,
	ErrorCodeDeploymentFrozen:        http.StatusUnprocessableEntity,
	ErrorCodeDeploymentNotApproved:   http.StatusUnprocessableEntity,
	ErrorCodeDeploymentSerialization: http.StatusUnprocessableEntity,
	ErrorCodeDeploymentStatusInvalid: http.StatusUnprocessableEntity,
	ErrorCodeEntityNotFound:          http.StatusNotFound,
	ErrorCodeEntityUnprocessable:     http.StatusUnprocessableEntity,
	ErrorCodeInternalError:           http.StatusInternalServerError,
	ErrorCodeLockAlreadyExist:        http.StatusUnprocessableEntity,
	ErrorCodeLicenseDecode:           http.StatusUnprocessableEntity,
	ErrorCodeLicenseRequired:         http.StatusPaymentRequired,
	ErrorCodeParameterInvalid:        http.StatusBadRequest,
	ErrorPermissionRequired:          http.StatusForbidden,
	ErrorRepoUniqueName:              http.StatusUnprocessableEntity,
}
