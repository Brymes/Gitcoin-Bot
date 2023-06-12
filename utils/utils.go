package utils

import (
	"bytes"
	"crypto/rand"
	"log"
	"runtime/debug"
	"time"
)

func GenerateUniqueID(length int, charset string) string {
	bytes := make([]byte, length)

	var chars string

	switch charset {
	case "alphabets":
		chars = "abcdefghijklmnopqrstuvwxyz"
	case "numeric":
		chars = "0123456789"
	case "alphanumeric":
		chars = "0123456789abcdefghijklmnopqrstuvwxyz"
	default:
		chars = "0123456789abcdefghijklmnopqrstuvwxyz"
	}

	if _, err := rand.Read(bytes); err != nil {
		panic("Internal Server Error whilst Generating password")
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes)
}

func OnErrorPanic(err error, helpText string, logger *log.Logger) {
	if err != nil {
		logger.Panicf("%s: \n, %v", helpText, err)
	}
}

func HandlePanicMacro(err interface{}, logger *log.Logger) bool {
	if err != nil {
		if logger == nil {
			logger = log.Default()
		}
		logger.Println(err)
		logger.Println(string(debug.Stack()))
		return true
	}
	return false
}
func PeriodicLog(buff *bytes.Buffer) {
	ticker := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-ticker.C:
			if buff.Len() == 0 {
				log.Println(buff)
			}
			buff.Reset()
		default:
			continue
		}
	}

}
