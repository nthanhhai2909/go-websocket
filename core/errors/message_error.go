package errors

type MessageError struct {
	message string
}

func (e MessageError) Error() string { return e.message }
