package exception

type BadRequestError struct {
	Message string
}

// Error function is the conventional function,
// for representing an error message.
func (b BadRequestError) Error() string {
	return b.Message
}
