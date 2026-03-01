package errors

type PreconditionFailureError struct {
	Message string `json:"message"`
}

func (e *PreconditionFailureError) Error() string {
	return e.Message
}

func NewPreconditionFailureError(message string) *PreconditionFailureError {
	return &PreconditionFailureError{Message: message}
}
