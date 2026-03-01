package errors

type InvalidArgumentError struct {
	Message string `json:"message"`
}

func (e *InvalidArgumentError) Error() string {
	return e.Message
}

func NewInvalidArgumentError(message string) *InvalidArgumentError {
	return &InvalidArgumentError{Message: message}
}
