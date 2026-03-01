package errors

type UnauthenticatedError struct {
	Message string `json:"message"`
}

func (e *UnauthenticatedError) Error() string {
	return e.Message
}

func NewUnauthenticatedError(message string) *UnauthenticatedError {
	return &UnauthenticatedError{Message: message}
}
