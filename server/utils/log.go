package utils

import "log"

func Log(message string, err interface{}) {
	if err != nil {
		log.Println(message, err)
	}
}
