package myerrors

type IncorrectPasswordError struct {
	Message string
}

func (e *IncorrectPasswordError) Error() string {
	return e.Message
}

type InvalidPasswordError struct {
	Message string
}

func (e *InvalidPasswordError) Error() string {
	return e.Message
}
