package response

import "fmt"

type HttpError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (e HttpError) Error() string {
	return fmt.Sprintf("status: %d, message: %s", e.StatusCode, e.Message)
}

func NewHttpError(statusCode int, message string) HttpError {
	return HttpError{
		StatusCode: statusCode,
		Message:    message,
	}
}
