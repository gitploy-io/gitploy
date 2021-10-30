package e

import "net/http"

var messages = map[ErrorCode]string{
	ErrorCodeConfigNotFound:         "The configuration file is not found.",
	ErrorCodeConfigParseError:       "The configuration is invalid.",
	ErrorCodeDeploymentConflict:     "The conflict occurs, please retry.",
	ErrorCodeDeploymentInvalid:      "The validation has failed.",
	ErrorCodeDeploymentLocked:       "The environment is locked.",
	ErrorCodeDeploymentUnapproved:   "The deployment is not approved",
	ErrorCodeDeploymentUndeployable: "There is merge conflict or a commit status check failed.",
	ErrorCodeLicenseDecode:          "Decoding the license is failed.",
	ErrorCodeRefNotFound:            "The reference is not found.",
	ErrorCodeInvalidRequest:         "Invalid request parameter.",
	ErrorCodeNotFound:               "It is not found.",
	ErrorCodeInternalError:          "Server internal error.",
}

func GetMessage(code ErrorCode) string {
	message, ok := messages[code]
	if !ok {
		return string(code)
	}

	return message
}

var httpCodes = map[ErrorCode]int{
	ErrorCodeConfigNotFound:         http.StatusNotFound,
	ErrorCodeConfigParseError:       http.StatusUnprocessableEntity,
	ErrorCodeDeploymentConflict:     http.StatusUnprocessableEntity,
	ErrorCodeDeploymentInvalid:      http.StatusUnprocessableEntity,
	ErrorCodeDeploymentLocked:       http.StatusUnprocessableEntity,
	ErrorCodeDeploymentUnapproved:   http.StatusUnprocessableEntity,
	ErrorCodeDeploymentUndeployable: http.StatusUnprocessableEntity,
	ErrorCodeLicenseDecode:          http.StatusUnprocessableEntity,
	ErrorCodeRefNotFound:            http.StatusNotFound,
	ErrorCodeInvalidRequest:         http.StatusBadRequest,
	ErrorCodeNotFound:               http.StatusNotFound,
	ErrorCodeInternalError:          http.StatusInternalServerError,
}

func GetHttpCode(code ErrorCode) int {
	httpCode, ok := httpCodes[code]
	if !ok {
		return http.StatusInternalServerError
	}

	return httpCode
}
