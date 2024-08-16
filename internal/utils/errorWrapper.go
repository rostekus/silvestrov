package utils

import "fmt"

func ErrorWrapper(msg string, err error) string {
	if err == nil {
		return msg
	}
	return fmt.Sprintf("%s: %v", msg, err)
}
