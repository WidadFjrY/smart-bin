package exception

type BadRequestError struct {
	error string
}

func NewBadRequestError(message string) BadRequestError {
	return BadRequestError{error: message}
}
