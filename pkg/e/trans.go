package e

import "net/http"

var messages = map[ErrorCode]string{
	ErrorCodeConfigInvalid:           "The configuration is invalid.",
	ErrorCodeConfigUndefinedEnv:      "The environment is not defined in the configuration.",
	ErrorCodeConfigRegexpError:       "The regexp is invalid.",
	ErrorCodeDeploymentConflict:      "The conflict occurs, please retry.",
	ErrorCodeDeploymentInvalid:       "The validation has failed.",
	ErrorCodeDeploymentLocked:        "The environment is locked.",
	ErrorCodeDeploymentNotApproved:   "The deployment is not approved",
	ErrorCodeDeploymentStatusInvalid: "The deployment status is invalid",
	ErrorCodeEntityNotFound:          "It is not found.",
	ErrorCodeEntityUnprocessable:     "Invalid request payload.",
	ErrorCodeInternalError:           "Server internal error.",
	ErrorCodeLockAlreadyExist:        "The environment is already locked",
	ErrorCodeLicenseDecode:           "Decoding the license is failed.",
	ErrorCodeLicenseRequired:         "The license is required.",
	ErrorCodeParameterInvalid:        "Invalid request parameter.",
	ErrorPermissionRequired:          "The permission is required",
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
	ErrorCodeConfigRegexpError:       http.StatusUnprocessableEntity,
	ErrorCodeDeploymentConflict:      http.StatusUnprocessableEntity,
	ErrorCodeDeploymentInvalid:       http.StatusUnprocessableEntity,
	ErrorCodeDeploymentLocked:        http.StatusUnprocessableEntity,
	ErrorCodeDeploymentNotApproved:   http.StatusUnprocessableEntity,
	ErrorCodeDeploymentStatusInvalid: http.StatusUnprocessableEntity,
	ErrorCodeEntityNotFound:          http.StatusNotFound,
	ErrorCodeEntityUnprocessable:     http.StatusUnprocessableEntity,
	ErrorCodeInternalError:           http.StatusInternalServerError,
	ErrorCodeLockAlreadyExist:        http.StatusUnprocessableEntity,
	ErrorCodeLicenseDecode:           http.StatusUnprocessableEntity,
	ErrorCodeLicenseRequired:         http.StatusPaymentRequired,
	ErrorCodeParameterInvalid:        http.StatusBadRequest,
	ErrorPermissionRequired:          http.StatusForbidden,
}

func GetHttpCode(code ErrorCode) int {
	httpCode, ok := httpCodes[code]
	if !ok {
		return http.StatusInternalServerError
	}

	return httpCode
}
