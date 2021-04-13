package utils

import "log"

func LogIfError(message string, err interface{}) {
	if err != nil {
		log.Println(message, err)
	}
}
