package errors

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewValidationError(
	msg string,
) *APIError {

	return &APIError{
		Code:    "VALIDATION_ERROR",
		Message: msg,
	}
}
