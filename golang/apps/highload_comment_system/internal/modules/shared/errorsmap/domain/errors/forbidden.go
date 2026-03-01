package errors

type ForbiddenError struct {
	Message string `json:"message"`
}

func (e *ForbiddenError) Error() string {
	return e.Message
}

func NewForbiddenError(message string) *ForbiddenError {
	return &ForbiddenError{Message: message}
}
