package exception

// Basic error checking,
// if error is not nil.
func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}
