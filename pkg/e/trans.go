package e

import "net/http"

var messages = map[ErrorCode]string{
	ErrorCodeConfigParseError:        "The configuration is invalid.",
	ErrorCodeConfigRegexpError:       "The regexp is invalid.",
	ErrorCodeDeploymentConflict:      "The conflict occurs, please retry.",
	ErrorCodeDeploymentInvalid:       "The validation has failed.",
	ErrorCodeDeploymentLocked:        "The environment is locked.",
	ErrorCodeDeploymentNotApproved:   "The deployment is not approved",
	ErrorCodeDeploymentStatusInvalid: "The deployment status is invalid",
	ErrorCodeLicenseDecode:           "Decoding the license is failed.",
	ErrorCodeLicenseRequired:         "The license is required.",
	ErrorCodeParameterInvalid:        "Invalid request parameter.",
	ErrorPermissionRequired:          "The permission is required",
	ErrorCodeEntityNotFound:          "It is not found.",
	ErrorCodeEntityUnprocessable:     "Invalid request payload.",
	ErrorCodeInternalError:           "Server internal error.",
}

func GetMessage(code ErrorCode) string {
	message, ok := messages[code]
	if !ok {
		return string(code)
	}

	return message
}

var httpCodes = map[ErrorCode]int{
	ErrorCodeConfigParseError:        http.StatusUnprocessableEntity,
	ErrorCodeConfigRegexpError:       http.StatusUnprocessableEntity,
	ErrorCodeDeploymentConflict:      http.StatusUnprocessableEntity,
	ErrorCodeDeploymentInvalid:       http.StatusUnprocessableEntity,
	ErrorCodeDeploymentLocked:        http.StatusUnprocessableEntity,
	ErrorCodeDeploymentNotApproved:   http.StatusUnprocessableEntity,
	ErrorCodeDeploymentStatusInvalid: http.StatusUnprocessableEntity,
	ErrorCodeLicenseDecode:           http.StatusUnprocessableEntity,
	ErrorCodeLicenseRequired:         http.StatusPaymentRequired,
	ErrorCodeParameterInvalid:        http.StatusBadRequest,
	ErrorCodeEntityNotFound:          http.StatusNotFound,
	ErrorPermissionRequired:          http.StatusForbidden,
	ErrorCodeEntityUnprocessable:     http.StatusUnprocessableEntity,
	ErrorCodeInternalError:           http.StatusInternalServerError,
}

func GetHttpCode(code ErrorCode) int {
	httpCode, ok := httpCodes[code]
	if !ok {
		return http.StatusInternalServerError
	}

	return httpCode
}
