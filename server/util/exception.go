package util

func ReturnErrorIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}
