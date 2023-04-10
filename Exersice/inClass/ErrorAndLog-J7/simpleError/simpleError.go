package simpleError

type SimpleError struct{
	Message string
	Operation string
}

func (s SimpleError) Error() string{
	return s.Message
}