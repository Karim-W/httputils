package httputils

type HttpErrorResponse struct {
	Code         int         `json:"code,omitempty"`
	Message      string      `json:"message,omitempty"`
	ErrorDetails string      `json:"errorDetails,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Transaction  *string     `json:"transaction,omitempty"`
}
type HttpSuccessResponse struct {
	Code        int         `json:"code,omitempty"`
	Message     string      `json:"message,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	Transaction *string     `json:"transaction,omitempty"`
}

func SuccessResponse(code int, message string, data interface{}, transactionId *string) (int, HttpSuccessResponse) {
	return code, HttpSuccessResponse{
		Code:        code,
		Message:     message,
		Data:        data,
		Transaction: transactionId,
	}
}

func ErrorResponse(code int, errorMessage string, errorDetails string, data interface{}, transactionId *string) (int, HttpErrorResponse) {
	return code, HttpErrorResponse{
		Code:         code,
		Message:      errorMessage,
		ErrorDetails: errorDetails,
		Data:         data,
		Transaction:  transactionId,
	}
}
