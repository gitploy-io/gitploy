package e

import "net/http"

var messages = map[ErrorCode]string{
	ErrorCodeDeploymentConflict:     "The conflict occurs, please retry.",
	ErrorCodeDeploymentUndeployable: "There is merge conflict or a commit status check failed.",
	ErrorCodeDeploymentInvalid:      "The validation has failed.",
	ErrorCodeLicenseDecode:          "Decoding the license is failed.",
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
	ErrorCodeDeploymentConflict:     http.StatusUnprocessableEntity,
	ErrorCodeDeploymentUndeployable: http.StatusUnprocessableEntity,
	ErrorCodeDeploymentInvalid:      http.StatusUnprocessableEntity,
	ErrorCodeLicenseDecode:          http.StatusUnprocessableEntity,
	ErrorCodeInternalError:          http.StatusInternalServerError,
}

func GetHttpCode(code ErrorCode) int {
	httpCode, ok := httpCodes[code]
	if !ok {
		return http.StatusInternalServerError
	}

	return httpCode
}
