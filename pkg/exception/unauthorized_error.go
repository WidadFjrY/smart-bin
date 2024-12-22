package exception

type Unauthorized struct {
	error string
}

func NewUnauthorized(message string) Unauthorized {
	return Unauthorized{error: message}
}
