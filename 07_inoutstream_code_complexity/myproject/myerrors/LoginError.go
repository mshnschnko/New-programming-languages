package myerrors

type InvalidLoginDataError struct {
	Message string
}

func (e *InvalidLoginDataError) Error() string {
	return e.Message
}
