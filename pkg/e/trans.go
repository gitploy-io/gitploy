package e

import "net/http"

var messages = map[ErrorCode]string{
	ErrorCodeMergeConflict: "There is merge conflict.",
	ErrorCodeLicenseDecode: "Decoding the license is failed.",
	ErrorCodeInternalError: "Server internal error.",
}

func GetMessage(code ErrorCode) string {
	message, ok := messages[code]
	if !ok {
		return string(code)
	}

	return message
}

var httpCodes = map[ErrorCode]int{
	ErrorCodeMergeConflict: http.StatusUnprocessableEntity,
	ErrorCodeLicenseDecode: http.StatusUnprocessableEntity,
	ErrorCodeInternalError: http.StatusInternalServerError,
}

func GetHttpCode(code ErrorCode) int {
	httpCode, ok := httpCodes[code]
	if !ok {
		return http.StatusInternalServerError
	}

	return httpCode
}
