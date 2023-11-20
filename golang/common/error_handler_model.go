package common

type ErrorHandler struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
