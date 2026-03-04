package errors

type NotFoundError struct {
	Message string `json:"message"`
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{Message: message}
}
