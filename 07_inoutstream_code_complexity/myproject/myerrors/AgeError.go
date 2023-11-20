package myerrors

type AgeError struct {
	Message string
}

func (e *AgeError) Error() string {
	return e.Message
}
