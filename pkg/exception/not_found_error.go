package exception

type NotFoundError struct {
	error string
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError{error: message}
}
