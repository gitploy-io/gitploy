package e

import "testing"

func Test_GetMessage(t *testing.T) {
	t.Run("Return code when the message is emtpy.", func(t *testing.T) {
		const ErrorCodeEmpty ErrorCode = "emtpy"

		message := GetMessage(ErrorCodeEmpty)
		if message != string(ErrorCodeEmpty) {
			t.Fatalf("GetMessage = %s, wanted %s", message, string(ErrorCodeEmpty))
		}
	})
}
