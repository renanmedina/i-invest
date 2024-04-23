package utils

import "log"

func NewApplicationLogger() *log.Logger {
	return log.Default()
}
