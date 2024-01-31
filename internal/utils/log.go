package utils

import (
	"log"
	"runtime"
)

func LogWithCallerInfo(message string) {
	// Retrieve caller information
	_, file, line, ok := runtime.Caller(1)

	if ok {
		// Print log message with caller information
		log.Printf("[%s:%d] %s", file, line, message)
	} else {
		// Unable to retrieve caller information
		log.Println("Unable to retrieve caller information.")
	}
}
